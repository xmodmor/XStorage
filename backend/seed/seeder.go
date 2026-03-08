package seed

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/xmodmor/XStorage/backend/internal/domain"
)

func Run(db *gorm.DB) {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		log.Println("seed: skipped (users already exist)")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("seed: failed to hash password: %v", err)
	}

	admin := domain.User{
		Email:        "admin@xstorage.local",
		PasswordHash: string(hash),
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatalf("seed: failed to create admin user: %v", err)
	}

	log.Println("seed: admin user created (admin@xstorage.local / admin123)")
}
