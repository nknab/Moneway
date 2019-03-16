/*
 * File: utility_functions.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:16:46 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief:
 * -----
 * Last Modified: Friday, 15th March 2019 7:16:59 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package util

import (
	"fmt"
	"log"
)

func CheckError(err error, msg string) bool {
	success := true
	if err != nil {
		success = false
		log.Fatal(msg, ": %v", err)
	}
	return success
}

func IntToString(number int32) string {
	return fmt.Sprintf("%f", number)
}

func FloatToString(number float64) string {
	return fmt.Sprintf("%f", number)
}
