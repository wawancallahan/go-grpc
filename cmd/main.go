package main

import (
	"log"
	"net"

	"github.com/wawancallahan/go-grpc/pb"
	"github.com/wawancallahan/go-grpc/protoserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	serveGrpcServer()
}

func serveGrpcServer() {
	grpcUserServer := protoserver.NewUserServer()
	grpcCategoryServer := protoserver.NewCategoryServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcUserServer)
	pb.RegisterCategoryServiceServer(grpcServer, grpcCategoryServer)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":3300")

	if err != nil {
		log.Fatal("Cannot Run Grpc Server")
	}

	log.Printf("start Grpc server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot Start Grpc Server")
	}
}
