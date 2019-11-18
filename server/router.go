package server

import (
	"gin_demo/api"
	"gin_demo/middleware"

	"github.com/gin-gonic/gin"
)

// ServerRouter 创建路由
func ServerRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		// 测试后台服务是否畅通接口
		v1.POST("ping", api.Ping)

		// 业务接口
		v1.GET("detail/:id", api.GetGoodsWithDetail)

	}
	return r
}
