package routes

import (
	"github.com/KENKUN-1031/seiji-backend/controller"
	"github.com/KENKUN-1031/seiji-backend/middleware"
	"github.com/gin-gonic/gin"
)

func DefineRoutes(r gin.IRouter) {
	auth := r.Group("/")
	{
		auth.GET("/test", controller.TestControl)
	}

	api := r.Group("/api")
	{
		api.POST("/auth/login", controller.Login)
	}

	private := r.Group("/api/private")
	private.Use(middleware.JWTAuthMiddleware()) // jwt認証が通った場合のみrouteを通す
	{
		private.GET("/politicians/generate", controller.GetRandomSeijika)
		private.POST("/likes", controller.CreateLike)
		private.GET("/likes", controller.GetLikedSeijikaList)
	}
}
