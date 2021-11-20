package main

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/customer-gateway/route"

	"github.com/fastmall/customer-gateway/dubbo"
	"github.com/gin-gonic/gin"
)

func main() {
	dubbo.StartDubboConsumer()
	r := gin.Default()
	route.SetCustomerRoute(r)
	err := r.Run(":80")
	if err != nil {
		logger.Fatal(err)
	}
}
