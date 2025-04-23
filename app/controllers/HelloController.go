package controllers

import "github.com/gin-gonic/gin"

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
