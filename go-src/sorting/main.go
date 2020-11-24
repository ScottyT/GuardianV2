package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Project struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Client string `json:"client"`
}
type By func(p1, p2 *Project) bool
type favSorter struct {
	by        By
	Option    string    `json:"option"`
	Favorites []Project `json:"project"`
}

//type By func(p1, p2 *Project)

func (by By) Sort(favs []Project) {
	ps := &favSorter{
		Favorites: favs,
		by:        by,
	}
	sort.Sort(ps)
}

func (a *favSorter) Len() int { return len(a.Favorites) }

func (a *favSorter) Swap(i, j int) { a.Favorites[i], a.Favorites[j] = a.Favorites[j], a.Favorites[i] }
func (a *favSorter) Less(i, j int) bool {
	//return a[i].Name < a[j].Name
	return a.by(&a.Favorites[i], &a.Favorites[j])
}

func sorting(w http.ResponseWriter, r *http.Request) {
	var data favSorter

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &data)
	// json.Unmarshal([]byte(reqBody), &proj)

	nameasc := func(p1, p2 *Project) bool {
		return p1.Name < p2.Name
	}
	namedesc := func(p1, p2 *Project) bool {
		return nameasc(p2, p1)
	}
	catasc := func(p1, p2 *Project) bool {
		return p1.Type < p2.Type
	}
	catdesc := func(p1, p2 *Project) bool {
		return catasc(p2, p1)
	}

	switch {
	case data.Option == "nameAsc":
		By(nameasc).Sort(data.Favorites)
	case data.Option == "nameDesc":
		By(namedesc).Sort(data.Favorites)
	case data.Option == "typeAsc":
		By(catasc).Sort(data.Favorites)
	case data.Option == "typeDesc":
		By(catdesc).Sort(data.Favorites)
	}
	json.NewEncoder(w).Encode(data)
}
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Web!\n"))
}
func main() {
	// lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// 	resp, err := handler(request)
	// 	return resp, err
	// })
	fs := http.FileServer(http.Dir("./dist"))
	/* http.Handle("/", fs)
	http.HandleFunc("/sorting", sorting)
	fmt.Printf("Starting server at port 8080\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":8080", nil)) */
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/hello", TestHandler)
	mux.HandleFunc("/sorting", sorting)
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Println("Hello world")
	// }
	// http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("server is listening at :8081")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":8081", mux))

}
