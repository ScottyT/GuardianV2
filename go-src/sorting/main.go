package main

import (
	"fmt"
	"log"
	"net/http"
)

type Project struct {
	Name   string `json:"project"`
	Type   string `json:"type"`
	Client string `json:"client"`
}
type FormData struct {
	Option    string                 `json:"option"`
	Favorites map[string]interface{} `json:"project"`
}
type ByName []Project

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func sorting(w http.ResponseWriter, r *http.Request) {
	//var data FormData
	//	var proj Project
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// json.Unmarshal([]byte(reqBody), &data)
	//json.Unmarshal([]byte(reqBody), &proj)

	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	fmt.Println(r.FormValue("project"))
	//fmt.Println(data)
	//sort.Sort(ByName(proj))

	// fmt.Println(proj)
}
func main() {
	// lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// 	resp, err := handler(request)
	// 	return resp, err
	// })
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.HandleFunc("/sorting", sorting)
	fmt.Printf("Starting server at port 1000\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":1000", nil))
}
