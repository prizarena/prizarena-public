package prizarena_interfaces

const MonthlyTournamentIDFormat = "200601"

type TournamentDto struct {
	ID string
	Name string
	GameID string
	Sponsor Sponsor
}
