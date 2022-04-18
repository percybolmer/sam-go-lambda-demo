package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler(event events.SQSEvent) error {
	log.Println(event)

	return nil
}

func main() {
	lambda.Start(lambdaHandler)
}
