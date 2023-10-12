package svc

import (
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"shorturl/internal/config"
	"shorturl/model"
	"shorturl/sequence"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map
	Sequence      sequence.Sequence
	BlackList     map[string]struct{}
	Filter        *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	List := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for i := 0; i < len(List); i++ {
		List[c.ShortUrlBlackList[i]] = struct{}{}
	}
	// 初始化布隆过滤器
	// 初始化 redisBitSet
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	// 声明一个bitSet, key="test_key"名且bits是1024位
	filter := bloom.New(store, "bloom_filter", 20*(1<<20))
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(sqlx.NewMysql(c.ShortUrlDb.DSN), c.CacheRedis),
		//Sequence:      sequence.NewMysql(c.SequenceDb.DSN),
		Sequence:  sequence.NewRedis(c.CacheRedis[0].Host),
		BlackList: List,
		Filter:    filter,
	}
}
