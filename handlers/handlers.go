package handlers

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func Handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Service 2!")
}
