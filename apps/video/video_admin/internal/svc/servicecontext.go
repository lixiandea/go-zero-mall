package svc

import (
	"github.com/lixiandea/go-zero-mall/apps/video/video_admin/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
