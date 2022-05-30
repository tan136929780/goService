/*
@Time : 2022/5/27 下午1:48
@Author : tan
@File : userGrpcService
@Software: GoLand
*/
package user

import (
	"context"
	"goMicroService/server-service/proto/proto"
)

type User struct{}

func NewUserGrpcService() *User {
	return &User{}
}

func (u *User) CreateUser(ctx context.Context, in *proto.UserRequest, out *proto.ResponseResult) error {
	return nil
}

func (u *User) UpdateUser(ctx context.Context, in *proto.UserRequest, out *proto.ResponseResult) error {
	return nil
}

func (u *User) SelectUser(ctx context.Context, in *proto.UserRequest, out *proto.ResponseResult) error {
	return nil
}

func (u *User) DeleteUser(ctx context.Context, in *proto.UserRequest, out *proto.ResponseResult) error {
	return nil
}
