package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ResAccessToken struct {
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func GetMpAccessToken() (result ResAccessToken) {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=fjaljasf&secret=fsdlafs")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	return
}

func SendTemplate(accessToken string) (result ResAccessToken) {
	resp, err := http.PostForm("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+accessToken, url.Values{"touser": {"OPENID"}})

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	return
}
