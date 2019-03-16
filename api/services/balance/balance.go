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

	bpb "github.com/nknab/Moneway/api/services/balance/protobuf"
	db "github.com/nknab/Moneway/database"
	ut "github.com/nknab/Moneway/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// "context"
	"net"
)

type balanceServer struct{}

func (balanceServer) GetBalance(data *bpb.GetBalanceRequest) (*bpb.GetBalanceReply, error) {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	//Initializing the Database package
	db.Init()
	table := "account"
	conditions := []string{"account_id", ut.IntToString(data.AccountID), "balance"}
	balance := db.Select(ctx, table, conditions)
	bal, _ := strconv.ParseFloat(balance, 32)
	response := &bpb.GetBalanceReply{
		Success: true,
		Amount:  bal,
	}

	return response, nil

}

func (balanceServer) UpdateBalance(ctx context.Context, data *bpb.UpdateBalanceRequest) (*bpb.UpdateBalanceReply, error) {
	//ctx, stop := context.WithCancel(context.Background())
	//defer stop()

	//Initializing the Database package
	db.Init()
	table := "account"
	params := []string{"account_id", ut.IntToString(data.AccountID), "balance", ut.FloatToString(data.Amount)}

	// Updating the Balance
	success := db.Update(ctx, table, params)

	response := &bpb.UpdateBalanceReply{
		Success: success,
	}

	return response, nil
}

//func (b *balanceServer) GetBalance(ctx context.Context, accountID int32) (*bpb.GetBalanceReply, error) {
//
//	//Initializing the Database package
//	db.Init()
//	table := "account"
//	conditions := []string{"account_id", ut.IntToString(accountID), "balance"}
//	balance := db.Select(ctx, table, conditions)
//	bal, _ := strconv.ParseFloat(balance, 32)
//	response := &bpb.GetBalanceReply{
//		Amount: bal,
//	}
//
//	return response, nil
//}
//
//func (b *balanceServer) UpdateBalance(ctx context.Context, accountID int32, amount float64) (*bpb.UpdateBalanceReply, error) {
//
//	//Initializing the Database package
//	db.Init()
//	table := "account"
//	params := []string{"account_id", ut.IntToString(accountID), "balance", ut.FloatToString(amount)}
//
//	// Updating the Balance
//	success := db.Update(ctx, table, params)
//
//	response := &bpb.UpdateBalanceReply{
//		Success: success,
//	}
//
//	return response, nil
//}

func main() {
	srv := grpc.NewServer()
	bpb.RegisterBalanceServer(srv, &balanceServer{})
	reflection.Register(srv)

	balanceConnect, err := net.Listen("tcp", ":8889")
	ut.CheckError(err, "Could not connect to :8889")

	if err = srv.Serve(balanceConnect); err != nil {
		log.Fatalf("Could not connect to: %v", err)

	}
}
