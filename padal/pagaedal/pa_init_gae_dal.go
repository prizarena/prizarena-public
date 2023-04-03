package pagaedal

import (
	"github.com/prizarena/prizarena-public/padal"
	"github.com/strongo/app/gaestandard"
)

func RegisterDal() {
	//padal.DB = gae.NewDatabase()
	padal.HandleWithContext = gaestandard.HandleWithContext
	padal.Tournament = tournamentGaeDal{}
}
