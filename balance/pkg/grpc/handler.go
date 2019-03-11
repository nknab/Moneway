package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/nknab/Moneway/balance/pkg/endpoint"
	pb "github.com/nknab/Moneway/balance/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeGetBalanceHandler creates the handler logic
func makeGetBalanceHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetBalanceEndpoint, decodeGetBalanceRequest, encodeGetBalanceResponse, options...)
}

// decodeGetBalanceResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeGetBalanceRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetBalanceRequest)
	return endpoint.GetBalanceRequest{AccountID: req.AccountID}, nil
}

// encodeGetBalanceResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetBalanceResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.GetBalanceResponse)
	return &pb.GetBalanceReply{Amount: reply.Balance}, nil
}
func (g *grpcServer) GetBalance(ctx context1.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceReply, error) {
	_, rep, err := g.getBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetBalanceReply), nil
}

// makeUpdateBalanceHandler creates the handler logic
func makeUpdateBalanceHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateBalanceEndpoint, decodeUpdateBalanceRequest, encodeUpdateBalanceResponse, options...)
}

// decodeUpdateBalanceResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeUpdateBalanceRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UpdateBalanceRequest)
	return endpoint.UpdateBalanceRequest{AccountID: req.AccountID, Amount: req.Amount}, nil
}

// encodeUpdateBalanceResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateBalanceResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.UpdateBalanceResponse)
	return &pb.UpdateBalanceReply{Success: reply.Success}, nil
}
func (g *grpcServer) UpdateBalance(ctx context1.Context, req *pb.UpdateBalanceRequest) (*pb.UpdateBalanceReply, error) {
	_, rep, err := g.updateBalance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateBalanceReply), nil
}
