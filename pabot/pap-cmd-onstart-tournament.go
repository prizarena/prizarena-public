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

func OnStartIfTournamentLink(whc bots.WebhookContext, prizarenaGameID, prizarenaToken string) (m bots.MessageFromBot, err error) {
	if whc.InputType() != bots.WebhookInputText {
		return
	}
	input := whc.Input().(bots.WebhookTextMessage)
	text := input.Text()
	if strings.HasPrefix(text, "/start ") {
		text = text[7:]
	}
	if !strings.HasPrefix(text, "t-") {
		return
	}
	tournamentGameID := strings.Split(text, "__")[0][2:]
	c := whc.Context()

	var tournament pamodels.Tournament

	tournamentFullID := prizarenaGameID + ":" + tournamentGameID

	if tournament, err = newPrizarenaApiUrlfetchClient(c, "", prizarenaGameID, prizarenaToken).GetTournament(c, tournamentFullID); err != nil {
		return
	}

	if m, err = RenderTournamentCard(whc.Context(), TournamentCardModeNewMessage, tournament); err != nil {
		return
	}
	return
}
