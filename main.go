package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	adocCfg, err := ParseAdocConfig("config.yml")
	if err != nil {
		panic(err)
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, adocCfg)
	})
	r.Run()
}
