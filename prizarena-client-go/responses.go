package prizarena

import "github.com/prizarena/prizarena-public/prizarena-interfaces"

//go:generate ffjson $GOFILE

type ErrorResponse struct {
	Code    string
	Message string
}

type TournamentCompletedResponse struct {
	Tournament prizarena_interfaces.TournamentDto
}

type PlayCompletedResponse struct {
	Tournament TournamentUserStats
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
