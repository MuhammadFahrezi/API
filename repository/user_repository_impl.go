package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(username, password, email, role_id, npm) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Email, user.Role_id, user.NPM)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.User_id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "select user_id, username, password, email, role_id from user where username = ? and password = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username, user.Password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email, &user.Role_id)
		if err != nil {
			panic(err)
		}
		return user
	} else {
		panic("username or password is incorrect")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select user_id, username, password, email, role_id from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email, &user.Role_id)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select user_id, username, password, email, role_id from user"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email, &user.Role_id)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByRole(ctx context.Context, tx *sql.Tx, role string) []domain.User {
	SQL := "select user_id, username, password, email, role_id from user where role = ?"
	rows, err := tx.QueryContext(ctx, SQL, role)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email, &user.Role_id)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "select user_id, username, password, email, role_id from user where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Username, &user.Password, &user.Email, &user.Role_id)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}
