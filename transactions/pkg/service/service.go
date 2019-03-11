/*
 * File: service.go
 * Project: Moneway Intern Assesment
 * File Created: Monday, 11th March 2019 6:36:20 AM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.0
 * Brief:
 * -----
 * Last Modified: Monday, 11th March 2019 10:03:26 AM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package service

import (
	"context"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nknab/MonewayV1.0/balance/pkg/grpc/pb"
	db "github.com/nknab/MonewayV1.0/database"
	"google.golang.org/grpc"
)

type Transaction struct {
	AccountID   string
	Description string
	Amount      float32
	Currency    string
}

// TransactionsService describes the service.
type TransactionsService interface {
	// Add your methods here
	Transct(ctx context.Context, transaction Transaction) error
}

type basicTransactionsService struct {
	balanceServiceClient pb.BalanceClient
}

func (b *basicTransactionsService) Transct(ctx context.Context, transaction Transaction) (e0 error) {
	// TODO implement the business logic of Transct
	//Initializing the Database package
	db.Init()

	//This value will comes from the balance service.
	va := pb.GetBalanceRequest{AccountID: transaction.AccountID}.AccountID

	//Converting the Value to a float32
	value, err := strconv.ParseFloat(va, 32)
	var oldBalance = float32(value)
	var newBalance float32 = 0.0

	table := "transactions"
	columns := []string{"account_id", "description", "amount", "old_balance", "new_balance", "currency"}
	values := []string{transaction.AccountID, transaction.Description, fmt.Sprintf("%f", transaction.Amount), fmt.Sprintf("%f", oldBalance), fmt.Sprintf("%f", newBalance), transaction.Currency}

	// Checking If It is a Debit Or A Credit
	if transaction.TransactionType == 0 {
		newBalance = oldBalance - transaction.Amount

		//Updating the balance via the balance service.
		_ = pb.UpdateBalanceRequest{AccountID: transaction.AccountID, Amount: fmt.Sprintf("%f", newBalance)}
	} else {
		newBalance = oldBalance + transaction.Amount
	}

	// Checking if there is enough Funds in the account before performing transaction
	if newBalance >= 0.0 {
		db.Insert(table, columns, values)
	} else {
		success = false
		fmt.Println("Not Enough Funds in Account")
	}

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return e0
}

// NewBasicTransactionsService returns a naive, stateless implementation of TransactionsService.
func NewBasicTransactionsService() TransactionsService {
	var etcdServer = "http://etcd:2379"

	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		log.Printf("Not able to connect to etcd: %s", err.Error())
		return new(basicTransactionsService)
	}

	entries, err := client.GetEntries("/services/balance/")
	if err != nil || len(entries) == 0 {
		log.Printf("Not able to get prefix entries: %s", err.Error())
		return new(basicTransactionsService)
	}

	conn, err := grpc.Dial(entries[0], grpc.WithInsecure())

	if err != nil {
		log.Printf("Not Able to Connect to Balance: %s", err.Error())
		return new(basicTransactionsService)
	}

	log.Printf("Connected to balance")

	return &basicTransactionsService{
		balanceServiceClient: pb.NewBalanceClient(conn),
	}
}

// New returns a TransactionsService with all of the expected middleware wired in.
func New(middleware []Middleware) TransactionsService {
	var svc TransactionsService = NewBasicTransactionsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
