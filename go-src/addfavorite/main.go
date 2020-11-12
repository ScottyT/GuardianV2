package main

import (
	"fmt"
	"log"
	"net/http"
)

func add(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fs := http.FileServer(http.Dir("./dist"))

	http.Handle("/", fs)
	fmt.Printf("Starting server at port 1000\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":1000", nil))
}
