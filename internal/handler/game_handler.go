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

var games []Game

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
			front := games[:i]
			back := games[i+1]
			games = append(front, back)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Game Deleted"))
			return
		}
	}

	http.Error(w, "Game not found", http.StatusNotFound)

}
