package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("about go MOd")

	greeter()
	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hello mod users")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Hello Mod users here</h1>"))
}
