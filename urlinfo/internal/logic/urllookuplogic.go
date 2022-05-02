package logic

import (
	"context"
	"fmt"

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

func (l *urlinfoLogic) UrlLookup(req *types.Request) (resp *types.Response, err error) {

	queryString := model.DBQueryString{Hostnameport: req.Hostnameport, Queryparamter: req.Queryparamter}

	url, err := l.svcCtx.Model.FindOne(l.ctx, queryString)

	if err != nil {
		return nil, err
	}

	fmt.Println(url.Queryparamter)

	return &types.Response{
		Message: url.Hostnameport + "--" + url.UpdateDate.GoString(),
	}, nil
}
