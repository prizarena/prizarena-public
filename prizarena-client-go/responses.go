package prizarena

import (
		"github.com/prizarena/prizarena-public/pamodels"
)

//go:generate ffjson $GOFILE

type ErrorResponse struct {
	Code    string
	Message string
}

type TournamentCompletedResponse struct {
	Tournament pamodels.Tournament
}

