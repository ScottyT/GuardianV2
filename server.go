package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("dist"))

	http.Handle("/", fs)
	fmt.Printf("Starting server at port 8080\n")
	//	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
