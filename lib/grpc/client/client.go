package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/pb"
	"log"
	"time"
)

const (
	defaultName = "goku"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	date = flag.String("date", time.Now().Format("2006-01-02"), "date to get awesome show")
)

func main() {
	flag.Parse()
	// Set up a Connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloGRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调SayHi接口
	r, err := c.SayHi(ctx, &pb.Req{Message: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// 调GetAwesomeShow接口
	r2, err := c.GetAwesomeShow(ctx, &pb.GetShowReq{Date: *date})
	if err != nil {
		log.Fatalf("could not get awesome show:%v\n", err)
	}
	log.Printf("there is an awesome show:%v\n", r2.GetShow())
}
