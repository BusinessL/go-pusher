package auth

import (
	"context"
	"fmt"
	authpb "pusher/proto/auth"

	wxapp "github.com/medivhzhan/weapp/v2"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	res, err := wxapp.Login("appid", "app", in.Code)
	if err != nil {
		fmt.Errorf("wxapp: %v", err)
	}

	return &authpb.LoginResponse{AccessToken: res.OpenID}, nil
}
