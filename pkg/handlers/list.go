package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProductList(c *gin.Context) {
	id, _ := c.Get(shopCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllProductsLists(c *gin.Context) {

}
func (h *Handler) getProductListById(c *gin.Context) {

}

func (h *Handler) updateProductList(c *gin.Context) {

}
func (h *Handler) deleteProductList(c *gin.Context) {

}
