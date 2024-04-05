package human_logfile

import (
	"errors"
)

var (
	ErrMissingFields = errors.New("Missing fields")
)

func NewHumanLogFileRow(playerWhoKilled string, playerWhoDied string, meanOfDeath string) HumanLogFileRowable {
	return &HumanLogFileRow{
		playerWhoKilled: playerWhoKilled,
		playerWhoDied:   playerWhoDied,
		meanOfDeath:     meanOfDeath,
	}
}

type HumanLogFileRow struct {
	playerWhoKilled string
	playerWhoDied   string
	meanOfDeath     string
}

type HumanLogFileRowable interface {
	Validation() error
	GetPlayerWhoKilled() string
	GetPlayerWhoDied() string
	GetMeanOfDeath() string
}

func (h HumanLogFileRow) Validation() error {

	if h.playerWhoKilled == "" || h.playerWhoDied == "" || h.meanOfDeath == "" {
		return ErrMissingFields
	}

	return nil
}

func (h HumanLogFileRow) GetPlayerWhoKilled() string {
	return h.playerWhoKilled
}

func (h HumanLogFileRow) GetPlayerWhoDied() string {
	return h.playerWhoDied
}

func (h HumanLogFileRow) GetMeanOfDeath() string {
	return h.meanOfDeath
}
