package route

import (
	activationAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/activation"
	forgotAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/forgot"
	loginAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/login"
	registerAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/register"
	resendAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/resend"
	resetAuth "github.com/fstar-dev/restapi-gin/controllers/auth-controllers/reset"
	handlerActivation "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/activation"
	handlerForgot "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/forgot"
	handlerLogin "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/login"
	handlerRegister "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/register"
	handlerResend "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/resend"
	handlerReset "github.com/fstar-dev/restapi-gin/handlers/auth-handlers/reset"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	activationRepository := activationAuth.NewRepositoryActivation(db)
	activationService := activationAuth.NewServiceActivation(activationRepository)
	activationHandler := handlerActivation.NewHandlerActivation(activationService)

	resendRepository := resendAuth.NewRepositoryResend(db)
	resendService := resendAuth.NewServiceResend(resendRepository)
	resendHandler := handlerResend.NewHandlerResend(resendService)

	forgotRepository := forgotAuth.NewRepositoryForgot(db)
	forgotService := forgotAuth.NewServiceForgot(forgotRepository)
	forgotHandler := handlerForgot.NewHandlerForgot(forgotService)

	resetRepository := resetAuth.NewRepositoryReset(db)
	resetService := resetAuth.NewServiceReset(resetRepository)
	resetHandler := handlerReset.NewHandlerReset(resetService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/activation/:token", activationHandler.ActivationHandler)
	groupRoute.POST("/resend-token", resendHandler.ResendHandler)
	groupRoute.POST("/forgot-password", forgotHandler.ForgotHandler)
	groupRoute.POST("/change-password/:token", resetHandler.ResetHandler)

}
