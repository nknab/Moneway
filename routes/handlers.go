/*
 * File: handlers.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Sunday, 17th March 2019 7:27:07 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This is to handle all route function calls
 * -----
 * Last Modified: Sunday, 17th March 2019 8:07:28 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package routes

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	balanceService "github.com/nknab/Moneway/api/services/balance/pb"
	transactionService "github.com/nknab/Moneway/api/services/transaction/pb"
	ut "github.com/nknab/Moneway/util"
)

// The structure of out balance reply services.
type Balance struct {
	AccountID int32   `json:"accountID"`
	Amount    float64 `json:"amount"`
}

//Various variables need for the class.
var (
	transactClient transactionService.TransactionClient
	balanceClient  balanceService.BalanceClient
	config         = "config/config.env"
)

/**
 * @brief This is the method call by our transaction route to make transactions.
 *
 * @param w http.ResponseWriter
 * @param r *http.Request
 *
 */
func transact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Getting the various details from the post request
	aid, _ := strconv.ParseInt(r.FormValue("AccountID"), 10, 64)
	amt, _ := strconv.ParseFloat(r.FormValue("Amount"), 64)
	request := transactionService.MakeTransaction{
		AccountID:       int32(aid),
		Description:     r.FormValue("Description"),
		Amount:          amt,
		Currency:        r.FormValue("Currency"),
		TransactionType: r.FormValue("TransactionType"),
	}

	//Pushing a transaction request to the transaction service.
	var response, _ = transactClient.Transact(context.Background(), &request)

	//Putting the reply in a json object for the user
	_ = json.NewEncoder(w).Encode(response)
}

/**
 * @brief This is the method call by our get balance route to get the balance.
 *
 * @param w http.ResponseWriter
 * @param r *http.Request
 *
 */
func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//Getting the ID from the Get request
	params := mux.Vars(r)
	accountID, _ := strconv.ParseInt(params["accountID"], 10, 64)
	request := balanceService.GetBalanceRequest{
		AccountID: int32(accountID),
	}

	//Pushing a get balance request to the balance service.
	response, err := balanceClient.GetBalance(context.Background(), &request)
	ut.CheckError(err, "There was an error when getting the balance")

	//Putting the reply in a json object for the user
	var balance Balance
	balance.AccountID = int32(accountID)
	balance.Amount = math.Round(response.Amount*100) / 100
	_ = json.NewEncoder(w).Encode(balance)
}

func Run() {
	var bport string
	var tport string

	//Loading in the environmant variables
	err := godotenv.Load(config)
	ut.CheckError(err, "Error loading config.env file")
	bport = ":" + os.Getenv("BALANCE_PORT")
	tport = ":" + os.Getenv("TRANSACTION_PORT")

	//Instantating the various services.
	transactClient = transactionService.NewTransactionClient(ut.Connect(tport, "Could not connect to "+tport))
	balanceClient = balanceService.NewBalanceClient(ut.Connect(bport, "Could not connect to "+bport))
}
