package prizarena_interfaces

//go:generate ffjson $GOFILE

type Impact struct {
	UserID string
	Points int
}

type PlayCompletedEvent struct {
	PlayID       string
	TournamentID string `json:",omitempty"`
	Impacts      []Impact
}
