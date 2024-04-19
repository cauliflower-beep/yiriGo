// 本篇实现参考grpc官方example
// https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/pb"
	"net"
)

/*
grpc 服务端的4个步骤：
1.取出server；
2.挂载方法；
3.注册服务；
4.创建监听
*/

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	fmt.Printf("received:%v\n", in.GetMessage())
	return &pb.Res{Message: "hi " + in.GetMessage()}, nil
}

func (s *server) GetAwesomeShow(ctx context.Context, in *pb.GetShowReq) (*pb.ShowRes, error) {
	fmt.Printf("received:%v\n", in.GetDate())
	return &pb.ShowRes{Show: "<friends>"}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen:%v\n", err)
	}

	// 创建 rpc 服务器
	s := grpc.NewServer()
	pb.RegisterHelloGRPCServer(s, &server{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err = s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
