package prizarena_interfaces

//go:generate ffjson $GOFILE

type Impact struct {
	UserID string
	Points int
}

type PlayCompletedEvent struct {
	PlayID       string
	TournamentID string `json:",omitempty"`
	Impacts      []Impact
}

type PlayCompletedResponse struct {
	Tournament PlayCompletedTournament
}

type PlayCompletedTournament struct {
	TournamentDto
	TournamentUserStats
}

type TournamentUserStats struct {
	ContestantsCount   int
	PlaysCount         int
	WinsCount          int                    `json:",omitempty"`
	DrawsCount         int                    `json:",omitempty"`
	Position           int                    `json:",omitempty"`
	ClosestContestants []TournamentContestant `json:",omitempty"`
}

type TournamentContestant struct {
	Position int
	Name     string
}
