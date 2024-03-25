// go run *.go
// http://localhost:8080/score

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Score int `json:"score"`
}

func scoreHandler(w http.ResponseWriter, r *http.Request) {
	results, err := getData()
	if err != nil {
		log.Printf("Error getting API counts: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	score := getScore(results)

	response := Response{
		Score: score,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/score", scoreHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server start failed, err:", err)
		return
	}
}
