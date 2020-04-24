package routers

import (
	"gomall-center/api"
	"gomall-center/pkg/web"

	"github.com/gin-gonic/gin"
)

// Register 配置路由
func Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/company", web.Handler(api.CompanyList))
}
