package main

import (
	"microservices/libraryManager/internal/controller"
	"microservices/libraryManager/internal/gateway"
	"microservices/libraryManager/internal/handler"
	"microservices/libraryManager/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	repo := repository.NewRepository()
	usecase := controller.NewUsecase(repo, *gateway.New())
	handler := handler.NewHandler(usecase)
	handler.Routes(e)

	e.Start(":3001")

}
