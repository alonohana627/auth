package main

import (
	"auth/database"
	_ "auth/database"
	pb "auth/pb"
	"auth/routes"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	database.InitDB()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	authRoutes := &routes.AuthRoutes{}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthRoutesServer(grpcServer, authRoutes)

	log.Printf("start gRPC server at %s\n", lis.Addr().String())
	log.Fatal(grpcServer.Serve(lis))
}
