package logic

import (
	"context"

	"urlinfo/urlinfo/internal/model"
	"urlinfo/urlinfo/internal/svc"
	"urlinfo/urlinfo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type urlinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewurlinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *urlinfoLogic {
	return &urlinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *urlinfoLogic) UrlLookup(req *types.Request) (resp *types.LookupResponse, err error) {

	queryString := model.DBQueryString{Hostnameport: req.Hostnameport, Queryparamter: req.Queryparamter}

	url, err := l.svcCtx.Model.FindOne(l.ctx, queryString)

	switch err {
	case nil:
		return &types.LookupResponse{
			Message: url.Hostnameport + "--" + url.Queryparamter + "is invalid",
			Allow:   false,
		}, nil
	case model.ErrNotFound:
		return &types.LookupResponse{
			Message: "the url is valid",
			Allow:   true,
		}, nil
	default:
		return nil, err
	}
}
