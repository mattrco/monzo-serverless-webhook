package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       http.StatusText(http.StatusMethodNotAllowed),
		}, nil
	}

	webhookBody := WebhookBody{}
	if err := json.Unmarshal([]byte(req.Body), &webhookBody); err != nil {
		log.Print(err)
		log.Printf("%+v", req.Body) // Body is a string so can be reused
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	log.Printf("%+v", webhookBody)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       http.StatusText(http.StatusAccepted),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
