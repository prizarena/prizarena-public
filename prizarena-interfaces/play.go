package prizarena_interfaces

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
	Tournament PlayCompletedTournament
}

type PlayCompletedTournament struct {
	TournamentDto
	TournamentStats
	Players []ContestantStats `json:",omitempty"`
}

type TournamentStats struct {
	ContestantsCount int
	PlaysCount       int
}

type ContestantStats struct {
	GameUserID string
	Position   int
	PlaysCount int
	WinsCount  int `json:",omitempty"`
	DrawsCount int `json:",omitempty"`
}
