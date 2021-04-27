package auth

import (
	"context"
	authpb "pusher/proto/auth"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{AccessToken: in.Code + " world"}, nil
}
