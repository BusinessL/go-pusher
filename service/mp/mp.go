package mp

import (
	"context"
	mppb "pusher/proto/mp"
)

type Server struct {
	mppb.UnimplementedMpServiceServer
}

func (s *Server) Push(ctx context.Context, in *mppb.Request) (*mppb.Response, error) {
	return &mppb.Response{Message: in.Content}, nil
}
