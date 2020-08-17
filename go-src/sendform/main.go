package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MyForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
type MyPayload struct {
	MyForm MyForm `json:"data"`
}

//proxy
type ProxyRequest struct {
	Headers     map[string]string `json:"headers"`
	Method      string            `json:"method"`
	Body        string            `json:"body"`
	QueryParams map[string]string `json:"queryParams"`
}

var defaultResponse = &events.APIGatewayProxyResponse{StatusCode: 200, Body: "success!"}

var apiKey = os.Getenv("SENDGRID_API_KEY")

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// p := ProxyRequest{
	// 	Headers:     request.Headers,
	// 	Method:      request.HTTPMethod,
	// 	Body:        request.Body,
	// 	QueryParams: request.QueryStringParameters,
	// }

	// dat := make(map[string]string)
	// json.Unmarshal([]byte(request.Body), &dat)
	var dat MyForm

	json.Unmarshal([]byte(request.Body), &dat)
	// Send email
	from := mail.NewEmail(dat.Name, dat.Email)
	subject := "Contact Form Submission"
	to := mail.NewEmail("Headquarters", "stuck04@gmail.com")
	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", dat.Name, dat.Email, dat.Message)
	htmlContent := fmt.Sprintf("<p style='font-size:16px;'>Name: %s</p>\n<p style='font-size:16px;'>Email: %s</p>\n<p style='font-size:16px'>Message: %s</p>", dat.Name, dat.Email, dat.Message)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
	}
	return defaultResponse, nil

}

// Use for local testing
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	var dat MyForm
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &dat)

	// Send email
	from := mail.NewEmail(dat.Name, dat.Email)
	subject := "Contact Form Submission"
	to := mail.NewEmail("Headquarters", "stuck04@gmail.com")
	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", dat.Name, dat.Email, dat.Message)
	htmlContent := fmt.Sprintf("<p style='font-size:16px;'>Name: %s</p>\n<p style='font-size:16px;'>Email: %s</p>\n<p style='font-size:16px'>Message: %s</p>", dat.Name, dat.Email, dat.Message)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Fprintf(w, "Thank you for contacting us! We will reach out to you shortly.")
	}
}
func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		resp, err := handler(request)
		return resp, err
	})
	// fs := http.FileServer(http.Dir("./dist"))

	// http.Handle("/", fs)
	// http.HandleFunc("/thankyou", formHandler)
	// fmt.Printf("Starting server at port 1000\n")
	// //	defer fmt.Println("Server ended")
	// log.Fatal(http.ListenAndServe(":1000", nil))

	//log.Fatal(http.ListenAndServe(":1000", r))
	// http.HandleFunc("/thankyou", altSubmission)
}
