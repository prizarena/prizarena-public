package prizarena_interfaces

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/prizarena/prizarena-public/pamodels"
)

const MonthlyTournamentIDFormat = "200601"

type TournamentDto struct {
	ID string
	Name string
	GameID string
	Sponsor Sponsor
}

type NewTournamentResponse struct {
	Tournament pamodels.Tournament
}

// IsMonthlyTournamentID returns true if
func IsMonthlyTournamentID(tournamentID string) bool {
	if len(tournamentID) <= 7 || strings.Count(tournamentID, ":") != 1 {
		return false
	}
	if tournamentID = strings.Split(tournamentID, ":")[1]; len(tournamentID) == 6 {
		if v, err := strconv.ParseInt(tournamentID, 10, 32); err == nil && v > 201801 {
			return true
		}
	}
	return false
}

func GetUrlForOpeningTournamentInGameTelegramBot(bot, tournamentID, referral string) string {
	var gameTournamentID string
	if i := strings.Index(tournamentID, ":"); i >= 0 {
		gameTournamentID = tournamentID[i+1:]
	} else {
		gameTournamentID = tournamentID
	}
	s := fmt.Sprintf("https://t.me/%v?start=t-%v", bot, gameTournamentID)
	if referral != "" {
		s += "__ref-" + referral
	}
	return s
}