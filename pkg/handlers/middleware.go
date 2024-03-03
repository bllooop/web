package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
