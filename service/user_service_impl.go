package service

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Error          error
}

type Claims struct {
	User_id  string
	Username string
	Role_id  bool
	jwt.RegisteredClaims
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Role, err := strconv.Atoi(request.Role_id)
	// if err != nil {
	// 	helper.PanicIfError(err)
	// }

	user := domain.User{
		Username: request.Username,
		Password: hashedPassword, // Use the hashed password
		Email:    request.Email,
		Role_id:  request.Role_id,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		return web.UserResponse{}, err
	}

	err = ComparePassword(user.Password, request.Password)
	if err != nil {
		return web.UserResponse{}, err // Passwords don't match
	}

	token, err := GenerateToken(user.Username, user.Role_id, user.User_id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse, nil
}

func GenerateToken(username string, role bool, userId int, secretKey string) (string, error) {
	// Set claims
	claims := &Claims{
		Username: username,
		User_id:  strconv.Itoa(userId),
		Role_id:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SetCookie(token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	}

	return cookie
}

func HashPassword(password string) (string, error) {
	// Generate a hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
