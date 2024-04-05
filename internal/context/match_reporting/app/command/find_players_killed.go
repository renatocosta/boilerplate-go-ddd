package command

import (
	"context"
	"encoding/json"
	"fmt"

	players_killed "github.com/ddd/internal/context/match_reporting/domain/model/player_killed"
	"github.com/ddd/pkg/building_blocks/app"
)

type FindPlayersKilledCommand struct {
	Data [][]string
}

type FindPlayersKilledHandler app.CommandHandler[FindPlayersKilledCommand, string]

type findPlayersKilledHandler struct {
	playerKilled players_killed.PlayerKilled
}

func NewFindPlayersKilledHandler(playerKilled players_killed.PlayerKilled) FindPlayersKilledHandler {
	return findPlayersKilledHandler{playerKilled: playerKilled}
}

func (h findPlayersKilledHandler) Handle(ctx context.Context, cmd FindPlayersKilledCommand) (string, error) {

	for _, row := range cmd.Data {
		h.playerKilled.Compute(players_killed.NewMatcher(row[0], row[1]))
	}

	entity := players_killed.NewGame(h.playerKilled.GetPlayers())
	entity.Sort()
	gameInfo := players_killed.GameInfo{
		TotalKills: entity.TotalKills(),
		Players:    entity.PlayerNames(),
		Kills:      entity.PlayerKills(),
	}
	jsonData, err := json.MarshalIndent(map[string]players_killed.GameInfo{"game_1": gameInfo}, "", "")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", nil
	}

	return string(jsonData), nil
}
