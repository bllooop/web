package handler

import (
	"web/pkg/service"

	"github.com/gin-gonic/gin"
)

// handlers
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

/*
var dvs []Device

	func init() {
		dvs = []Device{
			{1, "5F-33-CC-1F-43-82", "2.1.6"},
			{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
		}
	}
*/
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//reg := prometheus.NewRegistry()
	//m := h.newMetrics(reg)
	//m.devices.Set(float64(len(dvs)))
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/api", h.shopIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createProductList)
			lists.GET("/", h.getAllProductsLists)
			lists.GET("/:id", h.getProductListById)
			lists.PUT("/:id", h.updateProductList)
			lists.DELETE("/:id", h.deleteProductList)
			products := lists.Group(":id/product")
			{
				products.POST("/", h.createProduct)
				products.GET("/", h.getAllProducts)
			}
		}
		products := api.Group("products")
		{
			products.GET("/:id", h.getProductById)
			products.PUT("/:id", h.updateProduct)
			products.DELETE("/:id", h.deleteProduct)
		}
		metrics := api.Group("metrics")
		{
			metrics.GET("/", prometheusHandler())
		}
	}
	return router
}
