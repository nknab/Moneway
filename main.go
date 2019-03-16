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
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"context"

	"github.com/gorilla/mux"
	bpb "github.com/nknab/Moneway/api/services/balance/protobuf"
	tpb "github.com/nknab/Moneway/api/services/transaction/protobuf"
	ut "github.com/nknab/Moneway/util"
	"google.golang.org/grpc"
)

type Transaction struct {
	AccountID       int32   `json:"accountID"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TransactionType string  `json:"transactionType"`
}

type Balance struct {
	AccountID int32   `json:"accountID"`
	Amount    float64 `json:"amount"`
}

var (
	transactionClient tpb.TransactionClient
	balanceClient     bpb.BalanceClient
	ctx context.Context
)

// func transact(ctx context.Context, transClient tpb.TransactionClient, transaction *tpb.MakeTransaction) (*tpb.MakeTransactionReply, error) {
// 	response, err := transClient.Transact(ctx, &tpb.MakeTransaction{
// 		AccountID:       transaction.AccountID,
// 		Description:     transaction.Description,
// 		Amount:          transaction.Amount,
// 		Currency:        transaction.Currency,
// 		TransactionType: transaction.TransactionType,
// 	})
// 	ut.CheckError(err, "Tans")
// 	return response, nil
// }

// func getBalance(ctx context.Context, balClient bpb.BalanceClient, data *bpb.GetBalanceRequest) (*bpb.GetBalanceReply, error) {
// 	response, err := balClient.GetBalance(ctx, &bpb.GetBalanceRequest{
// 		AccountID: data.AccountID,
// 	})
// 	ut.CheckError(err, "GetBal")
// 	return response, nil
// }

// func updateBalance(ctx context.Context, balClient bpb.BalanceClient, data *bpb.UpdateBalanceRequest) (*bpb.UpdateBalanceReply, error) {
// 	response, err := balClient.GetBalance(ctx, &bpb.UpdateBalanceRequest{
// 		AccountID: data.AccountID,
// 		Amount:    data.Amount,
// 	})
// 	ut.CheckError(err, "GetBal")
// 	return response, nil
// }

func transact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transaction Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	request := &tpb.MakeTransaction{
		AccountID:       transaction.AccountID,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
		Currency:        transaction.Currency,
		TransactionType: transaction.TransactionType,
	}

	response, _ := transactionClient.Transact(ctx, request)
	_ = json.NewEncoder(w).Encode(response)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	accountID, _ := strconv.ParseInt(params["accountID"], 10, 64)

	request := &bpb.GetBalanceRequest{
		AccountID: int32(accountID),
	}

	response, _ := balanceClient.GetBalance(ctx, request)
	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	transactConn, err := grpc.Dial(":8888", grpc.WithInsecure())
	ut.CheckError(err, "TransConn")
	transactionClient = tpb.NewTransactionClient(transactConn)

	balConn, err := grpc.Dial(":8889", grpc.WithInsecure())
	ut.CheckError(err, "BalConn")
	balanceClient = bpb.NewBalanceClient(balConn)

	router := mux.NewRouter()

	router.HandleFunc("/transact", transact).Methods("POST")
	router.HandleFunc("/transact/{accountID}", getBalance).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
