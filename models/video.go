package models

import "time"

type Video struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	URL         string    `gorm:"primaryKey" json:"url"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Duration    uint      `gorm:"not null" json:"duration"`
	Resolution  string    `gorm:"not null" json:"resolution"`
	UploadDate  time.Time `gorm:"not null" json:"upload_date"`
}
