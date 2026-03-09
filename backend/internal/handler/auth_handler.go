package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/response"
	"github.com/xmodmor/XStorage/backend/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	res, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			response.Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid email or password")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "something went wrong")
		return
	}

	response.Success(c, http.StatusOK, res)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "UNAUTHORIZED", "user not found in context")
		return
	}

	user, err := h.authService.GetCurrentUser(c.Request.Context(), userID.(uint))
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "user not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "something went wrong")
		return
	}

	response.Success(c, http.StatusOK, user)
}
