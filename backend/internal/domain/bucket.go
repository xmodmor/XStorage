package domain

import "time"

type Bucket struct {
	ID         uint     `gorm:"primaryKey"`
	AppID      uint     `gorm:"not null;index"`
	App        App      `gorm:"foreignKey:AppID"`
	Name       string   `gorm:"not null"`
	Visibility string   `gorm:"not null;default:'private'"`
	CreatedAt  time.Time
	Objects    []Object `gorm:"foreignKey:BucketID"`
}
