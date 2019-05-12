package main

import (
	"encoding/json"
	_ "encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {

	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "First book", Author: "First Author", Year: "2011"},
		Book{ID: 2, Title: "Second book", Author: "Second Author", Year: "2012"},
		Book{ID: 3, Title: "Third book", Author: "Third Author", Year: "2013"},
		Book{ID: 4, Title: "Fourth book", Author: "Fourth Author", Year: "2014"},
		Book{ID: 5, Title: "Fifth book", Author: "Fifth Author", Year: "2015"})

	router.HandleFunc("/books", getbooks).Methods("GET")
	router.HandleFunc("/books/{id}", getbook).Methods("GET")
	router.HandleFunc("/books", addbook).Methods("POST")
	router.HandleFunc("/books", updatebook).Methods("PUT")
	router.HandleFunc("/books/{id}", removebook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getbooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(books)

}

func addbook(w http.ResponseWriter, r *http.Request) {

	log.Println("Add book")

}

func getbook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	log.Println(params) // params is a map

	log.Println(reflect.TypeOf(params["id"])) // usinng reflect library to get type of variable params["id"]

	i, _ := strconv.Atoi(params["id"]) // convert string to integer

	log.Println(reflect.TypeOf(i)) // get type of i

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}

}

func updatebook(w http.ResponseWriter, r *http.Request) {

	log.Println("Update book")

}

func removebook(w http.ResponseWriter, r *http.Request) {

	log.Println("Remove book")

}