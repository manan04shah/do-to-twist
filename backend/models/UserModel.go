package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Notes    []Note `gorm:"foreignKey:UserID" json:"notes"` // New field to store associated notes
}
