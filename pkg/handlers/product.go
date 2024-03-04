package handler

import (
	"net/http"
	"strconv"
	"web"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProduct(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input web.ProductItem
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.ProductItem.Create(shopid, listId, input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getAllProducts(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input web.ProductItem
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.ProductItem.Create(shopid, listId, input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getProductById(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	product, err := h.services.ProductItem.GetById(shopid, itemId)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input web.UpdateProductItemInput
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ProductItem.Update(shopid, id, input); err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
func (h *Handler) deleteProduct(c *gin.Context) {
	userId, err := getShopId(c)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.ProductItem.Delete(userId, itemId)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
