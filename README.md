### How to avoid dependency of struct to make your code testable

This is a service which sends data to outside world (for demo we just print to console).

```go
package sender

import (
	"context"
	"fmt"
)

type Service struct {
}

func (s *Service) Send(ctx context.Context, data map[string]interface{}) error {
	fmt.Println("Sending data to outside world", data)
	return nil
}

```
---

#### Example where we used struct

Here we used struct, and it made it difficult to test the code. You will
see we used the `sender.Service` struct in the `UserServiceWithStructDependency` struct.
- now this package must know about the sender package
```go
package main

import (
	"context"
	"go-guidelines/sender"
)

// UserServiceWithStructDependency is a user service with struct dependency
// Now this package must know about the sender package
type UserServiceWithStructDependency struct {
	senderService *sender.Service
}

func (u *UserServiceWithStructDependency) SendData(ctx context.Context, data map[string]interface{}) error {
	return u.senderService.Send(ctx, data)
}

func NewUserServiceWithStructDependency(senderService *sender.Service) *UserServiceWithStructDependency {
	return &UserServiceWithStructDependency{senderService: senderService}
}

```

Here is a test for the above code. 
- TestSendData => we are able to test this method
- TestSendData_ToSimulateError => you will see we are not able to
test the error scenario. If we use the `sender.Service` then it is
sending data to outside world (and we do not get error)

```go
package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-guidelines/sender"
	"testing"
)

func TestSendData(t *testing.T) {

	// Create a new user service with struct dependency
	userService := NewUserServiceWithStructDependency(&sender.Service{})

	// Test sending data method
	err := userService.SendData(context.Background(), map[string]interface{}{"name": "test"})
	assert.NoError(t, err)
}

func TestSendData_ToSimulateError(t *testing.T) {

	// Create a new user service with struct dependency
	userService := NewUserServiceWithStructDependency(&sender.Service{})

	// Test sending data method - here we want to simulate error
	// -> Here it is very difficult to simulate error and test
	err := userService.SendData(context.Background(), map[string]interface{}{"name": "test"})
	assert.NoError(t, err)
}

```
---

#### Example where we used interface
In this case we will avoid anu dependency on the `sender` package. Here you will notice we do not need to 
know about sender package
```go
package main

import (
	"context"
)

type Sender interface {
	Send(ctx context.Context, data map[string]interface{}) error
}

// UserServiceWithoutAnyDependencyOnOtherPackage has no dependency on other package
// This says I will work with any one who can send data
type UserServiceWithoutAnyDependencyOnOtherPackage struct {
	senderService Sender
}

func (u *UserServiceWithoutAnyDependencyOnOtherPackage) SendData(ctx context.Context, data map[string]interface{}) error {
	return u.senderService.Send(ctx, data)
}

func NewUserServiceWithStructDependency(senderService Sender) *UserServiceWithoutAnyDependencyOnOtherPackage {
	return &UserServiceWithoutAnyDependencyOnOtherPackage{senderService: senderService}
}
```

Here is the test for the above code.
- TestSendData => we are able to test this method (same as previous case)
- TestSendData_ToSimulateError => here you will notice we are able to simulate error using a stub sender
```go
package main

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go-guidelines/sender"
	"testing"
)

func TestSendData(t *testing.T) {

	// Create a new user service with struct dependency
	// NOTE - we don't need sender package here - just added it to show that it is same as previous case
	userService := NewUserServiceWithStructDependency(&sender.Service{})

	// Test sending data method
	err := userService.SendData(context.Background(), map[string]interface{}{"name": "test"})
	assert.NoError(t, err)
}

// TestSendData_ToSimulateError is a test to simulate error
// NOTE - here we can add any mock or stub to simulate error
func TestSendData_ToSimulateError(t *testing.T) {

	// Create a new user service with struct dependency
	userService := NewUserServiceWithStructDependency(&stubSender{err: errors.New("expected error")})

	// Test sending data method - here we want to simulate error
	// -> Here it is very difficult to simulate error and test
	err := userService.SendData(context.Background(), map[string]interface{}{"name": "test"})
	assert.Error(t, err)
	assert.Equal(t, "expected error", err.Error())
}

type stubSender struct {
	err error
}

func (s stubSender) Send(ctx context.Context, data map[string]interface{}) error {
	return s.err
}

```