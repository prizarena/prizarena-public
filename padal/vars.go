package padal

import (
	"github.com/strongo/db"
		"github.com/strongo/app"
)

var (
	DB db.Database
	Tournament TournamentDal
	HandleWithContext strongo.HandleWithContext
)
