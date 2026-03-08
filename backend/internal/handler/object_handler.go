package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/dto"
	"github.com/xmodmor/XStorage/backend/internal/response"
	"github.com/xmodmor/XStorage/backend/internal/service"
)

type ObjectHandler struct {
	objectService *service.ObjectService
}

func NewObjectHandler(objectService *service.ObjectService) *ObjectHandler {
	return &ObjectHandler{objectService: objectService}
}

func (h *ObjectHandler) Upload(c *gin.Context) {
	appID := c.GetUint("appID")
	bucketName := c.Param("bucket")
	key := strings.TrimPrefix(c.Param("key"), "/")

	if key == "" {
		response.Error(c, http.StatusBadRequest, "INVALID_KEY", "object key is required")
		return
	}

	mime := c.ContentType()
	if mime == "" {
		mime = "application/octet-stream"
	}

	object, err := h.objectService.Upload(c.Request.Context(), appID, bucketName, key, mime, c.Request.Body)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "BUCKET_NOT_FOUND", "bucket not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "UPLOAD_FAILED", "failed to upload object")
		return
	}

	response.Success(c, http.StatusCreated, dto.ObjectResponse{
		ID:        object.ID,
		Key:       object.Key,
		Size:      object.Size,
		Mime:      object.Mime,
		Checksum:  object.Checksum,
		CreatedAt: object.CreatedAt.Format("2006-01-02T15:04:05Z"),
	})
}

func (h *ObjectHandler) Download(c *gin.Context) {
	appID := c.GetUint("appID")
	bucketName := c.Param("bucket")
	key := strings.TrimPrefix(c.Param("key"), "/")

	reader, object, err := h.objectService.Download(c.Request.Context(), appID, bucketName, key)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "object not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "DOWNLOAD_FAILED", "failed to download object")
		return
	}
	defer reader.Close()

	c.Header("Content-Type", object.Mime)
	c.Header("Content-Length", fmt.Sprintf("%d", object.Size))
	c.Header("ETag", object.Checksum)
	c.DataFromReader(http.StatusOK, object.Size, object.Mime, reader, nil)
}

func (h *ObjectHandler) Delete(c *gin.Context) {
	appID := c.GetUint("appID")
	bucketName := c.Param("bucket")
	key := strings.TrimPrefix(c.Param("key"), "/")

	if err := h.objectService.Delete(c.Request.Context(), appID, bucketName, key); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "NOT_FOUND", "object not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "DELETE_FAILED", "failed to delete object")
		return
	}

	response.Success(c, http.StatusOK, nil)
}

func (h *ObjectHandler) List(c *gin.Context) {
	appID := c.GetUint("appID")
	bucketName := c.Param("bucket")

	var params dto.ListObjectsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}

	objects, total, err := h.objectService.List(c.Request.Context(), appID, bucketName, params.Page, params.PerPage)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			response.Error(c, http.StatusNotFound, "BUCKET_NOT_FOUND", "bucket not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to list objects")
		return
	}

	result := make([]dto.ObjectResponse, len(objects))
	for i, obj := range objects {
		result[i] = dto.ObjectResponse{
			ID:        obj.ID,
			Key:       obj.Key,
			Size:      obj.Size,
			Mime:      obj.Mime,
			Checksum:  obj.Checksum,
			CreatedAt: obj.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	response.Paginated(c, result, response.Meta{
		Page:    params.Page,
		PerPage: params.PerPage,
		Total:   total,
	})
}
