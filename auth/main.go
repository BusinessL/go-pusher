package main

import (
	"fmt"
	"net"
	"pusher/auth/hello"

	"google.golang.org/grpc"

	helloworldpb "pusher/auth/api/gen/v1/hello"
)

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		fmt.Print("error")
	}

	s := grpc.NewServer()

	helloworldpb.RegisterGreaterServer(s, &hello.Service{})

	s.Serve(lis)
	// authpb.RegisterAuthServiceServer(s, &auth.Service{
	// 	OpenIDResolver: &wechat.Service{
	// 		AppID:     "",
	// 		AppSecret: "",
	// 	},
	// 	Logger: logger,
	// })

	// s.Serve(lis)
}
