package pagaedal

import (
	"github.com/prizarena/prizarena-public/padal"
	"github.com/strongo/app/gaestandard"
	"github.com/strongo/db/gaedb"
)

func RegisterDal() {
	padal.DB = gaedb.NewDatabase()
	padal.HandleWithContext = gaestandard.HandleWithContext
	padal.Tournament = tournamentGaeDal{}
}
