package model

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
	Items []Item `gorm:"many2many:store_items;"`
}

type StoreJson struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StoreItem struct {
	ID        int `json:"id"`
	StoreID   int `json:"store_id"`
	ItemID    int `json:"item_id"`
	Quantity  int `json:"quantity"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-"`
}
