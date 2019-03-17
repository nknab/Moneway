/*
 * File: balance.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:01:17 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief:
 * -----
 * Last Modified: Friday, 15th March 2019 7:02:10 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	"log"
	"strconv"

	"github.com/nknab/Moneway/api/services/balance/pb"
	db "github.com/nknab/Moneway/database"
	ut "github.com/nknab/Moneway/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type server struct{}

func (s *server) GetBalance(ctx context.Context, data *balanceService.GetBalanceRequest) (*balanceService.GetBalanceReply, error) {

	db.Init("../../../config/config.toml")
	table := "account"

	conditions := []string{"account_id", ut.IntToString(data.AccountID), "balance"}
	balance := db.Select(ctx, table, conditions)
	value, _ := strconv.ParseFloat(balance, 32)
	response := &balanceService.GetBalanceReply{
		Success: true,
		Amount:  value,
	}
	return response, nil
}

func (s *server) UpdateBalance(ctx context.Context, data *balanceService.UpdateBalanceRequest) (*balanceService.UpdateBalanceReply, error) {

	db.Init("../../../config/config.toml")
	table := "account"
	params := []string{"account_id", ut.IntToString(data.AccountID), "balance", ut.FloatToString(data.Amount)}

	success := db.Update(ctx, table, params)
	response := &balanceService.UpdateBalanceReply{
		Success: success,
	}
	return response, nil
}

func main() {
	balanceConnect, err := net.Listen("tcp", ":8082")
	ut.CheckError(err, "Could not connect to :8082")

	srv := grpc.NewServer()
	balanceService.RegisterBalanceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(balanceConnect); err != nil {
		log.Fatalf("Could not connect to: %v", err)
	}
}
