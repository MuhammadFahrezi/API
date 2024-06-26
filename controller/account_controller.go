package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AccountController interface {
	UserDetailByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}