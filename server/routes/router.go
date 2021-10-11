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
			books.PUT("/", controllers.UpdateBook)
			books.POST("/", controllers.CreateBook)
			books.GET("/:id", controllers.ShowBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
