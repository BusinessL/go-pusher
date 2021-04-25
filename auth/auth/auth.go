package auth

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authpb "pusher/auth/api/gen/v1"
)

type Service struct {
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
	authpb.UnimplementedAuthServiceServer
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	//openID, err:= s.OpenIDResolver.Resolve()

	s.Logger.Info("received", zap.String("code", req.Code))

	openID, err := s.OpenIDResolver.Resolve(req.Code)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot", err)
	}

	return &authpb.LoginResponse{
		AccessToken: "token for" + openID,
		ExpiresIn:   7200,
	}, nil
}
