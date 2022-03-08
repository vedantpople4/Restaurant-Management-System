package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders/", controller.GetOrders())
	incomingRoutes.GET("/orders/: order_id", controller.GetOrders())
	incomingRoutes.POST("/orders", controller.CreateOrder())
	incomingRoutes.POST("/orders/:order_id", controller.UpdateOrder())
}
