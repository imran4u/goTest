package main

import (
	handler "imran4u/goTest/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", handler.PingHandler)

	r.Run(":9003")

}
