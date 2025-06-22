package main

import (
	"time"
)

type Payment struct {
	ID            int
	CustomerID    int
	Amount        float64
	PaymentDate   time.Time
	PaymentMethod string
}

func NewPayment(customerID int, amount float64, paymentMethod string) *Payment {
	return &Payment{
		CustomerID:    customerID,
		Amount:        amount,
		PaymentDate:   time.Now(),
		PaymentMethod: paymentMethod,
	}
}
