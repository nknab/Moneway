package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/nknab/Moneway/transactions/pkg/service"
)

// TransctRequest collects the request parameters for the Transct method.
type TransctRequest struct {
	Transaction service.Transaction `json:"transaction"`
}

// TransctResponse collects the response parameters for the Transct method.
type TransctResponse struct {
	E0 error `json:"e0"`
}

// MakeTransctEndpoint returns an endpoint that invokes Transct on the service.
func MakeTransctEndpoint(s service.TransactionsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TransctRequest)
		e0 := s.Transct(ctx, req.Transaction)
		return TransctResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r TransctResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Transct implements Service. Primarily useful in a client.
func (e Endpoints) Transct(ctx context.Context, transaction service.Transaction) (e0 error) {
	request := TransctRequest{Transaction: transaction}
	response, err := e.TransctEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TransctResponse).E0
}
