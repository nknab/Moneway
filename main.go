/*
 * File: main.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 10:26:00 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief:
 * -----
 * Last Modified: Friday, 15th March 2019 10:26:06 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	rt "github.com/nknab/Moneway/routes"
	"log"
	"net/http"
)




func main() {
	router := rt.Router()

	log.Fatal(http.ListenAndServe(":8000", router))

}
