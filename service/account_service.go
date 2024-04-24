package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type AccountService interface {
    UserDetailByID(ctx context.Context, userID int) web.UserResponse
}