package main

import "github.com/aws/aws-lambda-go/lambda"

var (
	alertmanagerBaseUrl string = "localhost:9093"
	silencesApiUrl      string = "/api/v2/silences/"
)

func main() {
	lambda.Start(handler)
}
