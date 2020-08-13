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

type Form struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

//proxy
type ProxyRequest struct {
	Headers     map[string]string `json:"headers"`
	Method      string            `json:"method"`
	Body        string            `json:"body"`
	QueryParams map[string]string `json:"queryParams"`
}

var defaultResponse = &events.APIGatewayProxyResponse{StatusCode: 200, Body: "success!"}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// p := ProxyRequest{
	// 	Headers:     request.Headers,
	// 	Method:      request.HTTPMethod,
	// 	Body:        request.Body,
	// 	QueryParams: request.QueryStringParameters,
	// }
	// b, err := json.Marshal(&p)
	// if err != nil {
	// 	log.Printf("failed to marshal event to JSON: %v", err)
	// 	return defaultResponse, nil
	// }
	dat := make(map[string]string)
	json.Unmarshal([]byte(request.Body), &dat)
	// Send email
	from := mail.NewEmail(dat["name"], dat["address"])
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
		log.Println(response.StatusCode)
	}
	return defaultResponse, nil

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
	lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		resp, err := handler(request)
		return resp, err
	})
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
