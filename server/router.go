package server

import (
	"github.com/gin-gonic/gin"
	"github.com/printSANO/go-gin-example/handlers"
)

// 라우터 설정
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		// User routes
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("/:id", handler.UserHandler.GetUser)
			userGroup.POST("/", handler.UserHandler.CreateUser)
		}

		// Post routes
		postGroup := apiGroup.Group("/posts")
		{
			postGroup.GET("/:id", handler.PostHandler.GetPost)
			postGroup.POST("/", handler.PostHandler.CreatePost)
		}
	}
	return router
}
