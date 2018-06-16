package prizarena_interfaces

import (
	"context"
	"github.com/prizarena/prizarena-public/pamodels"
)

type ApiClient interface {
	GetTournament(c context.Context, tournamentID string) (tournament pamodels.Tournament, err error)
	LeaveTournament(c context.Context, battleID string) error
	NewTournament(c context.Context, newTournament NewTournamentPayload) (response NewTournamentResponse, err error)
	PlayCompleted(c context.Context, payload PlayCompletedPayload) (response PlayCompletedResponse, err error)

	PairWithStranger(c context.Context, payload PairWithStrangerPayload) (response PairWithStrangerResponse, err error)
	PairedWithStranger(c context.Context, payload PairedWithStrangerPayload) (response PairedWithStrangerResponse, err error)
}

