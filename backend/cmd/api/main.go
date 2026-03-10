package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/xmodmor/XStorage/backend/internal/config"
	"github.com/xmodmor/XStorage/backend/internal/database"
	"github.com/xmodmor/XStorage/backend/internal/handler"
	"github.com/xmodmor/XStorage/backend/internal/repository"
	"github.com/xmodmor/XStorage/backend/internal/router"
	"github.com/xmodmor/XStorage/backend/internal/service"
	"github.com/xmodmor/XStorage/backend/internal/storage"
	"github.com/xmodmor/XStorage/backend/seed"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg.DatabaseURL)
	database.RunMigrations(db, "migrations")

	seed.Run(db)

	// Storage
	fs := storage.NewFilesystemStorage(cfg.StoragePath)

	// Repositories
	userRepo := repository.NewUserRepository(db)
	appRepo := repository.NewAppRepository(db)
	apiKeyRepo := repository.NewAPIKeyRepository(db)
	bucketRepo := repository.NewBucketRepository(db)
	objectRepo := repository.NewObjectRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	userService := service.NewUserService(userRepo)
	appService := service.NewAppService(appRepo)
	apiKeyService := service.NewAPIKeyService(apiKeyRepo, appRepo)
	bucketService := service.NewBucketService(bucketRepo)
	objectService := service.NewObjectService(objectRepo, bucketRepo, fs)

	// Handlers
	handlers := router.Handlers{
		Auth:   handler.NewAuthHandler(authService),
		User:   handler.NewUserHandler(userService),
		App:    handler.NewAppHandler(appService),
		APIKey: handler.NewAPIKeyHandler(apiKeyService),
		Bucket: handler.NewBucketHandler(bucketService),
		Object: handler.NewObjectHandler(objectService),
	}

	r := gin.Default()
	router.Setup(r, handlers, cfg.JWTSecret, apiKeyRepo)

	log.Printf("starting server on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
