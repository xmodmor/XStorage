package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/response"
	"github.com/xmodmor/XStorage/backend/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	search := c.Query("search")

	result, err := h.userService.List(c.Request.Context(), page, limit, search)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to list users")
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	user, err := h.userService.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrConflict) {
			response.Error(c, http.StatusConflict, "EMAIL_TAKEN", "email already in use")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to create user")
		return
	}

	response.Success(c, http.StatusCreated, user)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid user id")
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", "user not found")
		return
	}

	response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid user id")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	user, err := h.userService.Update(c.Request.Context(), uint(id), req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "user not found")
			return
		}
		if errors.Is(err, service.ErrConflict) {
			response.Error(c, http.StatusConflict, "EMAIL_TAKEN", "email already in use")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to update user")
		return
	}

	response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid user id")
		return
	}

	if err := h.userService.Delete(c.Request.Context(), uint(id)); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "user not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to delete user")
		return
	}

	response.Success(c, http.StatusOK, nil)
}
