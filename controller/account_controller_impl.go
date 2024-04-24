package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

func (controller *AccountControllerImpl) UserDetailByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	accountId := params.ByName("accountId")
	id, err := strconv.Atoi(accountId)
	helper.PanicIfError(err)

	accountResponse := controller.AccountService.UserDetailByID(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   accountResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
