// 创建带有拦截器的grpc-server
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/pb"
	"log"
	"net"
	"time"
)

/*
拦截器可以用来执行各种其他任务，例如：

1.验证请求身份验证
2.给请求授权
3.记录请求和响应
4.添加延迟或抖动以模拟网络条件
*/

// connectInterceptor
// @Description: 创建拦截器，记录处理的时间
// @param ctx 请求上下文。 该上下文包含有关请求的信息，例如客户端 IP 地址和请求超时
// @param req 请求消息。 使用 Protobuf 序列化和反序列化
// @param info  请求元数据。 该元数据包含有关请求的信息，例如方法名称和服务名称
// @param handler 要调用的下一个处理程序。 这是处理请求的实际代码
func connectInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	// defer会在handler结束之前触发
	defer func() {
		endTime := time.Now()
		log.Printf("Request %s from %s took %dms", info.FullMethod, ctx.Value("remote_ip"), endTime.Sub(startTime).Milliseconds())
	}()

	return handler(ctx, req)
}

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	fmt.Printf("received:%v\n", in.GetMessage())
	// 客户端默认等待1s超时，这里不要sleep太久喔
	// 当然也可以查阅文档，看看怎么延长客户端等待超时的时间
	time.Sleep(time.Millisecond * 100)
	return &pb.Res{Message: "hi " + in.GetMessage()}, nil
}

func (s *server) GetAwesomeShow(ctx context.Context, in *pb.GetShowReq) (*pb.ShowRes, error) {
	fmt.Printf("received:%v\n", in.GetDate())
	time.Sleep(time.Millisecond * 200)
	return &pb.ShowRes{Show: "<friends>"}, nil
}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen:%v\n", err)
	}

	// 创建带有拦截器的 rpc 服务器
	/*
		拦截器可以执行以下操作之一：

		调用 handler(ctx, req, info) 来处理请求。
		返回错误以阻止处理请求。
		修改 ctx、req 或 info 并调用 handler(ctx, req, info)。
	*/
	s := grpc.NewServer(grpc.UnaryInterceptor(connectInterceptor))
	pb.RegisterHelloGRPCServer(s, &server{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	if err = s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
	/*
		整个数据流如下

		客户端启动并连接到 gRPC 服务器；
		客户端调用 Hello() 方法；
		请求到达 gRPC 服务器；
		服务器上的 connect_interceptor 拦截器在处理请求之前执行。 拦截器记录请求的开始时间；
		服务器调用 Greeter 服务的 Hello() 方法实现；
		服务器上的 connect_interceptor 拦截器在处理请求之后执行。 拦截器记录请求的完成时间并计算处理时间；
		服务器向客户端发送响应；
		客户端接收响应；
		客户端关闭连接。
	*/
}
