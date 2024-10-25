package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Define the handler function
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Example: Parse a query string parameter
	name := req.QueryStringParameters["name"]
	if name == "" {
		name = "World" // Default value
	}

	// Create a response message
	message := fmt.Sprintf("Hello, %s!", name)

	// Prepare the API Gateway response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       message,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
