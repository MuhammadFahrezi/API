package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}

type Claims struct {
	Username string
	User_id  int
	Role_id  bool
	jwt.RegisteredClaims
}

func (middleware *Middleware) ApplyMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Extract the JWT token from the request header
		tokenString := request.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the JWT token and extract claims
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil // Replace "secret" with your actual secret key
		})

		if err != nil || !token.Valid || claims.Role_id == false {
			// If there's an error parsing the token or the token is invalid, respond with an Unauthorized status
			fmt.Println("Token invalid or parsing error:", err)
			helper.WriteResponse(writer, http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			})
			return
		}

		// Create a context with the user information and attach it to the request context
		ctx := context.WithValue(request.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "userId", claims.User_id)
		ctx = context.WithValue(ctx, "role", claims.Role_id)
		request = request.WithContext(ctx)

		// Proceed to the next handler with the updated request
		next(writer, request, params)
	}
}

func (middleware *Middleware) ApplyAdminMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Extract the JWT token from the request header
		tokenString := request.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the JWT token and extract claims
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil // Replace "secret" with your actual secret key
		})

		if err != nil || !token.Valid || claims.Role_id == true {
			// If there's an error parsing the token or the token is invalid, respond with an Unauthorized status
			fmt.Println("Token invalid or parsing error:", err)
			helper.WriteResponse(writer, http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			})
			return
		}

		// Create a context with the user information and attach it to the request context
		ctx := context.WithValue(request.Context(), "username", claims.Username)
		ctx = context.WithValue(ctx, "userId", claims.User_id)
		ctx = context.WithValue(ctx, "role", claims.Role_id)
		request = request.WithContext(ctx)

		// Proceed to the next handler with the updated request
		next(writer, request, params)
	}
}
