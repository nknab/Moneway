package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/nknab/Moneway/balance/pkg/service"
)

// GetBalanceRequest collects the request parameters for the GetBalance method.
type GetBalanceRequest struct {
	AccountID string `json:"account_id"`
}

// GetBalanceResponse collects the response parameters for the GetBalance method.
type GetBalanceResponse struct {
	Balance string `json:"balance"`
	E1      error  `json:"e1"`
}

// MakeGetBalanceEndpoint returns an endpoint that invokes GetBalance on the service.
func MakeGetBalanceEndpoint(s service.BalanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBalanceRequest)
		balance, e1 := s.GetBalance(ctx, req.AccountID)
		return GetBalanceResponse{
			Balance: balance,
			E1:      e1,
		}, nil
	}
}

// Failed implements Failer.
func (r GetBalanceResponse) Failed() error {
	return r.E1
}

// UpdateBalanceRequest collects the request parameters for the UpdateBalance method.
type UpdateBalanceRequest struct {
	AccountID string `json:"account_id"`
	Amount    string `json:"amount"`
}

// UpdateBalanceResponse collects the response parameters for the UpdateBalance method.
type UpdateBalanceResponse struct {
	Success string `json:"success"`
	E1      error  `json:"e1"`
}

// MakeUpdateBalanceEndpoint returns an endpoint that invokes UpdateBalance on the service.
func MakeUpdateBalanceEndpoint(s service.BalanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBalanceRequest)
		success, e1 := s.UpdateBalance(ctx, req.AccountID, req.Amount)
		return UpdateBalanceResponse{
			Success: success,
			E1:      e1,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateBalanceResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetBalance implements Service. Primarily useful in a client.
func (e Endpoints) GetBalance(ctx context.Context, accountID string) (string, error) {
	request := GetBalanceRequest{AccountID: accountID}
	response, err := e.GetBalanceEndpoint(ctx, request)
	if err != nil {
		return "", nil
	}
	return response.(GetBalanceResponse).Balance, nil
}

// UpdateBalance implements Service. Primarily useful in a client.
func (e Endpoints) UpdateBalance(ctx context.Context, accountID string, amount string) (string, error) {
	request := UpdateBalanceRequest{
		AccountID: accountID,
		Amount:    amount,
	}
	response, err := e.UpdateBalanceEndpoint(ctx, request)
	if err != nil {
		return "", nil
	}
	return response.(UpdateBalanceResponse).Success, nil
}
