package email

import (
	"context"
	"net/smtp"
	emailpb "pusher/proto/email"

	emailTool "github.com/jordan-wright/email"
)

type Server struct {
	emailpb.UnimplementedEmailServiceServer
}

func (s *Server) Send(ctx context.Context, req *emailpb.EmailRequest) (*emailpb.EmailResponse, error) {
	e := emailTool.NewEmail()
	e.From = "Jordan "
	e.To = []string{"max.li@madv360.com"}
	e.Text = []byte("Text body is")
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))

	return &emailpb.EmailResponse{Code: "success"}, nil
}
