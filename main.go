package main

import (
	"bytes"
	"flag"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ResponseStatusRequest struct { // LOL name
	Code int `uri:"code" binding:"required"`
}

func main() {
	var bindAddress string
	flag.StringVar(&bindAddress, "bind", "", "0.0.0.0:8888")
	flag.Parse()
	if bindAddress == "" {
		bindAddress = ":8080"
	}

	r := gin.New()
	r.Use(Logger(), gin.Recovery())
	r.Any("/status/:code", func(c *gin.Context) {
		var requestedStatus ResponseStatusRequest
		err := c.ShouldBindUri(&requestedStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		all, _ := io.ReadAll(rdr1)
		c.JSON(requestedStatus.Code, string(all))
	})
	r.Any("/authenticated", func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		all, _ := io.ReadAll(rdr1)
		if auth == "" {
			c.JSON(http.StatusUnauthorized, string(all))
			return
		}
		c.JSON(http.StatusOK, string(all))
	})
	r.NoRoute(func(c *gin.Context) {
		buf, _ := io.ReadAll(c.Request.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		all, _ := io.ReadAll(rdr1)
		c.JSON(http.StatusOK, string(all))
	})
	r.Run(bindAddress)
}
