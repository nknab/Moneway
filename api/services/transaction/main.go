/*
 * File: transaction.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:01:29 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief:
 * -----
 * Last Modified: Friday, 15th March 2019 7:02:16 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	//"context"
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/nknab/Moneway/api/services/balance/pb"
	"github.com/nknab/Moneway/api/services/transaction/pb"
	db "github.com/nknab/Moneway/database"
	ut "github.com/nknab/Moneway/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var balanceClient balanceService.BalanceClient

type server struct{}

func (s *server) Transact(ctx context.Context, transaction *transactionService.MakeTransaction) (*transactionService.MakeTransactionReply, error) {
	//ctx, stop := context.WithCancel(context.Background())
	//defer stop()
	fmt.Println("IN Transact")
	//Initializing the Database package
	db.Init()

	//This value will comes from the balance service.
	tempValue, _ := balanceClient.GetBalance(context.Background(), &balanceService.GetBalanceRequest{
		AccountID: transaction.AccountID,
	})

	//Converting the Value to a float32
	value := tempValue.Amount
	var oldBalance = math.Round(value*100) / 100
	var newBalance float64

	// Checking If It is a Debit Or A Credit
	if transaction.TransactionType == "DEBIT" {
		newBalance = oldBalance - transaction.Amount
	} else {
		newBalance = oldBalance + transaction.Amount
	}

	// Data for entry
	table := "transactions"
	columns := []string{"account_id", "description", "amount", "old_balance", "new_balance", "currency", "transaction_type"}
	values := []string{ut.IntToString(transaction.AccountID), transaction.Description, ut.FloatToString(transaction.Amount), ut.FloatToString(oldBalance), ut.FloatToString(newBalance), transaction.Currency, transaction.TransactionType}

	funcState := false
	stateStatus := "Status: Transaction Was Unsuccessful."
	// Checking if there is enough Funds in the account before performing transaction
	if newBalance >= 0.0 {
		db.Insert(ctx, table, columns, values)
		//Updating the balance via the balance service.
		state, _ := balanceClient.UpdateBalance(context.Background(), &balanceService.UpdateBalanceRequest{
			AccountID: transaction.AccountID,
			Amount:    newBalance,
		})
		funcState = state.Success
	}

	if funcState == true {
		stateStatus = "Status: Transaction Was Successful. Old Balance was: " + fmt.Sprintf("%.2f", oldBalance) + ". Current Balance is: " + fmt.Sprintf("%.2f", newBalance)
	} else if newBalance < 0.0 {
		stateStatus = "Status: Transaction Was Unsuccessful. Not Enough Funds in Account"
	}

	response := &transactionService.MakeTransactionReply{
		Success: true,
		Id:      2,
		Msg:     stateStatus,
	}

	return response, nil
}

func main() {

	balanceClient = balanceService.NewBalanceClient(ut.Connect(":8082", "Could not connect to :8082"))

	transactionSrv, err := net.Listen("tcp", ":8081")
	ut.CheckError(err, "Could not listen to TranactionServe on :8081")

	srv := grpc.NewServer()
	transactionService.RegisterTransactionServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(transactionSrv); err != nil {
		log.Fatalf("Could not connect to: %v", err)
	}
}
