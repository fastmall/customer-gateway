package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	customerApi "github.com/fastmall/customer/api"
	goodsApi "github.com/fastmall/goods/api"
)

var CustomerService = &customerApi.CustomerServiceClientImpl{}

var GoodsService = &goodsApi.GoodsServiceClientImpl{}

func StartDubboConsumer() {
	config.SetConsumerService(CustomerService)
	config.SetConsumerService(GoodsService)
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
