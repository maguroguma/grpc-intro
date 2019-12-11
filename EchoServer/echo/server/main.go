package main

import (
	"log"
	"net"

	pb "github.com/my0k/grpc-intro/EchoServer/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port) // TCPのリスナーを作成
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	srv := grpc.NewServer()                           // gRPCサーバを作成する
	pb.RegisterEchoServiceServer(srv, &echoService{}) // pbの登録関数でechoServiceをgRPCサーバへ登録する
	log.Printf("start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
