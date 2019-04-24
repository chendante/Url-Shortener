package main

import (
	"Url-Shortener/router"
)

func main() {
	r := router.InitRouter()

	_ = r.Run("localhost:8000")
}
