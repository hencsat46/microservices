package handler

import (
	"microservices/recordsManager/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	CreateUser(models.User) error
	ReadUser(int) (*models.User, error)
	UpdateUser(models.User) error
	DeleteUser(int) error
	CreateRecord(models.RecordModel) error
	ReadRecord(int) (*models.RecordModel, error)
	UpdateRecord(models.RecordModel) error
	DeleteRecord(int) error
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Routes(e *echo.Echo) {
	e.POST("/createuser", h.CreateUser)
	e.GET("/getuser", h.ReadUser)
	e.PUT("/updateuser", h.UpdateUser)
	e.DELETE("/deleteuser", h.DeleteUser)
	e.POST("/createrecord", h.CreateRecord)
	e.GET("/getrecord", h.ReadRecord)
	e.PUT("/updaterecord", h.UpdateRecord)
	e.DELETE("/deleterecord", h.DeleteRecord)
}

func (h *handler) CreateUser(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.CreateUser(models.User{Username: user.Username, Password: user.Password, Name: user.Name, Surname: user.Surname, Id: user.Id}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign up ok"})
}

func (h *handler) ReadUser(ctx echo.Context) error {
	value := ctx.QueryParam("id")
	intValue, _ := strconv.Atoi(value)

	response, err := h.usecase.ReadUser(intValue)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: response})
}

func (h *handler) UpdateUser(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.UpdateUser(models.User{Username: user.Username, Password: user.Password, Name: user.Name, Surname: user.Surname, Id: user.Id}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Update ok"})
}

func (h *handler) DeleteUser(ctx echo.Context) error {
	var user = userDTO{"", "", "", "", -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.DeleteUser(user.Id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Delete ok"})
}

func (h *handler) CreateRecord(ctx echo.Context) error {
	var user = recordDTO{-1, -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.CreateRecord(models.RecordModel{UserId: user.UserId, RecordId: user.RecordId}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign up ok"})
}

func (h *handler) ReadRecord(ctx echo.Context) error {
	value := ctx.QueryParam("id")
	intValue, _ := strconv.Atoi(value)

	response, err := h.usecase.ReadRecord(intValue)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: response})
}

func (h *handler) UpdateRecord(ctx echo.Context) error {
	var user = recordDTO{-1, -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.UpdateRecord(models.RecordModel{RecordId: user.RecordId, UserId: user.UserId}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Update ok"})
}

func (h *handler) DeleteRecord(ctx echo.Context) error {
	var user = recordDTO{-1, -1}

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Bad json"})
	}

	if err := h.usecase.DeleteRecord(user.UserId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Delete ok"})
}
