package api

import (
	"fmt"
	"gin_demo/serializer"
	"gin_demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取商品的id
func GetGoodsWithDetail(c *gin.Context) {
	// 获取id
	id := c.Param("id")
	fmt.Printf("id is: %s\n", id)

	// 访问数据库(假操作)
	// 初始化图片
	photos := service.GetPhotosService()
	// 初始化标题
	title := service.GetTitleService()
	// 初始化价格
	price := service.GetPriceService()

	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "查询成功",
		Data: gin.H{
			"id":     id,
			"photos": photos,
			"title":  title,
			"price":  price,
		},
	})
}
