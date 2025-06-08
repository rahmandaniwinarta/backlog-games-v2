package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Game struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var games = []Game{
	{ID: 1, Title: "Hollow Knight"},
	{ID: 2, Title: "Celeste"},
	{ID: 3, Title: "Elden Ring"},
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func PostGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	var NewGame Game

	err := json.NewDecoder(r.Body).Decode(&NewGame)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}

	NewGame.ID = len(games) + 1
	games = append(games, NewGame)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewGame)
}

func DeleteGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	deleteThis := strings.Split(r.URL.Path, "/") // DELETE /GAMES/3 --> []deleteThis{"","games","3"}

	if len(deleteThis) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	idStr := deleteThis[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	for i, game := range games {
		if game.ID == id {
			games = append(games[:i], games[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Game Deleted"))
			return
		}
	}
	http.Error(w, "Game not found", http.StatusNotFound)
}

func UpdateGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	//ambil ID dari path
	takeID := strings.Split(r.URL.Path, "/")
	if len(takeID) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}
	idStr := takeID[2] // PUT /games/2 index 0,1,2
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	//decode JSON body
	var updatedGames Game
	err = json.NewDecoder(r.Body).Decode(&updatedGames)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, game := range games {
		if game.ID == id {
			games[i].Title = updatedGames.Title

			w.Header().Set("Content Type", "application/json")
			json.NewEncoder(w).Encode(games[i])
			return
		}
	}
	http.Error(w, "Game not found", http.StatusNotFound)
}
