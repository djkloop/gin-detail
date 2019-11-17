package api

import (
	"gin_demo/serializer"
	"gin_demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取商品的id
func GetGoodsWithDetail(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "查询成功",
		Data: service.GetDetail(id),
	})
}
