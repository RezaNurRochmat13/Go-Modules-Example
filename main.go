package main

import (
	"fmt"
	"learning-gomod/domain/delivery/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Server has started")
	RunServer()
}

func RunServer() {
	routers := echo.New()
	groupingRoutes := routers.Group("/public/api/v1/")

	groupingRoutes.GET("users", handler.GetAllUsers)
	groupingRoutes.GET("users/:UsersParam", handler.GetDetailUsers)

	routers.Use(middleware.Logger())

	routers.Logger.Fatal(routers.Start(":8080"))
}
