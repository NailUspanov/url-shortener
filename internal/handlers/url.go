package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"regexp"
	"url-shortener/internal/helpers"
)

func (h *Handler) create(c *gin.Context) {

	type CreateRequest struct {
		LongUrl string `json:"long_url"`
	}

	var input CreateRequest
	if err := c.BindJSON(&input); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// check if url is valid
	_, err := url.ParseRequestURI(input.LongUrl)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid URI for request")
		return
	}

	generatedUrl, err := h.services.Create(input.LongUrl)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"short_url": generatedUrl,
	})
}

func (h *Handler) find(c *gin.Context) {

	type FindRequest struct {
		ShortUrl string `uri:"short_url"`
	}

	var input FindRequest
	if err := c.BindUri(&input); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !regexp.MustCompile(`^\w{10}$`).MatchString(input.ShortUrl) {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input type")
		return
	}

	longUrl, err := h.services.Find(input.ShortUrl)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"long_url": longUrl,
	})
}
