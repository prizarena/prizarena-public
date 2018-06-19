package pacached

import (
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"github.com/prizarena/prizarena-public/pamodels"
	"context"
	"time"
	"github.com/prizarena/prizarena-public/padal"
	"github.com/strongo/db"
	"github.com/strongo/log"
)

type cached struct {
	apiClient prizarena_interfaces.ApiClient
}

type Cached interface {
	GetTournamentByID(c context.Context, id string) (tournament pamodels.Tournament, err error)
}

func NewCached(apiClient prizarena_interfaces.ApiClient) Cached {
	return cached{apiClient: apiClient}
}

func (wrapper cached) GetTournamentByID(c context.Context, id string) (tournament pamodels.Tournament, err error) {
	if err = pamodels.VerifyIsFullTournamentID(id); err != nil {
		return
	}
	tournament.ID = id
	if padal.DB == nil {
		log.Warningf(c, "cached.GetTournamentByID() => padal.DB == nil")
	} else {
		err = padal.DB.Get(c, &tournament)
		if !db.IsNotFound(err) {
			if err != nil {
				log.Warningf(c, "Failed to get tournament from local DB: %v", err)
			} else if tournament.Cached.After(time.Now().Add(-time.Minute)) {
				return
			}
		}
	}
	tournament, err = wrapper.apiClient.GetTournament(c, id)
	if padal.DB != nil {
		tournament.Cached = time.Now()
		go func() {
			padal.DB.Update(c, &tournament)
		}()
	}
	return
}
