package main

import (
	"fmt"
	"net"
	"pusher/auth/auth"
	"pusher/auth/wechat"

	authpb "pusher/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()

	lis, err := net.Listen("tcp", "8081")
	if err != nil {
		fmt.Print("error")
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "",
			AppSecret: "",
		},
		Logger: logger,
	})

	s.Serve(lis)
}
