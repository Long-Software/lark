package api

import (
	"errors"

	"github.com/Long-Software/Bex/cmd/shorten-url/internal/helpers"
	"github.com/Long-Software/Bex/cmd/shorten-url/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) ResolveURL(c *gin.Context) {
	url, _ := c.Params.Get("url")

	value, err := s.store.Read(c, url)
	if err != nil {
		errorResponse(c, errors.New("Not found"))
		return
	}
	successResponse(c, value, 200)
}
func (s *Server) ShortenURL(c *gin.Context) {
	var req models.Request
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		errorResponse(c, errors.New("cannot parse JSON"))
		return
	}

	// implement rate limiting
	// check if the url is valid
	if !govalidator.IsURL(req.URL) {
		errorResponse(c, errors.New("Invalid URL"))
		return
	}

	// check for domain error

	// enforce https, SSL
	req.URL = helpers.EnforceHTTP(req.URL)
}
func errorResponse(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{"status": "error", "error": err.Error()})
}
func successResponse(ctx *gin.Context, data interface{}, code int) {
	ctx.JSON(code, gin.H{"status": "success", "data": data})
}
