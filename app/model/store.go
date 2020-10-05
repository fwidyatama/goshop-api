package model

import (
	"time"
)

type Store struct {
	ID        uint `json:"id" gorm:"primaryKey`
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
	Items []Item `gorm:"many2many:item_store;"`
}

type StoreJson struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


