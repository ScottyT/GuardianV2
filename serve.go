package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func main() {
	fmt.Println("Starting...")
	//fs := http.FileServer(http.Dir("/dist"))
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello world")
	}
	//http.Handle("/", fs)
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	/* fmt.Printf("Starting server at port 8080\n")
	//	defer fmt.Println("Server ended")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil) */
	//log.Fatal(http.ListenAndServe(":8080", nil))
	/* _, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("Connection error: %s\n", err)
	} else {
		fmt.Println("Listening")
	} */
}
