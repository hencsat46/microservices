package handler

import (
	"microservices/authentication/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	Create(models.User, int) error
	Read(int) (models.User, error)
	Update(models.User, int) error
	Delete(int) error
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Routes(e *echo.Echo) {
	e.POST("/create", h.Create)
	e.GET("/get", h.Read)
	e.PUT("/update", h.Update)
	e.DELETE("/delete", h.Delete)
}

func (h *handler) Create(ctx echo.Context) error {

	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.Create(models.User{Username: user.Username, Password: user.Password, Name: user.Name, Surname: user.Surname}, user.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign up ok"})
}

func (h *handler) Read(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	value := ctx.QueryParam("id")
	intValue, _ := strconv.Atoi(value)

	response, err := h.usecase.Read(intValue)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: response})
}

func (h *handler) Update(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.Update(models.User{Username: user.Username, Password: user.Password, Name: user.Name, Surname: user.Surname}, user.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Update ok"})

}

func (h *handler) Delete(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.Delete(user.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Delete ok"})

}
