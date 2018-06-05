package prizarena_interfaces

import (
	"context"
)

type ApiClient interface {
	NewTournament(c context.Context, newTournament NewTournament) (response NewTournamentResponseDto, err error)
	PlayCompleted(c context.Context, e PlayCompletedEvent) (response PlayCompletedResponse, err error)
}

