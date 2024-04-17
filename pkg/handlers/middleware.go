package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// methods for getting shop id and checking for headers of auth and parsing token
const (
	authorizationHeader = "Authorization"
	shopCtx             = "shopId"
)

func (h *Handler) shopIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newError(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 {
		newError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	shopId, err := h.services.Authorization.ParseToken(headerSplit[1])
	if err != nil {
		newError(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(shopCtx, shopId)
}

func getShopId(c *gin.Context) (int, error) {
	id, ok := c.Get(shopCtx)
	if !ok {
		newError(c, http.StatusInternalServerError, "shop id not found")
		return 0, errors.New("shop id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newError(c, http.StatusInternalServerError, "shop id is of invalid type")
		return 0, errors.New("shop id is of invalid type")
	}
	return idInt, nil
}
