package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDb struct {
		DSN string
	}

	SequenceDb struct {
		DSN string
	}

	CacheRedis        cache.CacheConf
	ShortDoamin       string
	ShortUrlBlackList []string
	BaseString        string // bas62指定基础字符串
}
