package pagaedal

import (
	"github.com/strongo/db/gaedb"
	"github.com/prizarena/prizarena-public/padal"
	"github.com/strongo/app/gaestandard"
)

func RegisterDal() {
	padal.DB = gaedb.NewDatabase()
	padal.HandleWithContext = gaestandard.HandleWithContext
	padal.Tournament = tournamentGaeDal{}
}
