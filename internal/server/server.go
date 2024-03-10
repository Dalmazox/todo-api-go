package server

import (
	"context"
	"log"
	"net/http"

	_ "github.com/dalmazox/todo-api/api"
	"github.com/dalmazox/todo-api/configs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start()
}

type ServerOptions struct {
	Context context.Context
}

type server struct {
	router  *gin.Engine
	options ServerOptions
}

func NewServer(options ServerOptions) Server {
	return server{
		router:  gin.New(),
		options: options,
	}
}

func (s server) Start() {
	s.initialSetup()
	s.addSwagger()
	s.addRoutes()
	s.start()
}

func (s server) initialSetup() {
	s.router.SetTrustedProxies(nil)
}

func (s server) addRoutes() {
	SetupRoutes(s.router)
}

func (s server) addSwagger() {
	if configs.GetConfig().Server.UseSwagger {
		s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (s server) start() {
	port := configs.GetConfig().Server.Port
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}

	log.Printf("Server is running on port %s", port)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
