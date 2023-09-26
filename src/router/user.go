package router

import (
	"application/src/controller"
	"application/src/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(userService *controller.UserService, e *gin.Engine) {
	r := e.Group("/user")
	r.POST("/users", userService.CreateUser)
	r.GET("/users", middleware.AuthMiddleware(), userService.GetAllUsers)
	r.GET("/users/:userID", middleware.AuthMiddleware(), userService.GetSingleUser)
	r.DELETE("/users/:userID", middleware.AuthMiddleware(), userService.DeleteUser)
}
