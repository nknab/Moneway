/*
 * File: router.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Sunday, 17th March 2019 11:11:21 AM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This basically creates our Routes Instance.
 * -----
 * Last Modified: Sunday, 17th March 2019 7:54:36 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/**
 * @brief This basically logs every users request
 *
 * @param inner http.Handler//The http Handle
 * @param name string //The name of the route
 *
 * @return http.Handler
 */
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

/**
 * @brief This creates an instace of our routes
 *
 * @return *mux.Router
 */
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
