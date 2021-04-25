package hello

import (
	"context"
	helloworldpb "pusher/auth/api/gen/v1/hello"
)

type Service struct {
	helloworldpb.UnimplementedGreaterServer
}

func (s *Service) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{Message: in.Name + "world"}, nil
}
