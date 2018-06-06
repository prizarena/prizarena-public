package prizarena_interfaces

import (
	"context"
)

type ApiClient interface {
	NewTournament(c context.Context, newTournament NewTournamentPayload) (response NewTournamentResponse, err error)
	PlayCompleted(c context.Context, payload PlayCompletedPayload) (response PlayCompletedResponse, err error)

	PairWithStranger(c context.Context, payload PairWithStrangerPayload) (response PairWithStrangerResponse, err error)
	PairedWithStranger(c context.Context, payload PairedWithStrangerPayload) (response PairedWithStrangerResponse, err error)
}

