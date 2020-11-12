package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var rxEmail = regexp.MustCompile(`^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`)
var rxName = regexp.MustCompile(`^[a-zA-Z\s]*$`)

type MyForm struct {
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Message     string                 `json:"message"`
	Project     string                 `json:"project"`
	Client      string                 `json:"client"`
	Type        string                 `json:"type"`
	DateRange   map[string]interface{} `json:"dateRange`
	Description string                 `json:"description"`
	FormType    string                 `json:"formtype"`
	Errors      map[string]string
}

type MyPayload struct {
	MyForm MyForm `json:"data"`
}

const apiKey = "SENDGRID_API_KEY"

func (msg *MyForm) Validate() bool {
	msg.Errors = make(map[string]string)
	formType := msg.FormType
	switch {
	case formType == "ProjectCreate":
		if strings.TrimSpace(msg.Project) == "" {
			msg.Errors["Project"] = "Please give the project a name"
		}
		if strings.TrimSpace(msg.Client) == "" {
			msg.Errors["Client"] = "Please specify a client"
		}
		if strings.TrimSpace(msg.Type) == "" {
			msg.Errors["Type"] = "Please enter a category"
		}
		if msg.DateRange["endDate"] == nil && msg.DateRange["startDate"] == nil {
			msg.Errors["DateRange"] = "Please select the start and end dates"
		}
		return len(msg.Errors) == 0
	case formType == "Contact":
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
	return false
}

func (msg *MyForm) Deliver() []byte {
	formType := msg.FormType
	switch {
	case formType == "Contact":
		from := mail.NewEmail(msg.Name, msg.Email)
		subject := "Contact Form Submission"
		to := mail.NewEmail("Scott Tucker", "scott@damage.click")
		plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", msg.Name, msg.Email, msg.Message)
		htmlContent := fmt.Sprintf("<p style='font-size:16px;'>Name: %s</p>\n<p style='font-size:16px;'>Email: %s</p>\n<p style='font-size:16px'>Message: %s</p>", msg.Name, msg.Email, msg.Message)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		return mail.GetRequestBody(message)

	case formType == "ProjectCreate":
		from := mail.NewEmail("No-reply", "no-reply@guardian.com")
		subject := "New project submission"
		to := mail.NewEmail("Headquarters", "stuck04@gmail.com")
		plainTextContent := fmt.Sprintf("Project name: %s\nClient: %s\nDate: %s\nDescription: %s\n", msg.Project, msg.Client, msg.DateRange, msg.Description)
		htmlContent := fmt.Sprintf("<p style='font-size:16px;'>Project name: %s</p>\n<p style='font-size:16px;'>Client: %s</p>\n<p style='font-size:16px'>Date: %s</p>\n<p style='font-size:16px'>Description: %s</p>", msg.Project, msg.Client, msg.DateRange, msg.Description)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		return mail.GetRequestBody(message)
	}
	return nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
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
	var message string
	var status int
	if response.StatusCode != 200 && response.StatusCode != 202 {
		return &events.APIGatewayProxyResponse{
			StatusCode: response.StatusCode,
			Body:       "Sorry, something went wrong on our end!",
		}, nil
	}
	switch {
	case dat.FormType == "Contact":
		message = "Thank you for contacting us! We will reach out to you shortly"
		status = 202
	case dat.FormType == "ProjectCreate":
		message = "Project was added successfully!"
		status = 202
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       message,
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

// use for project form submission
func projectHandler(w http.ResponseWriter, r *http.Request) {
	var dat MyForm
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &dat)
	fmt.Println(reqBody)
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
		fmt.Fprintf(w, "Project was added successfully!")
	}
}
func main() {
	// lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// 	resp, err := handler(request)
	// 	return resp, err
	// })
	fs := http.FileServer(http.Dir("./dist"))

	http.Handle("/", fs)
	http.HandleFunc("/sendform", formHandler)
	http.HandleFunc("/create", projectHandler)
	fmt.Printf("Starting server at port 1000\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":1000", nil))
}
