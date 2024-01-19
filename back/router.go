package main

import (
	"back/common"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controllers []common.Controller) *gin.Engine {
	router := gin.Default()
	
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
	
	return router
}