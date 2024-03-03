package handler

import (
	"web/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
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
				products.GET("/:product_id", h.getProductById)
				products.PUT("/:product_id", h.updateProduct)
				products.DELETE("/:product_id", h.deleteProduct)
			}
		}
	}
	return router
}
