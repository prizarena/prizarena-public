package pabot

import (
	"github.com/strongo/bots-framework/core"
	"strings"
			"github.com/prizarena/prizarena-public/pamodels"
)

const onStartTournamentCode = "PapOnStartTournament"

var OnStartTournament = bots.Command{
	Code: onStartTournamentCode,
	Action: func(whc bots.WebhookContext) (m bots.MessageFromBot, err error) {
		return
	},
}

func OnStartIfTournamentLink(whc bots.WebhookContext) (m bots.MessageFromBot, err error) {
	input := whc.Input().(bots.WebhookTextMessage)
	text := input.Text()
	if strings.HasPrefix(text, "/start ") {
		text = text[7:]
	}
	if !strings.HasPrefix(text, "t-") {
		return
	}
	tournamentID := strings.Split(text, "__")[0][2:]
	c := whc.Context()

	var tournament pamodels.Tournament

	httpClient := whc.BotContext().BotHost.GetHTTPClient(c)
	if tournament, err = newPrizarenaApiClient(httpClient).GetTournament(c, tournamentID); err != nil {
		return
	}

	if m, err = RenderTournamentCard(whc, TournamentCardModeNewMessage, tournament); err != nil {
		return
	}
	return
}
