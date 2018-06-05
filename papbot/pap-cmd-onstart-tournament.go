package papbot

import "github.com/strongo/bots-framework/core"

const onStartTournamentCode = "PapOnStartTournament"

var OnStartTournament = bots.Command{
	Code: onStartTournamentCode,
	Action: func(whc bots.WebhookContext) (m bots.MessageFromBot, err error) {
		return
	},
}
