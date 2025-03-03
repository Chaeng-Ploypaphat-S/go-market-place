package main

import (
	"time"

	"github.com/gocql/gocql"
)

type InventoryItem struct {
	ID          gocql.UUID
	Name        string
	Description string
	Quantity    int
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SellerID    gocql.UUID
}

func NewInventoryItem(name string, description string, quantity int, price float64, seller_id gocql.UUID) *InventoryItem {
	return &InventoryItem{
		ID:          gocql.TimeUUID(),
		Name:        name,
		Description: description,
		Quantity:    quantity,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		SellerID:    seller_id,
	}
}
