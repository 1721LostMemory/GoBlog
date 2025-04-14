package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Author    string `gorm:"foreignKey:author"`
	Title     string
	Content   string `gorm:"type:text"`
	ImagePath string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
