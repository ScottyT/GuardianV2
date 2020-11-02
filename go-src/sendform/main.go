package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var rxEmail = regexp.MustCompile(`^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`)
var rxName = regexp.MustCompile(`^[a-zA-Z\s]*$`)

type MyForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Errors  map[string]string
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

const apiKey = "SENDGRID_API_KEY"

func (msg *MyForm) Validate() bool {
	msg.Errors = make(map[string]string)
	matchEmail := rxEmail.MatchString(msg.Email)
	if matchEmail == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}
	matchName := rxName.MatchString(msg.Name)
	if matchName == false {
		msg.Errors["Name"] = "Name can only contain letters and spaces"
	}
	if len(msg.Name) >= 30 {
		msg.Errors["Name"] = "Name must be shorter than 30 characters"
	}
	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter a name"
	}
	if strings.TrimSpace(msg.Message) == "" {
		msg.Errors["Message"] = "Please enter a message"
	}
	return len(msg.Errors) == 0
}

func (msg *MyForm) Deliver() []byte {
	fmt.Println(msg.Name)
	from := mail.NewEmail(msg.Name, msg.Email)
	subject := "Contact Form Submission"
	to := mail.NewEmail("Headquarters", "scott@damage.click")
	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", msg.Name, msg.Email, msg.Message)
	htmlContent := fmt.Sprintf("<p style='font-size:16px;'>Name: %s</p>\n<p style='font-size:16px;'>Email: %s</p>\n<p style='font-size:16px'>Message: %s</p>", msg.Name, msg.Email, msg.Message)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	return mail.GetRequestBody(message)
}

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
	if dat.Validate() == false {
		bodyStr, _ := json.Marshal(dat)
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(bodyStr),
		}, nil
	}
	// Send email
	r := sendgrid.GetRequest(os.Getenv(apiKey), "/v3/mail/send", "https://api.sendgrid.com")
	r.Method = "POST"
	var Body = dat.Deliver()
	r.Body = Body
	response, _ := sendgrid.API(r)
	fmt.Println(response)
	if response.StatusCode != 200 && response.StatusCode != 202 {
		return &events.APIGatewayProxyResponse{
			StatusCode: response.StatusCode,
			Body:       "Sorry, something went wrong on our end!",
		}, nil
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 202,
		Body:       "Thank you for contacting us! We will reach out to you shortly",
	}, nil
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
	if dat.Validate() == false {
		json.NewEncoder(w).Encode(dat)
		return
	}

	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = dat.Deliver()
	request.Body = Body
	response, _ := sendgrid.API(request)
	if response.StatusCode != 200 && response.StatusCode != 202 {
		http.Error(w, "Sorry, something went wrong", response.StatusCode)
	} else {
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
	// http.HandleFunc("/sendform", formHandler)
	// fmt.Printf("Starting server at port 1000\n")
	// defer fmt.Println("Server ended")
	// log.Fatal(http.ListenAndServe(":1000", nil))
}
