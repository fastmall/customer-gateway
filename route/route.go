package route

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/fastmall/customer-gateway/controller"
	"github.com/gin-gonic/gin"
)

func SetCustomerRoute(r *gin.Engine) {
	g := r.Group("/customer")
	g.POST("/register", controller.Register)
	g.GET("/register", controller.Register)
}
