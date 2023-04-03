package pabot

import (
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/strongo/bots-framework/botsfw"
	"strings"
)

const onStartTournamentCode = "PapOnStartTournament"

var OnStartTournament = botsfw.Command{
	Code: onStartTournamentCode,
	Action: func(whc botsfw.WebhookContext) (m botsfw.MessageFromBot, err error) {
		return
	},
}

func OnStartIfTournamentLink(whc botsfw.WebhookContext, prizarenaGameID, prizarenaToken string) (m botsfw.MessageFromBot, err error) {
	if whc.InputType() != botsfw.WebhookInputText {
		return
	}
	input := whc.Input().(botsfw.WebhookTextMessage)
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
