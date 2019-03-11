/*
 * File: service.go
 * Project: Moneway Intern Assesment
 * File Created: Monday, 11th March 2019 5:53:24 AM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.0
 * Brief:
 * -----
 * Last Modified: Monday, 11th March 2019 10:03:50 AM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package service

import (
	"context"

	db "github.com/nknab/MonewayV1.0/database/database"
)

// BalanceService describes the service.
type BalanceService interface {
	// Add your methods here
	GetBalance(ctx context.Context, accountID string) (string, error)
	UpdateBalance(ctx context.Context, accountID string, amount string) (string, error)
}
type basicBalanceService struct{}

func (b *basicBalanceService) GetBalance(ctx context.Context, accountID string) (string, error) {
	// TODO implement the business logic of GetBalance

	//Initializing the Database package
	db.Init()
	table := "account"
	conditions := []string{"account_id", accountID}
	balance := db.Select(table, conditions)

	return balance, nil
}
func (b *basicBalanceService) UpdateBalance(ctx context.Context, accountID string, amount string) (string, error) {
	// TODO implement the business logic of UpdateBalance

	//Initializing the Database package
	db.Init()
	table := "account"
	param := []string{"account_id", accountID, "balance", amount}

	// Updating the Balance
	db.Update(table, param)

	return "", nil
}

// NewBasicBalanceService returns a naive, stateless implementation of BalanceService.
func NewBasicBalanceService() BalanceService {
	return &basicBalanceService{}
}

// New returns a BalanceService with all of the expected middleware wired in.
func New(middleware []Middleware) BalanceService {
	var svc BalanceService = NewBasicBalanceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
