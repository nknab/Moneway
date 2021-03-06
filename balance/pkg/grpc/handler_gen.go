// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/nknab/Moneway/balance/pkg/endpoint"
	pb "github.com/nknab/Moneway/balance/pkg/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	getBalance    grpc.Handler
	updateBalance grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.BalanceServer {
	return &grpcServer{
		getBalance:    makeGetBalanceHandler(endpoints, options["GetBalance"]),
		updateBalance: makeUpdateBalanceHandler(endpoints, options["UpdateBalance"]),
	}
}
