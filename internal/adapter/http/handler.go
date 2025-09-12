package http

import (
	"app/internal/adapter/parse"
	"app/internal/adapter/presenter"
	"app/internal/adapter/validate"
	"app/internal/usecase"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.DecodeInputPort
}

func NewHandler(uc usecase.DecodeInputPort) *Handler {
	return &Handler{uc: uc}
}

// POST /v1/decode
// - text/plain or application/json {"input":"..."} をサポート
type jsonReq struct {
	Input string `json:"input"`
}

func (h *Handler) Decode(c *gin.Context) {
	var raw string

	ct := c.ContentType()
	switch ct {
	case "application/json":
		var jr jsonReq
		if err := c.ShouldBindJSON(&jr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "BadRequest", "message": "invalid json"}})
			return
		}
		raw = jr.Input
	default:
		// text/plain などは生ボディで
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "BadRequest", "message": "cannot read body"}})
			return
		}
		raw = string(body)
	}

	out, err := h.uc.Decode(usecase.DecodeInput{Raw: raw})
	if err != nil {
		code, msg := mapError(err)
		c.JSON(code, gin.H{"error": gin.H{"code": httpStatusCodeName(code), "message": msg}})
		return
	}

	c.JSON(http.StatusOK, presenter.FromUsecase(out))
}

func mapError(err error) (int, string) {
	switch err {
	case parse.ErrParse:
		return http.StatusBadRequest, err.Error()
	case validate.ErrEmptyPayload, validate.ErrNonASCII, validate.ErrBadNumberToken:
		return http.StatusUnprocessableEntity, err.Error()
	default:
		// domain/decoder のエラー等も一律422に
		return http.StatusUnprocessableEntity, err.Error()
	}
}

func httpStatusCodeName(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "BadRequest"
	case http.StatusUnprocessableEntity:
		return "UnprocessableEntity"
	default:
		return "InternalServerError"
	}
}
