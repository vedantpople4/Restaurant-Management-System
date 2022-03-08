package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems/", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/: orderItem_id", controller.GetorderItem())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.POST("/orderItems/: orderItem_id", controller.UpdateTables())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
}
