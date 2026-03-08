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

type APIKeyHandler struct {
	apiKeyService *service.APIKeyService
}

func NewAPIKeyHandler(apiKeyService *service.APIKeyService) *APIKeyHandler {
	return &APIKeyHandler{apiKeyService: apiKeyService}
}

func (h *APIKeyHandler) Create(c *gin.Context) {
	appID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid app id")
		return
	}

	var req dto.CreateAPIKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	ownerID := c.GetUint("userID")

	key, err := h.apiKeyService.Create(c.Request.Context(), uint(appID), ownerID, req)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "app not found")
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			response.Error(c, http.StatusForbidden, "FORBIDDEN", "access denied")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to create api key")
		return
	}

	response.Success(c, http.StatusCreated, dto.APIKeyResponse{
		ID:          key.ID,
		AppID:       key.AppID,
		AccessKey:   key.AccessKey,
		SecretKey:   key.SecretKey,
		Permissions: key.Permissions,
		CreatedAt:   key.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *APIKeyHandler) List(c *gin.Context) {
	appID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid app id")
		return
	}

	ownerID := c.GetUint("userID")

	keys, err := h.apiKeyService.ListByApp(c.Request.Context(), uint(appID), ownerID)
	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			response.Error(c, http.StatusForbidden, "FORBIDDEN", "access denied")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to list api keys")
		return
	}

	result := make([]dto.APIKeyResponse, len(keys))
	for i, key := range keys {
		result[i] = dto.APIKeyResponse{
			ID:          key.ID,
			AppID:       key.AppID,
			AccessKey:   key.AccessKey,
			Permissions: key.Permissions,
			CreatedAt:   key.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Success(c, http.StatusOK, result)
}

func (h *APIKeyHandler) Delete(c *gin.Context) {
	appID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid app id")
		return
	}

	keyID, err := strconv.ParseUint(c.Param("keyId"), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_ID", "invalid key id")
		return
	}

	ownerID := c.GetUint("userID")

	if err := h.apiKeyService.Delete(c.Request.Context(), uint(keyID), uint(appID), ownerID); err != nil {
		if errors.Is(err, service.ErrForbidden) {
			response.Error(c, http.StatusForbidden, "FORBIDDEN", "access denied")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to delete api key")
		return
	}

	response.Success(c, http.StatusOK, nil)
}
