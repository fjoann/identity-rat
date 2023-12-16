package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type requestBody struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body requestBody
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Hi %s", body.Name),
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
