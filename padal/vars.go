package padal

import (
	"github.com/strongo/app"
	"github.com/strongo/dalgo/dal"
)

var (
	DB                dal.Database
	Tournament        TournamentDal
	HandleWithContext strongo.HandleWithContext
)
