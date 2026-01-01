package main

import (
	"QAProject/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.ListenAndServe(":8080", nil)
}
