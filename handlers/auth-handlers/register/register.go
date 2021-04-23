package handlerRegister

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/gin-rest-api/controllers/auth-controllers/register"
	util "github.com/restuwahyu13/gin-rest-api/utils"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.InputRegister

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

	resultRegister, errRegister := h.service.RegisterService(&input)

	if errRegister == "REGISTER_CONFLICT_409" {
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return
	}

	if errRegister == "REGISTER_FAILED_403" {
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
	accessToken, err := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

	if err != nil {
		util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	_, errSendMail := util.SendGridMail(resultRegister.Fullname, resultRegister.Email, "Activation Account", "template_register", accessToken)

	if errSendMail != nil {
		util.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	util.APIResponse(ctx, "Register new account successfully", http.StatusOK, http.MethodPost, nil)
}
