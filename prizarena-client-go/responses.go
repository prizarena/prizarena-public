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

