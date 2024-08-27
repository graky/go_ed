package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"go_ed/pkg/config"
	"go_ed/pkg/handlers"
	"go_ed/pkg/middleware"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *sql.DB
}

func New(cfg *config.Config, db *sql.DB) *Server {
	s := &Server{
		router: gin.Default(),
		cfg:    cfg,
		db:     db,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.POST("/sign-up", handlers.SignUp(s.db))
	s.router.POST("/sign-in", handlers.SignIn(s.db))
	s.router.POST("/upload", middleware.AuthMiddleware(), handlers.Upload(s.db))
	s.router.GET("/files", middleware.AuthMiddleware(), handlers.GetFiles(s.db))
}

func (s *Server) Run() error {
	return s.router.Run(s.cfg.ServerAddress)
}