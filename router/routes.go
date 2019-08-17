package router

import (
	handler "hades/handlers"
	"net/http"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"Generate OTP",
		"POST",
		"/otp",
		handler.CreateNewOTP,
	},
	Route{
		"Validate OTP",
		"POST",
		"/otp/validate",
		handler.ValidateOTP,
	},
}
