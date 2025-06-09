package storage

import "backlog-games-v2/internal/model"

var Games = []model.Game{
	{ID: 1, Title: "The Witcher 3"},
	{ID: 2, Title: "Hollow Knight"},
	{ID: 3, Title: "Elden Ring"},
}

func GetAllGames() []model.Game {
	return Games
}

func AddGame(g model.Game) {
	Games = append(Games, g)
}

func DeleteGameByID(id int) bool {
	for i, game := range Games {
		if game.ID == id {
			Games = append(Games[:i], Games[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateGameByID(id int, updating model.Game) (*model.Game, bool) {
	for i, games := range Games {
		if games.ID == id {
			Games[i].Title = updating.Title
			return &Games[i], true
		}
	}
	return nil, false
}
