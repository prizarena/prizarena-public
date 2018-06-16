package padal

import (
	"fmt"
	"context"
	"github.com/prizarena/prizarena-public/pamodels"
)

func GetTournamentByID(c context.Context, tournamentID string) (tournament pamodels.Tournament, err error) {
	tournament.ID = tournamentID
	err = DB.Get(c, &tournament)
	return
}

func GetTournamentID(gameID, gameTournamentID string) string {
	return fmt.Sprintf("%v:%v", gameID, gameTournamentID)
}