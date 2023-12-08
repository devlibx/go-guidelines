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
