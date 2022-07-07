package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	type CreateRequest struct {
		LongUrl string `json:"long_url"`
	}
	var input CreateRequest
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	generatedUrl, err := h.services.Create(input.LongUrl)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"shortUrl": generatedUrl,
	})
}

func (h *Handler) find(c *gin.Context) {
	type FindRequest struct {
		ShortUrl string `uri:"short_url"`
	}
	var input FindRequest
	if err := c.BindUri(&input); err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	longUrl, err := h.services.Find(input.ShortUrl)
	if err != nil {
		logrus.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"long_url": longUrl,
	})
}
