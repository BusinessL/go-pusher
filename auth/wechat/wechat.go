package wechat

import (
	weapp "github.com/medivhzhan/weapp/v2"
)

type Service struct {
	AppID     string
	AppSecret string
}

func (s *Service) Resolve(code string) (string, error) {
	resp, err := weapp.Login(s.AppID, s.AppSecret, code)

	if err != nil {
		return "错误", err
	}

	return resp.OpenID, nil
}
