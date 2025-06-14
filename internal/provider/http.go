package provider

import (
	"auth-service/internal/config"
	"auth-service/internal/controller"

	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {

	appController := controller.NewAppController()
	appController.Route(router)
}
