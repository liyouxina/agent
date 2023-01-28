package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func handle(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	if body != nil {
		fmt.Println(string(body))
	}
	c.JSON(200, nil)
}

func main() {
	r := gin.Default()
	r.GET("/event", handle)
	r.POST("/event", handle)
	r.Run("0.0.0.0:8899")
}