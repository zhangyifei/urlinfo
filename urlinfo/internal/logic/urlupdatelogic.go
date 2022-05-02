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

type UrlUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUrlUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UrlUpdateLogic {
	return &UrlUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UrlUpdateLogic) UrlUpdate(req *types.UpdateRequest) (resp *types.Response, err error) {

	l.Logger.Info(req.Hostnameport + "----" + req.Queryparamter)

	testurl := model.Url{Hostnameport: req.Hostnameport, Queryparamter: req.Queryparamter, UpdateDate: time.Now()}

	info, err := l.svcCtx.Model.Upsert(l.ctx, &testurl)

	l.Logger.Info(info)

	return &types.Response{
		Message: strconv.Itoa(info.Updated),
	}, err
}
