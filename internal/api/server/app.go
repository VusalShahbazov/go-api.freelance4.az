package server

import (
	"github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"github.com/VusalShahbazov/api.freelance4.az/internal/services/migrations"
	"github.com/VusalShahbazov/api.freelance4.az/internal/services/validations"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func (s *ApiServer) Run() error {

	s.Engine =  gin.Default()

	time.Sleep(10 * time.Second)
	models.ConnectToDatabase()

	s.InitRoutes()

	if err := validations.RegisterCustomValidation(); err != nil {
		return err
	}

	if err := migrations.Run(os.Getenv("MIGRATION_DIR")); err != nil {
		return err
	}

	return s.Engine.Run(os.Getenv("API_PORT"))
}
