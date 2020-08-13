package main

import (
	"context"
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

//
type Body struct {
	Payload Payload `json:"payload"`
}

func HandleRequest(ctx context.Context, form Form) (string, error) {
	return fmt.Sprintf("Hello %s!", form.Name), nil
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body Body
	json.Unmarshal([]byte(request.Body), &body)
	// Send email
	from := mail.NewEmail(body.Payload.Form.Name, body.Payload.Form.Address)
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example user", "stuck04@gmail.com")
	plainTextContent := "Hello there."
	htmlContent := "<p style='font-size:16px;'>Hello there.</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
	}
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
	// name := r.FormValue("name")
	// address := r.FormValue("address")
	// var body Body

	// fields := Form{
	// 	Name:    name,
	// 	Address: address,
	// }

	// resp, err := http.Get("https://restorewithguardian.com/.netlify/functions/sendform")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(string(body))

	//fmt.Fprintf(w, "POST request successful")
	// name := r.FormValue("name")
	// address := r.FormValue("address")
	//v := url.Values{}

	// Send email
	// from := mail.NewEmail(name, address)
	// subject := "Sending with Twilio SendGrid is Fun"
	// to := mail.NewEmail("Example user", "stuck04@gmail.com")
	// plainTextContent := "Hello there."
	// htmlContent := "<p style='font-size:16px;'>Hello there.</p>"
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// client := sendgrid.NewSendClient("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")
	// response, err := client.Send(message)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(response.StatusCode)
	// 	fmt.Println(response.Body)
	// 	fmt.Println(response.Headers)
	// }
}

// func altSubmission(w http.ResponseWriter, r *http.Request) {
// 	requestBody, err := json.Marshal(map[string]string{
// 		"name":    "Scott",
// 		"address": "swtucker@twc.com",
// 	})
// 	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(string(body))
// }

func main() {
	lambda.Start(handler)
	// fileServer := http.FileServer(http.Dir("dist"))
	// http.Handle("/", fileServer)
	// fmt.Printf("Starting server at port 8080\n")
	// //	defer fmt.Println("Server ended")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// //http.HandleFunc("/thankyou", formHandler)
	// http.HandleFunc("/thankyou", altSubmission)
}
