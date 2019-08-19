package main

import (
	"fmt"
	"github.com/rs/cors"
	"hades/router"
	"log"
	"net/http"
	"os"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	PORT := os.Getenv("PORT")
	fmt.Println("PORT is", PORT);
	if PORT == "" {
		PORT = "8080"
	}

	// create router and start listen on port
	newRouter := router.NewRouter()
	fmt.Println("Server starts on PORT", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, setupGlobalMiddleware(newRouter)))
}
