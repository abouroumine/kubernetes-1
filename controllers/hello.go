package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct{}

func (h *HelloController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello World",
		"Detail": gin.H{
			"Detail 1": "Hello Again",
		},
	})
}

func (h *HelloController) ReceiveWorld(c *gin.Context) {
	word := c.PostForm("word")
	c.JSON(http.StatusOK, gin.H{
		"Message": word,
	})
}
