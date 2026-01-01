package main

import (
	"QAProject/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handler2)
	http.ListenAndServe(":8081", nil)
}
