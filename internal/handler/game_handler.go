package handler

import (
	"backlog-games-v2/internal/model"
	"backlog-games-v2/internal/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetAllGames())
}

func PostGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	//cek method pas curl kalau ga POST berarti no no

	var newGame model.Game
	err := json.NewDecoder(r.Body).Decode(&newGame) //decode body : misal{"title":"Pokemon"}. decode JSON nya salin ke struct Title karena di struct sudah `json:title`
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	newGame.ID = len(storage.GetAllGames()) + 1 //set ID jadi sekarang struct newGames(id +1)
	storage.AddGame(newGame)                    // new game id + 1 , title dapet dari body

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGame)
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

	if storage.DeleteGameByID(id) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Game Deleted"))
		return
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
	var updatedGames model.Game
	err = json.NewDecoder(r.Body).Decode(&updatedGames)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	updated, ok := storage.UpdateGameByID(id, updatedGames)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
		return
	}
	http.Error(w, "Game not found", http.StatusNotFound)
}
