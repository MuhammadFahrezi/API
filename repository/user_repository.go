package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByRole(ctx context.Context, tx *sql.Tx, role string) []domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
