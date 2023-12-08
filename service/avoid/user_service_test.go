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
