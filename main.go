package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"hades/router"
	"log"
	"net/http"
	"os"
)

func init(){
	// Loading env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	PORT := os.Getenv("PORT")

	// create router and start listen on port
	newRouter := router.NewRouter()
	fmt.Println("Server starts on PORT", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, setupGlobalMiddleware(newRouter)))
}
