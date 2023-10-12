package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"shorturl/model"
	"shorturl/pkg/base62"
	"shorturl/pkg/connect"
	"shorturl/pkg/md5"
	"shorturl/pkg/urlTool"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1.校验数据
	// 1.1 数据不为空
	// 1.2 输入的长链接 能够ping 通
	if !connect.Get(req.LongUrl) {
		return &types.ConvertResponse{}, errors.New("无效的链接")
	}
	// 1.3 判断之前是否已经转链过 （数据库中是否已存在该长链）
	md5Value := md5.Sum([]byte(req.LongUrl))
	one, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: md5Value,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转为%s", one.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	// 1.4 该长链是不是一个短连接（避免循环转链）
	basePath, err := urlTool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urlTool.GetBasePath failed", logx.Field("lUrl", req.LongUrl), logx.Field("err", err))
		return nil, errors.New("查询基础地址失败")
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该链接已经是短链了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	var sUrl string
	// 2. 取号
	for {
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			logx.Errorw("Sequence.Next failed", logx.LogField{Key: "err", Value: err.Error()})
			return nil, err
		}
		// 3. 号码转短链
		sUrl = base62.Int2String(int64(seq))
		fmt.Println("-> seq:", seq, "	->sUrl:", sUrl)
		fmt.Println("->:tets：", base62.String2Int(sUrl))
		//todo: 查看转后的短链接是否在 blackList 中
		if _, ok := l.svcCtx.BlackList[sUrl]; ok {
			continue
		}
		break
	}
	// 存入布隆过滤器中
	if err := l.svcCtx.Filter.Add([]byte(sUrl)); err != nil {
		logx.Errorw("Filter.Add() failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 4. 存储长短链接映射关系
	l.svcCtx.ShortUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{req.LongUrl, true},
		Md5:  sql.NullString{md5Value, true},
		Surl: sql.NullString{sUrl, true},
	})
	// 5. 返回响应
	shortUrl := l.svcCtx.Config.ShortDoamin + "/" + sUrl
	return &types.ConvertResponse{ShortUrl: shortUrl}, nil
}
