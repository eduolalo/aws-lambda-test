package clients

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// ToProvider manda el body del request a un webhook emulando ser el proveedor
func ToProvider(c *gin.Context) {

	url := os.Getenv("WEBHOOK_URL")

	httpClient := &http.Client{}
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		log.Println("Err -> New Request", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Walabi-Hook-Service")

	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	// Generamos el post
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println("Err -> ClientDo", err)
		log.Println("Tip -> Recuerda haber dado de alta la variable de entorno WEBHOOK_URL")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Set("clientResponse", res.StatusCode)

	c.Next()
}
