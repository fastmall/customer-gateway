package controller

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/customer-gateway/dubbo"
	"github.com/fastmall/customer/api"
	"github.com/gin-gonic/gin"
)

var customerService = dubbo.CustomerService

func Register(c *gin.Context) {

	createCustomerResponse, err := customerService.CreateCustomer(context.Background(), &api.CreateCustomerRequest{
		Username: "zhurungen",
		Password: "test01",
	})

	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(createCustomerResponse.UserId)
}
