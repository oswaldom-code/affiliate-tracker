package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/http/rest/dto"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{Status: true, Message: "pong"})
}
