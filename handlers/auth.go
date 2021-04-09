package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers"
	"github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service auth.Service
}

func NewHandler(service auth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input auth.InputRegister

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	resultRegister, errRegister := h.service.RegisterService(input)

	if errRegister != nil {
		response := utils.APIResponse("Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := utils.APIResponse("Register new account successfully", http.StatusOK, http.MethodPost, resultRegister)
	ctx.JSON(http.StatusOK, response)
}
