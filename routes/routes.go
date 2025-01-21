package routes

import (
    "Service2f/controller"

    "github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
    r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
    r.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
    r.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
    r.HandleFunc("/customers/{id:[0-9]+}", controllers.GetCustomerByID).Methods("GET")

    return r
}
