package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Println("You submitted a form")
	lambda.Start(Handler)
	// from := mail.NewEmail("Example User", "test@example.com")
	// subject := "Sending with Twilio SendGrid is Fun"
	// to := mail.NewEmail("Example User", "stuck04@gmail.com")
	// plainTextContent := "and easy to do anywhere, even with Go"
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// client := sendgrid.NewSendClient(os.Getenv("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"))
	// response, err := client.Send(message)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(response.StatusCode)
	// 	fmt.Println(response.Body)
	// 	fmt.Println(response.Headers)
	// }
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello World!",
	}, nil
}
