package main

import (
	"context"

	pb "github.com/my0k/grpc-intro/EchoServer/echo/proto"
)

type echoService struct{}

// Echoメソッドを実装
// EchoRequestからGetMessageメソッドでクライアントが送信した文言を受け取りResponseに含める
func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}
