package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus/", controller.GetMenus())
	incomingRoutes.GET("/menus/: menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.POST("/menus/:menu_id", controller.UpdateMenu())
}
