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

func (h *handler) RegisterHandler(c *gin.Context) {
	var input auth.InputRegister

	err := c.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	_, errRegister := h.service.RegisterService(input)

	if errRegister != nil {
		logrus.Fatal(errRegister.Error())
		return
	}

	response := utils.APIResponse("Register new account successfully", http.StatusOK, http.MethodGet, nil)
	c.JSON(http.StatusOK, response)
}
