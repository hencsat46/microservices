package main

import (
	"microservices/authentication/internal/controller"
	"microservices/authentication/internal/handler"
	"microservices/authentication/internal/repostitory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	repo := repostitory.NewRepository()
	usecase := controller.NewUsecase(repo)
	handler := handler.NewHandler(usecase)
	handler.Routes(e)

	e.Start(":3000")

}
