package svc

import (
	"urlinfo/urlinfo/internal/config"
	"urlinfo/urlinfo/internal/middleware"
	"urlinfo/urlinfo/internal/model"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Model      model.UrlModel
	Tokenlimit rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Model:      model.NewUrlModel(c.DataSource, c.Collection, c.Cache),
		Tokenlimit: middleware.NewTokenlimitMiddleware(c).Handle,
	}
}
