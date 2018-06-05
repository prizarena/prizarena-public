package prizarena_interfaces

import (
	"context"
)

type ApiClient interface {
	NewTournament(c context.Context, newTournament NewTournament) (tournament TournamentDto, err error)
	PlayCompleted(c context.Context, e PlayCompletedEvent) (err error)
}

