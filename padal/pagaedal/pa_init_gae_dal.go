package pagaedal

import (
	"github.com/strongo/db/gaedb"
	"github.com/strongo/app/gaestandard"
	"github.com/prizarena/prizarena-public/padal"
)

func RegisterDal() {
	padal.DB = gaedb.NewDatabase()
	padal.HandleWithContext = gaestandard.HandleWithContext
	padal.Tournament = tournamentGaeDal{}
}
