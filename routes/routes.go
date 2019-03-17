/*
 * File: routes.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Sunday, 17th March 2019 10:48:02 AM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This is where all Routes in the projects are defined
 * -----
 * Last Modified: Sunday, 17th March 2019 7:52:54 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package routes

import (
	"net/http"
)

//Structure of a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//A slice of routes
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
