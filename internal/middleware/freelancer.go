package middleware

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/helper"
	"github.com/VusalShahbazov/api.freelance4.az/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Customer(ctx *gin.Context) {
	user, err := helper.GetUser(ctx)
	if err != nil {
		response.UnexpectedError(ctx, nil)
		ctx.Abort()
		return
	}
	if user.IsFreelancer {
		response.ErrorResponse(ctx, gin.H{"message": "Have to be customer"}, http.StatusForbidden)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func Freelancer(ctx *gin.Context) {
	user, err := helper.GetUser(ctx)
	if err != nil {
		response.UnexpectedError(ctx, nil)
		ctx.Abort()
		return
	}
	if !user.IsFreelancer {
		response.ErrorResponse(ctx, gin.H{"message": "Have to be freelancer"}, http.StatusForbidden)
		ctx.Abort()
		return
	}

	ctx.Next()
}

