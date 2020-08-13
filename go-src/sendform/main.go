package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//fsdf
type Form struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// type
type Payload struct {
	Form Form `json:"data"`
}

// type
type Body struct {
	Payload Payload `json:"payload"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body Body
	json.Unmarshal([]byte(request.Body), &body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "success!",
	}, nil
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	// name := r.FormValue("name")
	// address := r.FormValue("address")
	name := r.PostFormValue("name")
	address := r.PostFormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	// Send email
	from := mail.NewEmail("Example User", address)
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example user", "stuck04@gmail.com")
	plainTextContent := "Hello there."
	htmlContent := "<p style='font-size:16px;'>Hello there.</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

// func altFormSub() {
// 	formData :=
// }
func main() {
	fileServer := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fileServer)
	//http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	lambda.Start(handler)
}
