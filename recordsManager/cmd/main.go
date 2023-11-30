package main

import (
	"microservices/recordsManager/internal/controller"
	"microservices/recordsManager/internal/gateway"
	"microservices/recordsManager/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	usecase := controller.NewUsecase(gateway.New())
	handler := handler.NewHandler(usecase)
	handler.Routes(e)

	e.Start(":3002")
}
