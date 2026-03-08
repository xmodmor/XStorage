package domain

import "time"

type Object struct {
	ID          uint   `gorm:"primaryKey"`
	BucketID    uint   `gorm:"not null;index"`
	Bucket      Bucket `gorm:"foreignKey:BucketID"`
	Key         string `gorm:"not null"`
	Size        int64  `gorm:"not null"`
	Mime        string `gorm:"not null"`
	StoragePath string `gorm:"not null"`
	Checksum    string `gorm:"not null"`
	CreatedAt   time.Time
}
