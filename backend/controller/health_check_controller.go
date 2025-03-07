package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController interface {
	HealthCheck(ctx *gin.Context)
}

type healthCheckController struct{}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (c *healthCheckController) HealthCheck(ctx *gin.Context) {
	log.Println("Health check")
	ctx.Status(http.StatusOK)
}
