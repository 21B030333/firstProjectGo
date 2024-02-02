package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi! My name is Elnar.\n")
	fmt.Fprintf(w, "Tusaukeser is one of the ancient Kazakh traditions that has survived to this day,\n")
	fmt.Fprintf(w, "procedure is intended for a small child and is done so that the child can get back on his feet faster.")
	
}


func Items(w http.ResponseWriter, r *http.Request){
	var response Response

	items := prepareResponse()

	response.Items = items

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func prepareResponse() []Item {
	var items []Item

	var item Item
	item.Id = 1
	item.Name = "Apple"
	item.Destiny = "Full"
	items = append(items, item)

	item.Id = 2
	item.Name = "Book"
	item.Destiny = "Smart"
	items = append(items, item)

	item.Id = 3
	item.Name = "Money"
	item.Destiny = "Rich"
	items = append(items, item)

	item.Id = 4
	item.Name = "Dombra"
	item.Destiny = "Creative"
	items = append(items, item)

	return items
}