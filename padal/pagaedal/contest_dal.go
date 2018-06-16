package pagaedal

import (
	"github.com/prizarena/greed-game/server-go/greedgame/dal"
	"github.com/prizarena/greed-game/server-go/greedgame/models"
	"context"
	"github.com/pkg/errors"
	"github.com/prizarena/arena/arena-go"
	"github.com/strongo/log"
	"google.golang.org/appengine/datastore"
	"time"
	"github.com/prizarena/prizarena-public/padal"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
)

type tournamentGaeDal struct {
}

var _ padal.TournamentDal = (*tournamentGaeDal)(nil)

func (tournamentGaeDal) FindStranger(c context.Context, tournamentID, userID string, ignoreIDs []string) (strangerUserID string, err error) {
	log.Debugf(c, "tournamentGaeDal.FindStranger(tournamentID=%v, userID=%v, ignoreIDs=%v)", tournamentID, userID, ignoreIDs)
	if tournamentID == "" {
		err = errors.New("Parameter tournamentID is empty string")
		return
	} else if prizarena_interfaces.IsMonthlyTournamentID(tournamentID) {
		err = errors.New("Parameter tournamentID is monthly tournament ID")
		return
	}

	iterator := datastore.NewQuery(pamodels.ContestantKind).
		Filter("TournamentID =", tournamentID).
		Filter("StrangerCreated >", time.Time{}).
		Order("Stranger").
		KeysOnly().
		Run(c)

OUTER:
	for {
		var key *datastore.Key
		if key, err = iterator.Next(nil); err != nil {
			if err == datastore.Done {
				err = nil
			}
			break
		}

		if strangerUserID = pamodels.ContestantID(key.StringID()).UserID(); strangerUserID == userID {
			continue OUTER
		} else {
			for _, ignoreID := range ignoreIDs {
				if strangerUserID == ignoreID {
					continue OUTER
				}
			}
		}

		var strangerUser models.User
		if strangerUser, err = dal.User.GetUserByID(c, strangerUserID); err != nil {
			break
		}

		if rivalBid := strangerUser.GetBattles().GetBattleByRivalID(arena.NewStrangerBattleID(tournamentID)); rivalBid != nil {
			return // Stranger found
		}
	}
	strangerUserID = "" // Stranger not found
	return
}
