/*
 * File: main.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:01:17 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This this Balance Service class.
 * -----
 * Last Modified: Sunday, 17th March 2019 8:26:21 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package main

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	balanceService "github.com/nknab/Moneway/api/services/balance/pb"
	"github.com/nknab/Moneway/database"
	ut "github.com/nknab/Moneway/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type server struct{}

var (
	config = "../../../config/config.env"
	db     *sql.DB
)

/**
 * @brief This gets a user's balance.
 *
 * @param ctx context.Context //Context
 * @param data *balanceService.GetBalanceRequest // Reference to the GetBalanceRequest Object
 *
 * @return *balanceService.GetBalanceReply
 * @return error
 */
func (s *server) GetBalance(ctx context.Context, data *balanceService.GetBalanceRequest) (*balanceService.GetBalanceReply, error) {
	table := "account"

	conditions := []string{"account_id", ut.IntToString(data.AccountID), "balance"}
	balance := database.Select(ctx, table, conditions, db)
	value, _ := strconv.ParseFloat(balance, 32)
	response := &balanceService.GetBalanceReply{
		Success: true,
		Amount:  value,
	}
	return response, nil
}

/**
 * @brief This Update the balance in the database.
 *
 * @param ctx context.Context //Context
 * @param data *balanceService.UpdateBalanceRequest // Reference to the UpdateBalanceRequest Object
 *
 * @return *balanceService.UpdateBalanceReply
 * @return error
 */
func (s *server) UpdateBalance(ctx context.Context, data *balanceService.UpdateBalanceRequest) (*balanceService.UpdateBalanceReply, error) {
	table := "account"
	params := []string{"account_id", ut.IntToString(data.AccountID), "balance", ut.FloatToString(data.Amount)}

	success := database.Update(ctx, table, params, db)
	response := &balanceService.UpdateBalanceReply{
		Success: success,
	}
	return response, nil
}

func main() {
	//Initializing the database
	db = database.Init(config)

	//Reading in the configuration variables
	err := godotenv.Load(config)
	ut.CheckError(err, "Error loading config.env file")
	var port string
	port = ":" + os.Getenv("BALANCE_PORT")

	//Listening to the balance service server
	balanceConnect, err := net.Listen("tcp", port)
	ut.CheckError(err, "Could not connect to "+port)

	//Declaring and initializing the gRPC server
	srv := grpc.NewServer()
	balanceService.RegisterBalanceServer(srv, &server{})
	reflection.Register(srv)

	err = srv.Serve(balanceConnect)
	ut.CheckError(err, "Could not connect to balance connect")
}
