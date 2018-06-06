package prizarena_interfaces

type PairWithStrangerPayload struct {
	TournamentID string
	GameUserID   string
	Move         *MoveDto `json:",omitempty"`
}

type PairWithStrangerResponse struct {
	RivalGameUserID string
	RivalMove       *MoveDto `json:",omitempty"`
}

type PairedWithStrangerPayload struct {
	GameUserID      string
	RivalGameUserID string
}

type PairedWithStrangerResponse struct {
}
