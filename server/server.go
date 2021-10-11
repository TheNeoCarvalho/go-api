package server

import (
	"log"

	"github.com/TheNeoCarvalho/go-api/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigueRoutes(s.server)

	log.Print("Server is running... ")
	log.Fatal(router.Run(":" + s.port))
}
