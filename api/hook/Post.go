package hook

import "github.com/gin-gonic/gin"

// Post emula la recepción de un post al hook, envía un post a un webhook y responderá
// con el resultado de la petición.
func Post(c *gin.Context) {

	clientResponse := c.GetInt("clientResponse")
	description := `
		Lo que se amula aquí es:
		1. Tú eres eres el proveedor e hiciste un post a la IP del NAT-Gateway,
		2. El NAT-Gateway redireccionó la petición a una lambda que ejecuta este código,
		3. Esta respuesta debería tener la IP del Nat-Gateway como "origen".
	 `
	c.JSON(200, gin.H{
		"lambda_mood":     "hook",
		"client_response": clientResponse,
		"description":     description,
	})
}
