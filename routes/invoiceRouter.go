package routes

import (
	controller "restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices/", controller.GetInvoices())
	incomingRoutes.GET("/invoices/: invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.POST("/invoices/:invoices_id", controller.UpdateInvoice())
}
