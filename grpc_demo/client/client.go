package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "启动端口号")
	flag.Parse()
}

func main() {
	//创建与给定目标（服务端）的连接句柄。
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	//创建 Greeter 的客户端对象。
	client := pb.NewGreeterClient(conn)
	// 发送 RPC 请求，等待同步响应，得到回调后返回响应结果。
	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
