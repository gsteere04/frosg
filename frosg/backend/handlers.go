package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Get all customer
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	DB.Find(&customers)
	json.NewEncoder(w).Encode(&customers)
}

// Create a new customer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	DB.Create(&customers)
	json.NewEncoder(w).Encode(&customers)
}

// Get a single customer by ID
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer Customer
	DB.Find(&customer, params["id"])
	json.NewEncoder(w).Encode(&customer)
}

// Get all transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []Transaction
	DB.Find(&transactions)
	json.NewEncoder(w).Encode(&transactions)
}

// Create a new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions []Transaction
	DB.Create(&transactions)
	json.NewEncoder(w).Encode(&transactions)
}

//
