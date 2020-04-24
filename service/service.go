package service

import (
	"gomall-center/pkg/storage"
	"gomall-center/pkg/web"

	"github.com/jinzhu/gorm"
)

// Service 基础结构
type Service struct {
	*web.RequestContext
	*gorm.DB
	//*redis.Client
}

// InitService init base service
func InitService(ctx *web.RequestContext) Service {
	srv := Service{
		ctx,
		storage.GetDB(ctx.Host),
		//storage.GetRedis(ctx.Host),
	}
	return srv
}
