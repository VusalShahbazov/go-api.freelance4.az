package controllers

import (
	"github.com/gin-gonic/gin"
)

func  Index(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{"message":"Maybe word :D"})
}
func Test(ctx *gin.Context)  {
	ctx.JSON(200 , gin.H{"message":"text"})
}

