package http

import "github.com/gin-gonic/gin"

func NewRouter(h *Handler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/decode", h.Decode)
	return r
}
