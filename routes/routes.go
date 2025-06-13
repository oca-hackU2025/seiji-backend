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
}
