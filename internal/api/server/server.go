package server

import "github.com/gin-gonic/gin"

type ApiServer struct {
	Engine *gin.Engine
	Config Config
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		Engine: gin.Default(),
		Config: Config{},
	}
}
