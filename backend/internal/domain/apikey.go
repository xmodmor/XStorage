package domain

import "time"

type APIKey struct {
	ID          uint   `gorm:"primaryKey"`
	AppID       uint   `gorm:"not null;index"`
	App         App    `gorm:"foreignKey:AppID"`
	AccessKey   string `gorm:"uniqueIndex;not null"`
	SecretKey   string `gorm:"not null"`
	Permissions string `gorm:"type:text;not null;default:'read,write'"`
	CreatedAt   time.Time
}
