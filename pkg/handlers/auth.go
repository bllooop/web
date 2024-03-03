package handler

import (
	"net/http"
	"web"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input web.Shop

	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
	}
	id, err := h.services.Authorization.CreateShop(input)
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	ShopName string `json:"shopname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newError(c, http.StatusBadRequest, err.Error())
	}
	token, err := h.services.Authorization.GenerateToken(input.ShopName, input.Password)
	if err != nil {
		newError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
