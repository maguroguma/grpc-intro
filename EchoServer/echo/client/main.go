package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/my0k/grpc-intro/EchoServer/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	// grpc.ClientConnを生成
	// 本来はデフォルトでTLSが必須なので、insecureオプションを有効にして簡略化している
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn) // clientの生成（クライアントスタブの実体）
	msg := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg}) // リクエストに入力文字列を乗せる
	if err != nil {
		log.Println(err)
	}
	log.Println(r.GetMessage()) // レスポンスの文字列を表示して終了する
}
