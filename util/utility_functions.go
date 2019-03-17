/*
 * File: utility_functions.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:16:46 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This are just basic utility functions used through out the entire project.
 * -----
 * Last Modified: Sunday, 17th March 2019 7:51:15 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package util

import (
	"log"
	"strconv"

	"google.golang.org/grpc"
)

/**
 * @brief This dials a gRPC connection
 *
 * @param host string //The IP you want to dial
 * @param msg string //the error you want to display when things go wrong.
 *
 * @return *grpc.ClientConn
 */
func Connect(host string, msg string) *grpc.ClientConn {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	CheckError(err, msg)

	return conn
}

/**
 * @brief This checks if there is an error
 *
 * @param err error //The error you want to check
 * @param msg string //the error you want to display when things go wrong.
 *
 * @return bool
 */
func CheckError(err error, msg string) bool {
	success := true
	if err != nil {
		success = false
		log.Fatal(msg, " : %v", err)
	}
	return success
}

/**
 * @brief This converts int to string
 *
 * @param number int32 //The number you want to convert
 *
 * @return string
 */
func IntToString(number int32) string {
	return strconv.FormatInt(int64(number), 10)
}

/**
 * @brief This converts float to string
 *
 * @param number float64 //The number you want to convert
 *
 * @return string
 */
func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'E', -1, 64)
}
