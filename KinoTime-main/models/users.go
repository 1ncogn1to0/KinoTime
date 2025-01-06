package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:created_at" json:"updated_at"`
}
