package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"order-service/publisher"
)

type Order struct {
	OrderID string `json:"order_id"`
	User    string `json:"user"`
	Item    string `json:"item"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	orderID := r.URL.Query().Get("order_id")
	user := r.URL.Query().Get("user")
	item := r.URL.Query().Get("item")

	// Validate required parameters
	if orderID == "" || user == "" || item == "" {
		http.Error(w, "Missing required query parameters: order_id, user, item", http.StatusBadRequest)
		return
	}

	// Create the order
	order := Order{
		OrderID: orderID,
		User:    user,
		Item:    item,
	}

	// Marshal the order to JSON
	body, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal order", http.StatusInternalServerError)
		return
	}

	// Publish the order
	err = publisher.Publish(body)
	if err != nil {
		http.Error(w, "Failed to publish order", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Order sent to queue: %s", order.OrderID)
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/create-order", createOrderHandler)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	failOnError(err, "Failed to start server")
}
