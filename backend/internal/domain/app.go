package domain

import "time"

type App struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	OwnerID   uint   `gorm:"not null;index"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
	APIKeys   []APIKey `gorm:"foreignKey:AppID"`
	Buckets   []Bucket `gorm:"foreignKey:AppID"`
}
