package clients

import (
	"os"

	"github.com/gin-gonic/gin"
)

// End emula la recepción de un request al API y responde con el resultado de la petición
// hecha en el handler anterior.
func End(c *gin.Context) {

	clientResponse := c.GetInt("clientResponse")
	prov := os.Getenv("WEBHOOK_URL")
	description := `
		Lo que se amula aquí es:
		1. Tú eres eres el Cliente e hiciste un post a la lambda API.
		2. Este API hizo una petición al proveedor (` + prov + `) y,
		3. La IP registrada en el webhook.site deberá ser la IP del NAT-Gateway
	 `
	c.JSON(200, gin.H{
		"lambda_mood":     "api",
		"client_response": clientResponse,
		"description":     description,
	})
}
