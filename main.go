package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func lambdaHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var input Event

	if err := json.Unmarshal([]byte(event.Body), &input); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	output := fmt.Sprintf("Hello from Lambda %s, the variable is %s", input.Name, os.Getenv("my-cool-variable"))

	return events.APIGatewayProxyResponse{
		Body:       output,
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
