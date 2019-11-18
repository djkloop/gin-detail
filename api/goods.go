package api

import (
	"gin_demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGoodsWithDetail 获取商品的id
func GetGoodsWithDetail(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	res := service.GetDetail(id)
	c.JSON(http.StatusOK, res)
}
