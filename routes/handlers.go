package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nknab/Moneway/api/services/balance/pb"
	"github.com/nknab/Moneway/api/services/transaction/pb"
	ut "github.com/nknab/Moneway/util"
	"math"
	"net/http"
	"strconv"
)

//type Transaction struct {
//	AccountID       int32   `json:"accountID"`
//	Description     string  `json:"description"`
//	Amount          float64 `json:"amount"`
//	Currency        string  `json:"currency"`
//	TransactionType string  `json:"transactionType"`
//}

type Balance struct {
	AccountID int32   `json:"accountID"`
	Amount    float64 `json:"amount"`
}

var (
	transactClient transactionService.TransactionClient
	balanceClient  balanceService.BalanceClient
)

func transact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//var transaction Transaction
	//
	//err := json.NewDecoder(r.Body).Decode(&transaction)
	//ut.CheckError(err, "Tran")
	//
	//request := transactionService.MakeTransaction{
	//	AccountID:       transaction.AccountID,
	//	Description:     transaction.Description,
	//	Amount:          transaction.Amount,
	//	Currency:        transaction.Currency,
	//	TransactionType: transaction.TransactionType,
	//}

	aid, _ := strconv.ParseInt(r.FormValue("AccountID"), 10, 64)
	amt, _ := strconv.ParseFloat(r.FormValue("Amount"), 64)
	request := transactionService.MakeTransaction{
		AccountID:       int32(aid),
		Description:     r.FormValue("Description"),
		Amount:          amt,
		Currency:        r.FormValue("Currency"),
		TransactionType: r.FormValue("TransactionType"),
	}
	var response, _ = transactClient.Transact(context.Background(), &request)
	_ = json.NewEncoder(w).Encode(response)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	accountID, _ := strconv.ParseInt(params["accountID"], 10, 64)
	fmt.Println(accountID)
	request := balanceService.GetBalanceRequest{
		AccountID: int32(accountID),
	}
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	response, err := balanceClient.GetBalance(ctx, &request)
	ut.CheckError(err, "Resp")
	var balance Balance
	balance.AccountID = int32(accountID)
	balance.Amount = math.Round(response.Amount * 100) / 100
	_ = json.NewEncoder(w).Encode(balance)
}

func Run() {
	transactClient = transactionService.NewTransactionClient(ut.Connect(":8081", "Transcon"))
	balanceClient = balanceService.NewBalanceClient(ut.Connect(":8082", "BalConn"))
}
