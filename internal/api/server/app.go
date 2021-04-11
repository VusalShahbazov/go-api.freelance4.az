package server

import (
	"fmt"
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"github.com/VusalShahbazov/api.freelance4.az/internal/services/migrations"
	"github.com/VusalShahbazov/api.freelance4.az/internal/services/validations"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func (s *ApiServer) Run() error {

	//ginMode := os.Getenv("APP_MODE")
	//if ginMode == "" {
	//	ginMode = gin.ReleaseMode
	//}
	//gin.SetMode(ginMode)

	s.Engine =  gin.Default()

	time.Sleep(time.Second * 5)
	models.ConnectToDatabase()

	s.InitRoutes()

	if err := validations.RegisterCustomValidation(); err != nil {
		return err
	}

	if err := migrations.Run(os.Getenv("MIGRATION_DIR")); err != nil {
		return err
	}

	return s.Engine.Run(fmt.Sprintf(":%v" ,  os.Getenv("API_PORT")))
}
