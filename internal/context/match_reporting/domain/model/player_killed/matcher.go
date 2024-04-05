package players_killed

import "errors"

var (
	ErrMissingFields = errors.New("Missing fields")
)

type Matchable interface {
	GetPlayerWhoKilled() string
	GetPlayerWhoDied() string
	Validation() error
}

func NewMatcher(whoKilled string, whoDied string) Matchable {
	return Matcher{
		whoKilled: whoKilled,
		whoDied:   whoDied,
	}
}

type Matcher struct {
	whoDied   string
	whoKilled string
}

func (m Matcher) GetPlayerWhoKilled() string {
	return m.whoKilled
}

func (m Matcher) GetPlayerWhoDied() string {
	return m.whoDied
}

func (m Matcher) Validation() error {

	if m.whoKilled == "" || m.whoDied == "" {
		return ErrMissingFields
	}

	return nil
}
