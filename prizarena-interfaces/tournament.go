package prizarena_interfaces

import (
	"strings"
	"strconv"
)

const MonthlyTournamentIDFormat = "200601"

type TournamentDto struct {
	ID string
	Name string
	GameID string
	Sponsor Sponsor
}

type NewTournamentResponse struct {
	Tournament TournamentDto
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
