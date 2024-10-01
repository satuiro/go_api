package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	// Create order
	fmt.Println("Create an order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	// Create order
	fmt.Println("List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	// Create order
	fmt.Println("Get order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	// Create order
	fmt.Println("Update an order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	// Create order
	fmt.Println("Delete an order by ID")
}
