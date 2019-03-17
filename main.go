/*
 * File: main.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 8:32:15 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This initializes the main server for request to come in from the user
 * -----
 * Last Modified: Sunday, 17th March 2019 7:44:39 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	rt "github.com/nknab/Moneway/routes"
	ut "github.com/nknab/Moneway/util"
)

func main() {
	var port string

	//Decalring and initializing the various routes
	router := rt.Router()
	rt.Run()

	//Reading in the configuration variables
	err := godotenv.Load("config/config.env")
	ut.CheckError(err, "Error loading config.env file")
	port = ":" + os.Getenv("MAIN_PORT")

	//Serving and Listening to the main server
	log.Fatal(http.ListenAndServe(port, router))

}
