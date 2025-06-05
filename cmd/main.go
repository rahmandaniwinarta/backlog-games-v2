package main

import (
	"fmt"
	"net/http"

	"backlog-games-v2/internal/handler"
)

func main() {
	http.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetGames(w, r)
		case http.MethodPost:
			handler.PostGames(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	fmt.Println("server running at http://localhost: " + port)
	http.ListenAndServe(port, nil)
}
