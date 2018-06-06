package prizarena_interfaces

type PairWithStrangerPayload struct {
	TournamentID string
	GameUserID string
}

type PairWithStrangerResponse struct {
	RivalGameUserID string
}

type PairedWithStrangerPayload struct {
	GameUserID string
	RivalGameUserID string
}

type PairedWithStrangerResponse struct {
}