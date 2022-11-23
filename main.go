package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/kalmecak/aws-lambda-test/api"
)

var ginLambda *ginadapter.GinLambda

// Handler majea las peticiones de las Lambdas a Gin
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {

	g := api.Router()

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		ginLambda = ginadapter.New(g)

		lambda.Start(Handler)
	} else {
		if err := g.Run(":8080"); err != nil {
			panic(err)
		}
	}

}
