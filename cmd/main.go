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
			/*
				curl http://localhost:8080/games
			*/
		case http.MethodPost:
			handler.PostGames(w, r)
			/*
				 curl -X POST http://localhost:8080/games \
						-H "Content-Type: application/json" \
						-d '{"title": "Albion Online"}'
			*/
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/games/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			handler.DeleteGames(w, r)
			/*
				curl -X DELETE http://localhost:8080/games/1
			*/
		case http.MethodPut:
			handler.UpdateGames(w, r)
			/*
				curl -X PUT http://localhost:8080/games/1 \
						-H "Content-Type: application/json" \
				  		-d '{"title": "Elden Ring"}'
			*/
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	fmt.Println("server running at http://localhost: " + port)
	http.ListenAndServe(port, nil)
}
