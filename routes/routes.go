package routes

import (
	"github.com/KENKUN-1031/seiji-backend/controller"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	auth := r.Group("/")
	{
		auth.GET("/test", controller.TestControl)
	}

	api := r.Group("/api")
	{
		api.GET("/politicians/generate", controller.GetRandomSeijika)
		api.POST("/auth/login", controller.Login)
	}
}
