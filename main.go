package main

import (
	"fmt"
	"learning-gomod/domain/delivery/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server has started")
	RunServer()
}

func RunServer() {
	routers := gin.Default()
	groupingRoutes := routers.Group("/public/api/v1/")
	{
		groupingRoutes.GET("users", handler.GetAllUserHandler)
		groupingRoutes.GET("users/:UsersParam", handler.GetDetailUserHandler)
		groupingRoutes.POST("users", handler.CreateNewUserHandler)
		groupingRoutes.PUT("users/:UsersParam", handler.UpdateUserHandler)
	}

	routers.Run(":8080")

}
