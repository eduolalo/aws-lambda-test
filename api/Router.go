package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kalmecak/aws-lambda-test/api/clients"
	"github.com/kalmecak/aws-lambda-test/api/hook"
)

func Router() *gin.Engine {
	g := gin.Default()

	switch os.Getenv("LAMBDA_MOOD") {
	case "api":
		// Rutas para probar el hook
		g.DELETE("/api", clients.ToProvider, clients.End)
		g.POST("/api", clients.ToProvider, clients.End)
		g.GET("/api", clients.ToProvider, clients.End)
		g.PUT("/api", clients.ToProvider, clients.End)
	case "hook":
		// Ruta para probar el hook
		g.POST("/hook", hook.Post)
	default:
		// Ruta de default para probar que el servicio está funcionando
		g.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}

	return g
}
