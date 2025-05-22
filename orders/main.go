package main

import (
	"context"
	"log"
	"net"

	"github.com/DavidReque/order-management-system/common"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	grpcpAddr string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	grpcpAddr = common.EnvString("GRPC_ADDR", "localhost:2000")

}

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)

	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	log.Println("Starting grpc server on ", grpcpAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
