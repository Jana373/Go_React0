package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func get_text() string {
	return "this is a test text"
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	response := Response{Message: get_text()}
	json.NewEncoder(w).Encode(response)
}
func main() {
	fs := http.FileServer(http.Dir("./build"))
	http.Handle("/", fs)

	http.HandleFunc("/api/quote", textHandler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
