
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    // "math/rand"
    "net/http"
    // "strconv"
    // "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

const (
    DBHost  = "127.0.0.1"
    DBUser  = "root"
    DBPass  = "Kush@123456"
    DBDbase = "pro"
    PORT    = ":3000"
)

var database *sql.DB

func initDB() {
    dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
    db, err := sql.Open("mysql", dbConn)
    if err != nil {
        log.Fatalf("Database connection error: %v", err)
    }

    database = db
   
    if err := database.Ping(); err != nil {
        log.Fatalf("Database ping error: %v", err)
    }
    log.Println("Database connected successfully!")
}

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO product (name) VALUES (?)"
	result, err := database.Exec(query, product.Name)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	product.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name FROM product"
	rows, err := database.Query(query)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			http.Error(w, "Error scanning product row", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO customer (name, email) VALUES (?, ?)"
	result, err := database.Exec(query, customer.Name, customer.Email)
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
	rows, err := database.Query(query)
	if err != nil {
		http.Error(w, "Error fetching customers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email); err != nil {
			http.Error(w, "Error scanning customer row", http.StatusInternalServerError)
			return
		}
		customers = append(customers, customer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	query := "SELECT id, name FROM product WHERE id = ?"
	err := database.QueryRow(query, id).Scan(&product.ID, &product.Name)
	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var customer Customer
	query := "SELECT id, name, email FROM customer WHERE id = ?"
	err := database.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Email)
	if err == sql.ErrNoRows {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching customer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func main() {

	initDB()
	defer database.Close()

	r := mux.NewRouter()

	
	r.HandleFunc("/products", CreateProduct).Methods("POST")
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/customers", CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", GetProductByID).Methods("GET")
	r.HandleFunc("/customers/{id:[0-9]+}", GetCustomerByID).Methods("GET")


	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
