package routes

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"Transaction",
		"POST",
		"/transact",
		transact,
	},
	Route{
		"GetBalance",
		"GET",
		"/getBalance/{accountID}",
		getBalance,
	},
}