package controller

import (
	"Url-Shortener/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortenUrl(c *gin.Context) {
	url := c.PostForm("url")
	res := model.ShortenUrl(url)

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"res": res,
	})
}

func GetOriginUrl(c *gin.Context) {
	url := c.Params.ByName("url")
	res, ok := model.SearchOriginUrl(url)
	res = model.AddScript(res)
	if ok {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, res)
	} else {
		c.String(http.StatusNotFound, "")
	}
}
