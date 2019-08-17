package main

import (
	"github.com/rs/cors"
	"hades/router"
	"log"
	"net/http"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	// create router and start listen on port 8000
	newRouter := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", setupGlobalMiddleware(newRouter)))
}
