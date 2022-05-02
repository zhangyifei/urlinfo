package logic

import (
	"context"
	"strconv"
	"time"

	"urlinfo/urlinfo/internal/model"
	"urlinfo/urlinfo/internal/svc"
	"urlinfo/urlinfo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UrlBatchUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUrlBatchUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UrlBatchUpdateLogic {
	return &UrlBatchUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UrlBatchUpdateLogic) UrlBatchUpdate(req *types.BatchUpdateRequest) (resp *types.Response, err error) {

	urls := make([]*model.Url, 0)

	for _, req := range req.Requests {
		testurl := model.Url{Hostnameport: req.Hostnameport, Queryparamter: req.Queryparamter, UpdateDate: time.Now()}
		urls = append(urls, &testurl)
	}

	err = l.svcCtx.Model.BatchInsert(l.ctx, urls)

	if err != nil {
		return nil, err
	}

	return &types.Response{
		Message: strconv.Itoa(len(req.Requests)) + "  inserted",
	}, nil
}
