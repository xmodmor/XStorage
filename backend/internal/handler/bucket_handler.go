package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/response"
	"github.com/xmodmor/XStorage/backend/internal/service"
)

type BucketHandler struct {
	bucketService *service.BucketService
}

func NewBucketHandler(bucketService *service.BucketService) *BucketHandler {
	return &BucketHandler{bucketService: bucketService}
}

func (h *BucketHandler) Create(c *gin.Context) {
	var req dto.CreateBucketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	appID := c.GetUint("appID")

	bucket, err := h.bucketService.Create(c.Request.Context(), appID, req)
	if err != nil {
		if errors.Is(err, service.ErrConflict) {
			response.Error(c, http.StatusConflict, "BUCKET_EXISTS", "bucket with this name already exists")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to create bucket")
		return
	}

	response.Success(c, http.StatusCreated, dto.BucketResponse{
		ID:         bucket.ID,
		Name:       bucket.Name,
		Visibility: bucket.Visibility,
		CreatedAt:  bucket.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *BucketHandler) List(c *gin.Context) {
	appID := c.GetUint("appID")

	buckets, err := h.bucketService.ListByApp(c.Request.Context(), appID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to list buckets")
		return
	}

	result := make([]dto.BucketResponse, len(buckets))
	for i, b := range buckets {
		result[i] = dto.BucketResponse{
			ID:         b.ID,
			Name:       b.Name,
			Visibility: b.Visibility,
			CreatedAt:  b.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Success(c, http.StatusOK, result)
}

func (h *BucketHandler) Delete(c *gin.Context) {
	appID := c.GetUint("appID")
	bucketName := c.Param("bucket")

	if err := h.bucketService.Delete(c.Request.Context(), appID, bucketName); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "bucket not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to delete bucket")
		return
	}

	response.Success(c, http.StatusOK, nil)
}
