package server

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/controllers"
	"github.com/VusalShahbazov/api.freelance4.az/internal/controllers/v1"
	"github.com/VusalShahbazov/api.freelance4.az/internal/middleware"
)

func (s *ApiServer) InitRoutes() {
	s.Engine.GET("/", controllers.Index)

	api := s.Engine.Group("api")
	{
		v1Group := api.Group("v1")
		{
			auth := v1Group.Group("auth")
			{
				auth.POST("signup", v1.SignUp)
				auth.POST("login", v1.Login)
				auth.POST("logout", middleware.Auth, v1.Logout)
				auth.GET("me", middleware.Auth, v1.Me)
				password := auth.Group("password")
				{
					password.POST("reset" , middleware.Auth ,  v1.ResetPassword)
					//password.POST("forget" , middleware.Auth ,  v1.ResetPassword)
				}
			}
		}
	}
}
