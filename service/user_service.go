package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
}
