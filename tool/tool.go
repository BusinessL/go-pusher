package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	redis "github.com/go-redis/redis/v8"
)

type ResAccessToken struct {
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func GetMpAccessToken(ctx context.Context) (result ResAccessToken) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	//判断缓存是否有值
	// val, err := rdb.Get(ctx, "access_token").Result()
	// if err == nil && val != "" {
	// 	result.AccessToken = val
	// 	return
	// }

	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=fjaljasf&secret=fsdlafs")

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	// rdb.Set(ctx, "access_token", result.AccessToken, 3600*time.Second)

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
