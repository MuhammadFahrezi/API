package domain

import "time"

//user with token
type User struct {
	User_id    int
	Role_id    bool
	NPM        int
	Username   string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
	Token      string
}
