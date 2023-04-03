package pabot

import (
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"github.com/strongo/bots-framework/botsfw"
	"net/url"
)

const refreshTournamentCommandCode = "refresh-tournament"

func refreshTournamentCallbackData(tournamentID string) string {
	return refreshTournamentCommandCode + "?id=" + tournamentID
}

var refreshTournamentCommand = func(prizarenaApiFactory prizarena_interfaces.ApiClientFactory) botsfw.Command {
	return botsfw.NewCallbackCommand(
		refreshTournamentCommandCode,
		func(whc botsfw.WebhookContext, callbackUrl *url.URL) (m botsfw.MessageFromBot, err error) {
			var tournament pamodels.Tournament
			tournament.ID = callbackUrl.Query().Get("id")
			if err = pamodels.VerifyIsFullTournamentID(tournament.ID); err != nil {
				return
			}
			c := whc.Context()
			prizarenaAPI := prizarenaApiFactory(c)
			if tournament, err = prizarenaAPI.GetTournament(c, tournament.ID); err != nil {
				return
			}
			m, err = RenderTournamentCard(whc.Context(), TournamentCardModeEditCallbackMessage, tournament)
			return
		},
	)
}
