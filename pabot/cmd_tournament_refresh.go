package pabot

import (
	"github.com/strongo/bots-framework/core"
	"net/url"
	"github.com/prizarena/prizarena-public/pamodels"
)

const refreshTournamentCommandCode = "refresh-tournament"

func refreshTournamentCallbackData(tournamentID string) string {
	return refreshTournamentCommandCode + "?id=" + tournamentID
}

var refreshTournamentCommand = bots.NewCallbackCommand(
	refreshTournamentCommandCode,
	func(whc bots.WebhookContext, callbackUrl *url.URL) (m bots.MessageFromBot, err error) {
		var tournament pamodels.Tournament
		tournament.ID = callbackUrl.Query().Get("id")
		if err = pamodels.VerifyIsFullTournamentID(tournament.ID); err != nil {
			return
		}
		c := whc.Context()
		httpClient := whc.BotContext().BotHost.GetHTTPClient(c)
		prizarenaApiClient := newPrizarenaApiClient(httpClient)
		if tournament, err = prizarenaApiClient.GetTournament(c, tournament.ID); err != nil {
			return
		}
		m, err = RenderTournamentCard(whc.Context(), TournamentCardModeEditCallbackMessage, tournament)
		return
	},
)
