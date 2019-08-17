package router

import (
	"fmt"
	"github.com/gorilla/mux"
)

// NewRouter builds and returns a new router from routes
func NewRouter() *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()

	fmt.Println("ROUTES IN OPERATION")
	fmt.Println(".............................")


	for _, route := range routes {
		fmt.Println(route.Name, "\t", route.Method, "\t", "/api/v1" + route.Pattern)
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
