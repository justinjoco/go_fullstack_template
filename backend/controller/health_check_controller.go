package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (c *HealthCheckController) HealthCheck(ctx *gin.Context) {
	log.Println("Health check")
	ctx.Status(http.StatusOK)
}
