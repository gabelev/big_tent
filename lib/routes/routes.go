package routes

import (
	"github.com/gabelev/get_heard/lib/data"

	"github.com/gin-gonic/gin"
)

type Service struct {
	host   string
	port   string
	db     data.DB
	router *gin.Engine
}

// New creates an instance of the Service.
func New(
	host string,
	port string,
	db data.DB,
	router *gin.Engine,
) *Service {
	return &Service{
		host:   host,
		port:   port,
		db:     db,
		router: router,
	}
}

// New Server Generates an instance of the Gin Server.
func NewServer() (router *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	// TODO: should this be New or Default?
	router = gin.Default()
	return router
}

func (s *Service) Start() (err error) {
	router := s.router
	router.GET("big_tent/hc", s.healthCheckHandler)
	err = router.Run(s.port)
	return err
}
