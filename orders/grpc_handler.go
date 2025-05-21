package main

import (
	"context"
	"log"

	pb "github.com/DavidReque/order-management-system/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received: %v", req)
	o := &pb.Order{
		ID: "42",
	}

	return o, nil

}
