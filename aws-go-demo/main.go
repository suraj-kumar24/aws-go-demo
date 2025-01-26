package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message   string   `json:"message"`
	FruitList []string `json:"fruit_list"`
	Status    bool     `json:"status"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Message:   "Hello, World!",
			FruitList: []string{"apple", "banana", "cherry"},
			Status:    true,
		})
	})
	log.Println("Server is running on port 3005 ")
	if err := http.ListenAndServe(":3005", nil); err != nil {
		log.Fatal(err)
	}
}
