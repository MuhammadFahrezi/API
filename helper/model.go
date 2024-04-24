package helper

import (
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:       category.Id,
		Category: category.Category,
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	role := "user"
	if !user.Role_id {
		role = "admin"
	}
	return web.UserResponse{
		User_id:  user.User_id,
		Role_id:  role,
		NPM:      user.NPM,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
