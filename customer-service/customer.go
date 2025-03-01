package main

import (
	"time"
)

type Customer struct {
	ID             int
	CustomerType   int
	Name           string
	Email          string
	RegisteredDate time.Time
	LastUpdated    time.Time
}

func NewCustomer(id int, name string, email string) *Customer {
	return &Customer{
		ID:             id,
		Name:           name,
		Email:          email,
		RegisteredDate: time.Now(),
		LastUpdated:    time.Now(),
	}
}

func (c *Customer) UpdateCustomer(name string, email string, customerType int) {
	c.Name = name
	c.Email = email
	c.CustomerType = customerType
	c.LastUpdated = time.Now()
}
