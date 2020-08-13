package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("dist"))
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port 8080\n")
	//	defer fmt.Println("Server ended")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
