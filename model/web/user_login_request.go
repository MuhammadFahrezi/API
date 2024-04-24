package web

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,max=200,min=1"`
	Password string `json:"password" validate:"required,max=200,min=1"`
}