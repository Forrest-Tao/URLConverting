package sequence

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client redis.Client
}

func NewRedis(DSN string) Sequence {
	rdb := redis.NewClient(&redis.Options{
		Addr: DSN,
	})
	return &Redis{
		Client: *rdb,
	}
}

func (r *Redis) Next() (seq uint64, err error) {
	cmd := r.Client.Incr(context.Background(), "sequence_num")
	if cmd != nil {
		return uint64(cmd.Val()), nil
	}
	return 0, err
}
