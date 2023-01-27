package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	r := gin.New()
	r.Use(Logger(), gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		all, _ := io.ReadAll(rdr1)
		c.JSON(200, string(all))
	})
	r.Run("0.0.0.0:8888")
}
