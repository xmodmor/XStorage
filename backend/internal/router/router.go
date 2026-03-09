package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/handler"
	"github.com/xmodmor/XStorage/backend/internal/middleware"
	"github.com/xmodmor/XStorage/backend/internal/repository"
)

type Handlers struct {
	Auth   *handler.AuthHandler
	App    *handler.AppHandler
	APIKey *handler.APIKeyHandler
	Bucket *handler.BucketHandler
	Object *handler.ObjectHandler
}

func Setup(r *gin.Engine, h Handlers, jwtSecret string, apiKeyRepo repository.APIKeyRepository) {
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit(100, 1*time.Minute))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")

	// Public
	v1.POST("/auth/login", h.Auth.Login)

	// Dashboard routes (JWT auth)
	dashboard := v1.Group("")
	dashboard.Use(middleware.JWTAuth(jwtSecret))
	{
		dashboard.GET("/auth/me", h.Auth.Me)

		dashboard.POST("/apps", h.App.Create)
		dashboard.GET("/apps", h.App.List)
		dashboard.GET("/apps/:id", h.App.GetByID)
		dashboard.DELETE("/apps/:id", h.App.Delete)

		dashboard.POST("/apps/:id/keys", h.APIKey.Create)
		dashboard.GET("/apps/:id/keys", h.APIKey.List)
		dashboard.DELETE("/apps/:id/keys/:keyId", h.APIKey.Delete)
	}

	// Storage routes (API key auth)
	storage := v1.Group("")
	storage.Use(middleware.APIKeyAuth(apiKeyRepo))
	{
		storage.POST("/buckets", h.Bucket.Create)
		storage.GET("/buckets", h.Bucket.List)
		storage.DELETE("/buckets/:bucket", h.Bucket.Delete)

		storage.PUT("/buckets/:bucket/objects/*key", h.Object.Upload)
		storage.GET("/buckets/:bucket/objects/*key", h.Object.Download)
		storage.DELETE("/buckets/:bucket/objects/*key", h.Object.Delete)
		storage.GET("/buckets/:bucket/objects", h.Object.List)
	}
}
