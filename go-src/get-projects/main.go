package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	//ctx := context.Background()
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello",
	}, nil
}
func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		resp, err := handler(request)
		return resp, err
	})
}
