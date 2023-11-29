package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pratikshakuldeep456/receipeBook/routes"
	//"github.com/pratiksha456/receipeBook/routes"
)

func main() {
	// Initialize some sample data
	routes.InitSampleData()

	// Initialize the router
	router := mux.NewRouter()

	// Define API endpoints
	routes.SetupRoutes(router)

	// Start the server
	port := 8080
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server is running on :%d\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
