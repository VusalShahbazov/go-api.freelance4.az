package v1

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/helper"
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"github.com/VusalShahbazov/api.freelance4.az/internal/repositories/userRepository"
	"github.com/VusalShahbazov/api.freelance4.az/internal/requests"
	"github.com/VusalShahbazov/api.freelance4.az/internal/response"
	"github.com/VusalShahbazov/api.freelance4.az/internal/services/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)


func  SignUp(ctx *gin.Context)  {
	var json requests.SignUpRequest

	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(ctx, err)
		return
	}

	exists , err := userRepository.CheckEmailIsExists(json.Email)
	if err != nil {
		response.InternalServerError(ctx , nil)
		return
	}

	if exists {
		//todo change to multi lang
		response.ValidationErrorByKey(ctx , "email" , []string{"Email Must be unique"})
		return
	}

	user, err := userRepository.CreateUser(&json)
	if err != nil {
		response.ErrorResponse(ctx , gin.H{"message" : err.Error()} , 400)
		return
	}

	responseWithToken(ctx, user)
}

func  Login(ctx *gin.Context)  {
	var json requests.Login
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		response.GenerateValidationResponse(ctx, err)
		return
	}
	user, err := userRepository.GetUserByEmail(json.Email)
	if err != nil {
		response.ErrorResponse(ctx , gin.H{"message" : err.Error()} , 400)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		response.Unauthorized(ctx , nil)
		return
	}
	responseWithToken(ctx, user)
}

func  Me(ctx *gin.Context) {
	get, exists := ctx.Get("user")
	if exists {
		response.SuccessResponse(ctx , gin.H{"user":get})
		return
	}
	response.ErrorResponse(ctx , gin.H{"message":"Unexpected error"} , http.StatusInternalServerError)
}

func  Logout(ctx *gin.Context) {
	_, exists := ctx.Get("user")
	if exists {
		ctx.Set("user" , nil)
		response.SuccessResponse(ctx , gin.H{"message"  :"Successful logged out"})
		return
	}
	response.ErrorResponse(ctx , gin.H{"message":"Unexpected error"} , http.StatusInternalServerError)
}

func ResetPassword(ctx *gin.Context)  {
	var json requests.ResetPassword
	if err := ctx.ShouldBindJSON(&json) ; err != nil {
		response.GenerateValidationResponse(ctx, err)
		return
	}
	user, err := helper.GetUser(ctx)
	if err != nil {
		response.UnexpectedError(ctx , err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.OldPassword))
	if err != nil {
		response.BadRequest(ctx , gin.H{"old_password" : "Password is not valid"})
		return
	}
	err = userRepository.ResetPassword(user, json.NewPassword)
	if err != nil {
		response.InternalServerError(ctx , err.Error())
		return
	}
	response.Ok(ctx)
}

func responseWithToken(ctx  *gin.Context, user *models.User)  {
	token, err := jwt.GenerateTokenFromUser(user)
	if err != nil {
		response.ErrorResponse(ctx , gin.H{"message" : err.Error()} , 400)
		return
	}
	response.SuccessResponse(ctx , gin.H{"token" : token ,  "user" : user})
}