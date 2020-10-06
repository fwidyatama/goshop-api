package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty" `
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	OrderDate time.Time          `json:"order_date" bson:"order_date"`
	OrderItem []OrderItem        `bson:"order_item" json:"order_item"`
}

type OrderItem struct {
	ItemId    int    `json:"name" bson:"item_id"`
	ItemName  string `json:"item_name" bson:"item_name"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	StoreId   int    `json:"store_id" bson:"store_id"`
	StoreName string `json:"store_name" bson:"store_name"`
}

func NewOrder() *Order {
	return &Order{
		OrderDate: time.Now(),
	}
}
