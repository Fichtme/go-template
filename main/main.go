package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello world!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)
	_ = http.ListenAndServe(":8080", router)
}
