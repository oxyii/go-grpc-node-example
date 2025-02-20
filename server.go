package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/oxyii/go-grpc-node-example/proto"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTestServer(grpcServer, &Svc{})
	if grpcServer.Serve(lis) != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Svc struct {
	proto.UnimplementedTestServer
}

func (s *Svc) Unary(ctx context.Context, in *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	return &wrapperspb.StringValue{Value: "Hello, " + in.Value + "!"}, nil
}

func (s *Svc) ClientStream(stream proto.Test_ClientStreamServer) error {
	var resp []string
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: fmt.Sprintf("Hello, %v!", resp)})
		}
		if err != nil {
			return err
		}
		resp = append(resp, in.Value)
	}
}

func (s *Svc) ServerStream(in *wrapperspb.StringValue, stream proto.Test_ServerStreamServer) error {
	for i := 0; i < 2; i++ {
		if err := stream.Send(&wrapperspb.StringValue{Value: fmt.Sprintf("Hello, %v! %d", in.Value, i)}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Svc) BidiStream(stream proto.Test_BidiStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&wrapperspb.StringValue{Value: "Hello, " + in.Value + "!"}); err != nil {
			return err
		}
	}
}
