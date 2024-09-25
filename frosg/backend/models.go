package main

import (
	"gorm.io/gorm"
)

// Customer model
type Customer struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phone_number"`
}

// Transaction model
type Transaction struct {
	gorm.Model
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}

// Migrate the database schema
func migrate() {
	DB.AutoMigrate(&Customer{}, &Transaction{})
}
