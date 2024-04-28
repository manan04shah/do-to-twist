package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `gorm:"size:255;not null;" json:"title"`
	Body   string `gorm:"not null;" json:"body"`
	UserID uint   `json:"user_id"` // New field to store the associated user's ID
}
