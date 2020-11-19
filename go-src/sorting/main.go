package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Project struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Client string `json:"client"`
}

// type ByName []Project

// func (a ByName) Len() int           { return len(a) }
// func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func sort(w http.ResponseWriter, r *http.Request) {
	var proj Project
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &proj)
	fmt.Println(proj)
	//sort.Sort(ByName(proj))

	fmt.Println(proj)
}

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.HandleFunc("/sorting", sort)
	fmt.Printf("Starting server at port 1000\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":1000", nil))
}
