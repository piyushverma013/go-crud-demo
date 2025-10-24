package router

import (
	"github.com/gin-gonic/gin"
	service "github.com/piyushverma013/simple-go-crud/Service"
	"github.com/piyushverma013/simple-go-crud/handler"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	userGroup := r.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/", userHandler.GetUsers)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}
