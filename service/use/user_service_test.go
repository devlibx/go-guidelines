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
