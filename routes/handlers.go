package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nknab/Moneway/api/services/balance/pb"
	"github.com/nknab/Moneway/api/services/transaction/pb"
	ut "github.com/nknab/Moneway/util"
	"net/http"
	"strconv"
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

var(
	transactClient = transactionService.NewTransactionClient(ut.Connect(":8081", "Transcon"))
	balanceClient = balanceService.NewBalanceClient(ut.Connect(":8082", "BalConn"))
)

func transact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var transaction Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	ut.CheckError(err, "Tran")

	request := transactionService.MakeTransaction{
		AccountID:       transaction.AccountID,
		Description:     transaction.Description,
		Amount:          transaction.Amount,
		Currency:        transaction.Currency,
		TransactionType: transaction.TransactionType,
	}

	//fmt.Println(transaction.AccountID)

	//aid, _ := strconv.ParseInt(r.FormValue("AccountID"), 10,64)
	//amt, _ := strconv.ParseFloat(r.FormValue("Amount"), 64)
	//request := transactionService.MakeTransaction{
	//	AccountID:       int32(aid),
	//	Description:     r.FormValue("Description"),
	//	Amount:          amt,
	//	Currency:        r.FormValue("Currency"),
	//	TransactionType: r.FormValue("TransactionType"),
	//}

	//fmt.Println(r.Body)
	//ctx, stop := context.WithTimeout(context.Background(), 30*time.Second)
	//defer stop()

	fmt.Println("Before")
	//var response, _ = transactClient.Transact(ctx, &request)
	fmt.Println("After")
	_ = json.NewEncoder(w).Encode(request)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	accountID, _ := strconv.ParseInt(params["accountID"], 10, 64)

	request := balanceService.GetBalanceRequest{
		AccountID: int32(accountID),
	}
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	fmt.Println("Before")
	response, err := balanceClient.GetBalance(ctx, &request)
	fmt.Println("After")
	ut.CheckError(err, "Resp")
	var balance Balance
	balance.AccountID = int32(accountID)
	balance.Amount = response.Amount
	_ = json.NewEncoder(w).Encode(balance)
}