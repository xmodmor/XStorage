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

type AppHandler struct {
	appService *service.AppService
}

func NewAppHandler(appService *service.AppService) *AppHandler {
	return &AppHandler{appService: appService}
}

func (h *AppHandler) Create(c *gin.Context) {
	var req dto.CreateAppRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	ownerID := c.GetUint("userID")

	app, err := h.appService.Create(c.Request.Context(), ownerID, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to create app")
		return
	}

	response.Success(c, http.StatusCreated, dto.AppResponse{
		ID:        app.ID,
		Name:      app.Name,
		OwnerID:   app.OwnerID,
		CreatedAt: app.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *AppHandler) List(c *gin.Context) {
	ownerID := c.GetUint("userID")

	apps, err := h.appService.ListByOwner(c.Request.Context(), ownerID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to list apps")
		return
	}

	result := make([]dto.AppResponse, len(apps))
	for i, app := range apps {
		result[i] = dto.AppResponse{
			ID:        app.ID,
			Name:      app.Name,
			OwnerID:   app.OwnerID,
			CreatedAt: app.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Success(c, http.StatusOK, result)
}

func (h *AppHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid app id")
		return
	}

	app, err := h.appService.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "NOT_FOUND", "app not found")
		return
	}

	response.Success(c, http.StatusOK, dto.AppResponse{
		ID:        app.ID,
		Name:      app.Name,
		OwnerID:   app.OwnerID,
		CreatedAt: app.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *AppHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid app id")
		return
	}

	ownerID := c.GetUint("userID")

	if err := h.appService.Delete(c.Request.Context(), uint(id), ownerID); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "app not found")
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			response.Error(c, http.StatusForbidden, "FORBIDDEN", "access denied")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to delete app")
		return
	}

	response.Success(c, http.StatusOK, nil)
}
