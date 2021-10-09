package routes

import (
	"github.com/TheNeoCarvalho/go-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigueRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBooks)
		}
	}

	return router
}
