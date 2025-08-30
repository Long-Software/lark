package api

import (
	"github.com/Long-Software/Bex/cmd/shorten-url/internal/config"
	"github.com/Long-Software/Bex/packages/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Redis
	router *gin.Engine
	config config.Config
}

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	return config
}

func NewServer(config config.Config, store db.Redis) (*Server, error) {
	s := &Server{config: config, store: store}
	// Routes
	s.SetupRouter()
	return s, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (s *Server) SetupRouter() {
	r := gin.Default()
	r.Use(cors.New(corsConfig())) // Cors
	r.GET("/:url", s.ResolveURL)
	r.POST("/api/v1", s.ShortenURL)
}
