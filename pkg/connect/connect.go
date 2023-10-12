package connect

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var Client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: time.Second * 2,
}

func Get(url string) bool {
	resp, err := Client.Get(url)
	if err != nil {
		logx.Errorw("connect Client.Get failed", logx.Field("err", err.Error()))
		return false
	}
	defer resp.Body.Close()
	return http.StatusOK == resp.StatusCode
}
