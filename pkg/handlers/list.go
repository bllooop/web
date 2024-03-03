package handler

import (
	"net/http"
	"strconv"
	"web"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createProductList(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	var input web.ProductList
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	//calling service method
	id, err := h.services.ProductList.Create(shopid, input)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllProductsListsResponse struct {
	Data []web.ProductList `json:"data"`
}

func (h *Handler) getAllProductsLists(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	lists, err := h.services.ProductList.GetAllLists(shopid)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllProductsListsResponse{
		Data: lists,
	})
}
func (h *Handler) getProductListById(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	list, err := h.services.ProductList.GetById(shopid, id)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateProductList(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input web.UpdateProductListInput
	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ProductList.Update(shopid, id, input); err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusReponse{
		Status: "ok",
	})
}
func (h *Handler) deleteProductList(c *gin.Context) {
	shopid, err := getShopId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newError(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.services.ProductList.Delete(shopid, id)
	if err != nil {
		newError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusReponse{
		Status: "ok",
	})
}
