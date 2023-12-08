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
