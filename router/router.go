package router

import "github.com/gin-gonic/gin"
import . "Url-Shortener/controller"

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/shorten-url", ShortenUrl)
	router.GET("/origin-url/:url", GetOriginUrl)
	return router
}
