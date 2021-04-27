package main

import (
	"fmt"
	"log"
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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// authpb.RegisterAuthServiceServer(s, &auth.Service{
	// 	OpenIDResolver: &wechat.Service{
	// 		AppID:     "",
	// 		AppSecret: "",
	// 	},
	// 	Logger: logger,
	// })

	// s.Serve(lis)

	// conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithInsecure())

	// gwmux := runtime.NewServeMux()

	// err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	// gwServer := &http.Server{
	// 	Addr:    ":8090",
	// 	Handler: gwmux,
	// }

	// gwServer.ListenAndServe()
}
