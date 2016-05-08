package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Service) healthCheckHandler(c *gin.Context) {
	host, err := os.Hostname()
	if err != nil {
		c.String(
			http.StatusServiceUnavailable,
			err.Error(),
		)
	}
	status := fmt.Sprintf("Service up and running on: %s\n",
		host,
	)
	c.String(
		http.StatusOK,
		status,
	)
}
