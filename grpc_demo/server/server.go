package main

import (
	"context"
	"flag"
	"net"

	"google.golang.org/grpc"
	// 设置引用别名
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "启动端口号")
	flag.Parse()
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello, world"}, nil
}

func main() {
	//创建 gRPC Server 对象
	server := grpc.NewServer()
	//将 GreeterServer（其包含需要被调用的服务端接口）注册到 gRPC Server的内部注册中心
	pb.RegisterGreeterServer(server, &GreeterServer{})
	//创建 Listen，监听 TCP 端口
	lis, _ := net.Listen("tcp", ":"+port)
	//gRPC Server 开始 lis.Accept，直到 Stop 或 GracefulStop
	server.Serve(lis)
}
