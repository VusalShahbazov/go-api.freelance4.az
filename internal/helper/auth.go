package helper

import (
	"errors"
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*models.User, error) {
	stUser, ok := ctx.Get("user")
	if !ok {
		return nil, errors.New("User not exists")
	}
	user, ok := stUser.(models.User)
	if !ok {
		return nil, errors.New("User key is not valid")
	}
	return &user, nil
}
