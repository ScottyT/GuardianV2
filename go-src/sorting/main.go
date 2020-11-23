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
	var proj Project
	res1, _ := json.Marshal(data.Favorites)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &data)
	json.Unmarshal([]byte(res1), &proj)
	fmt.Println(proj)
	// json.Unmarshal([]byte(reqBody), &proj)

	//fmt.Println(r.FormValue("project"))
	// switch {
	// case data.Option == "nameAsc":
	// 	sort.Slice(data.Favorites, func(i, j int) bool { return data.Favorites[i].Name < data.Favorites[j].Name })
	// case data.Option == "nameDesc":
	// 	sort.Slice(data.Favorites, func(i, j int) bool { return data.Favorites[i].Name > data.Favorites[j].Name })
	// }

	nameasc := func(p1, p2 *Project) bool {
		return p1.Name < p2.Name
	}
	namedesc := func(p1, p2 *Project) bool {
		return p1.Name > p2.Name
	}
	cat := func(p1, p2 *Project) bool {
		return p1.Type < p2.Type
	}

	switch {
	case data.Option == "nameAsc":
		By(nameasc).Sort(data.Favorites)
	case data.Option == "nameDesc":
		By(namedesc).Sort(data.Favorites)
	case data.Option == "typeAsc":
		By(cat).Sort(data.Favorites)

	}
	json.NewEncoder(w).Encode(data)
	//json.NewEncoder(w).Encode(data)
}

func main() {
	// lambda.Start(func(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// 	resp, err := handler(request)
	// 	return resp, err
	// })
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.HandleFunc("/sorting", sorting)
	fmt.Printf("Starting server at port 8080\n")
	defer fmt.Println("Server ended")
	log.Fatal(http.ListenAndServe(":8080", nil))
	/* addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", TestHandler) */
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Println("Hello world")
	// }
	// http.HandleFunc("/hello", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	/* log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux)) */

}
