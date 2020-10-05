package model

import (
	"time"
)

type Item struct {
	ID          uint    `json:"id" gorm:"primaryKey`
	Name        string  `json:"name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type ItemJson struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
