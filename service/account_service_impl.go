package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	// "github.com/go-playground/validator/v10"
)

type AccountServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewAccountService(userRepository repository.UserRepository, DB *sql.DB) AccountService {
	return &AccountServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *AccountServiceImpl) UserDetailByID(ctx context.Context, userID int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userID)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}
