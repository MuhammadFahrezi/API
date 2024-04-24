package web

type UserResponse struct {
	User_id  int    `json:"user_id"`
	Role_id  string `json:"role_id"`
	NPM      int    `json:"npm"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
