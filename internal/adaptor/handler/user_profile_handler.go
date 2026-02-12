package handler

import (
	"net/http"
	"time"

	"github.com/NuttayotSukkum/user-profile/internal/core/port"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc port.UserProfileSvc
}

func NewUserHandler(svc port.UserProfileSvc) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req port.UserProFileRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Code: CodeBadRequest,
			Time: time.Now().Format(time.RFC3339),
			Data: "invalid request body",
		})
	}

	if req.Email == "" || req.FirstName == "" || req.LastName == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, MessageResponse{
			Code: CodeBadRequest,
			Time: time.Now().Format(time.RFC3339),
			Data: "email or name or password is empty",
		})
	}

	if err := h.svc.CreateUserProfile(req); err != nil {
		return c.JSON(http.StatusInternalServerError, MessageResponse{
			Code: CodeInternalError,
			Time: time.Now().Format(time.RFC3339),
			Data: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, MessageResponse{
		Code: CodeSuccess,
		Time: time.Now().Format(time.RFC3339),
		Data: "user created",
	})
}
