package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shorturl/internal/svc"
	"shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 在布隆过滤器中查询是否存在，不存在的请求就不要访问redis了
	exists, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil || !exists {
		logx.Errorw("Bloom Filter failed", logx.LogField{Value: err.Error(), Key: "err"})
		return nil, Err404
	}
	fmt.Println("开始数据库查询...")
	// 拿到短链接，返回长链接
	one, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: req.ShortUrl,
		Valid:  true,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, Err404
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.Field("err", err))
		return nil, err
	}
	return &types.ShowResponse{LongUrl: one.Lurl.String}, nil
}
