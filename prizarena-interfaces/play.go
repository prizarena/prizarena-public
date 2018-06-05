package prizarena_interfaces

type Impact struct {
	UserID string
	Points int
}

type PlayCompletedEvent struct {
	PlayID       string
	TournamentID string
	Impacts      []Impact
}