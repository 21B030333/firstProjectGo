package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)


type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Destiny string `json:"destiny"`
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/items", Items).Methods("GET")

	router.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		id := vars["id"]
		i, err := strconv.Atoi(id)
		if err != nil || i < 1 || i > 4{
			fmt.Fprintf(w, `<p>404 page not found</p>`)
			return
		}

		items := prepareResponse()

		fmt.Println(items[i - 1])

		if i == 1 {
			http.ServeFile(w, r, "./ui/html/alma.html")
		} else if i == 2{
			http.ServeFile(w, r, "./ui/html/book.html")
		} else if i == 3{
			http.ServeFile(w, r, "./ui/html/money.html")
		} else {
			http.ServeFile(w, r, "./ui/html/dombra.html")
		}
		
	})

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
