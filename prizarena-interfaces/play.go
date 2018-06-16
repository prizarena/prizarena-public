package prizarena_interfaces

import "github.com/prizarena/prizarena-public/pamodels"

//go:generate ffjson $GOFILE

type Impact struct {
	UserID string
	Points int
	Result string // options: win|lost|draw
}

type PlayCompletedPayload struct {
	PlayID       string
	TournamentID string `json:",omitempty"`
	Impacts      []Impact
}

type PlayCompletedResponse struct {
	Tournament pamodels.Tournament
	Players []ContestantStats `json:",omitempty"`
}

type ContestantStats struct {
	GameUserID string
	Position   int
	PlaysCount int
	WinsCount  int `json:",omitempty"`
	DrawsCount int `json:",omitempty"`
}
