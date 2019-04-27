package controller

import (
	"Url-Shortener/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortenUrl(c *gin.Context) {
	url := c.PostForm("url")
	res := model.ShortenUrl(url)
	c.JSON(200, gin.H{
		"success":  true,
		"res": res,
	})
}

func GetOriginUrl(c *gin.Context) {
	url := c.Params.ByName("url")
	res, ok := model.SearchOriginUrl(url)
	c.JSON(http.StatusOK, gin.H{"url": res, "success": ok})
}
