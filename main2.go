package main

import (
	"fmt"
	"log"
	"net/http"
	"Service2f/config"
	"Service2f/routes"
)

func main() {
	config.InitDB()
	defer config.Database.Close()

	r := routes.SetupRoutes()

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
