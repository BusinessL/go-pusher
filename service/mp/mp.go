package mp

import (
	"context"
	"fmt"
	mppb "pusher/proto/mp"
	"pusher/tool"
)

type Server struct {
	mppb.UnimplementedMpServiceServer
}

func (s *Server) Push(ctx context.Context, in *mppb.Request) (*mppb.Response, error) {
	te := tool.GetMpAccessToken()

	fmt.Printf("%v", te.ErrCode)
	res := tool.SendTemplate(te.AccessToken)

	return &mppb.Response{Code: res.ErrCode}, nil
}
