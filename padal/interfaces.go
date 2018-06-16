package padal

import "context"

type TournamentDal interface {
	FindStranger(c context.Context, tournamentID, userID string, friends []string) (strangerID string, err error)
}
