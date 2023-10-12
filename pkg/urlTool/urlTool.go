package urlTool

import (
	"errors"
	"net/url"
	"path"
)

// GetBasePath 获取URL路径的最后一节
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("路径无域名")
	}
	return path.Base(myUrl.Path), nil
}
