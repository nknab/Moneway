/*
 * File: transaction.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:01:29 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This this Transaction Service class.
 * -----
 * Last Modified: Friday, 15th March 2019 7:02:16 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"net"
	"os"

	"github.com/joho/godotenv"

	balanceService "github.com/nknab/Moneway/api/services/balance/pb"
	transactionService "github.com/nknab/Moneway/api/services/transaction/pb"
	"github.com/nknab/Moneway/database"
	ut "github.com/nknab/Moneway/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	balanceClient balanceService.BalanceClient
	config        = "../../../config/config.env"
	db            *sql.DB
)

type server struct{}

/**
 * @brief This gets a user's balance.
 *
 * @param ctx context.Context //Context
 * @param transaction *transactionService.MakeTransaction // Reference to the MakeTransaction Object
 *
 * @return *transactionService.MakeTransactionReply
 * @return error
 */
func (s *server) Transact(ctx context.Context, transaction *transactionService.MakeTransaction) (*transactionService.MakeTransactionReply, error) {
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

	// Data for transaction
	table := "transactions"
	columns := []string{"account_id", "description", "amount", "old_balance", "new_balance", "currency", "transaction_type"}
	values := []string{ut.IntToString(transaction.AccountID), transaction.Description, ut.FloatToString(transaction.Amount), ut.FloatToString(oldBalance), ut.FloatToString(newBalance), transaction.Currency, transaction.TransactionType}

	funcState := false
	stateStatus := "Status: Transaction Was Unsuccessful."

	var id int32
	// Checking if there is enough Funds in the account before performing transaction
	if newBalance >= 0.0 {
		id, _ = database.Insert(ctx, table, columns, values, db)
		//Updating the balance via the balance service.
		state, _ := balanceClient.UpdateBalance(context.Background(), &balanceService.UpdateBalanceRequest{
			AccountID: transaction.AccountID,
			Amount:    newBalance,
		})
		funcState = state.Success
	} else{
		id = -1
	}

	//Checking to see of the transaction was successful.
	if funcState == true {
		stateStatus = "Status: Transaction Was Successful. Old Balance was: " + fmt.Sprintf("%.2f", oldBalance) + ". Current Balance is: " + fmt.Sprintf("%.2f", newBalance)
	} else if newBalance < 0.0 {
		stateStatus = "Status: Transaction Was Unsuccessful. Not Enough Funds in Account"
	}
	fmt.Println(id)

	response := &transactionService.MakeTransactionReply{
		Success: true,
		Id:      id,
		Msg:     stateStatus,
	}

	return response, nil
}

func main() {
	var bport string
	var tport string

	//Initializing the database
	db = database.Init(config)

	//Reading in the configuration variables
	err := godotenv.Load(config)
	ut.CheckError(err, "Error loading config.env file")
	bport = ":" + os.Getenv("BALANCE_PORT")
	tport = ":" + os.Getenv("TRANSACTION_PORT")

	//Connect to my balance service client.
	balanceClient = balanceService.NewBalanceClient(ut.Connect(bport, "Could not connect to "+bport))

	transactionSrv, err := net.Listen("tcp", tport)
	ut.CheckError(err, "Could not listen to TranactionServe on "+tport)

	//Declaring and initializing the gRPC server
	srv := grpc.NewServer()
	transactionService.RegisterTransactionServer(srv, &server{})
	reflection.Register(srv)

	err = srv.Serve(transactionSrv)
	ut.CheckError(err, "Could not connect to transaction connect")
}
