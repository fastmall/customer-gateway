package controller

import (
	"github.com/fastmall/customer-gateway/dubbo"
	"github.com/fastmall/goods/api"
	"github.com/gin-gonic/gin"
	"strconv"
)

var goodsService = dubbo.GoodsService

func GetGoodsDetail(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	detail, err := goodsService.GetItemDetail(c, &api.GetItemDetailRequest{ItemId: id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "%v", detail)
}
