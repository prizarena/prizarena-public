package padal

import (
	"context"
	"github.com/prizarena/prizarena-public/pamodels"
)

type TournamentDal interface {
	FindStranger(c context.Context, tournamentID, userID string, friends []string) (strangerID string, err error)
	GetUserTournaments(c context.Context, userID, orderBy string, limit int, keysOnly bool) (tournaments []pamodels.Tournament, err error)
}
