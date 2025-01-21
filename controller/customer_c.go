package controllers

import (
	"encoding/json"
	"net/http"
	"Service2f/config"
	"Service2f/models"

	"github.com/gorilla/mux"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO customer (name, email) VALUES (?, ?)"
	result, err := config.Database.Exec(query, customer.Name, customer.Email)
	if err != nil {
		http.Error(w, "Error creating customer", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	customer.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, email FROM customer"
	rows, err := config.Database.Query(query)
	if err != nil {
		http.Error(w, "Error fetching customers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email); err != nil {
			http.Error(w, "Error scanning customer row", http.StatusInternalServerError)
			return
		}
		customers = append(customers, customer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var customer models.Customer
	query := "SELECT id, name, email FROM customer WHERE id = ?"
	err := config.Database.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Email)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
