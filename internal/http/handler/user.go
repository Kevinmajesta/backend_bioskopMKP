package handler

import (
	"net/http"

	"Kevinmajesta/backend_bioskopMKP/internal/http/binder"
	"Kevinmajesta/backend_bioskopMKP/internal/service"
	"Kevinmajesta/backend_bioskopMKP/pkg/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c echo.Context) error {
	input := binder.UserRegisterRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	user, err := h.userService.Register(input.Name, input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "register success", user))
}

func (h *UserHandler) Login(c echo.Context) error {
	input := binder.UserLoginRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "ada kesalahan input"))
	}

	token, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "login success", binder.UserLoginResponse{
		Token: token,
	}))
}
