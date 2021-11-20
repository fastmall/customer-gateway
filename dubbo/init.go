package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/fastmall/customer/api"
)

var CustomerService = &api.CustomerServiceClientImpl{}

func StartDubboConsumer() {
	config.SetConsumerService(CustomerService)
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
