package pabot

import (
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"github.com/strongo/bots-framework/botsfw"
	"net/url"
)

const playStrangerCommandCode = "play-stranger"

func getPlayStrangerCallbackData(tournamentID string) string {
	return playStrangerCommandCode + "?tournament=" + tournamentID
}

func playStrangerCommand(prizarenaApiClientFactory prizarena_interfaces.ApiClientFactory) botsfw.Command {
	return botsfw.NewCallbackCommand(playStrangerCommandCode,
		func(whc botsfw.WebhookContext, callbackUrl *url.URL) (m botsfw.MessageFromBot, err error) {
			c := whc.Context()
			var tournament pamodels.Tournament
			// // tournament.ID = prizarenaGameID + pamodels.TournamentIDSeparator + callbackUrl.Query().Get("tournament")
			tournament.ID = callbackUrl.Query().Get("tournament")
			// if tournament, err = padal.GetTournamentByID(c, tournament.ID); err != nil {
			// 	return
			// }

			payload := prizarena_interfaces.PairWithStrangerRequest{
				TournamentID: tournament.ID,
				GameUserID:   whc.AppUserStrID(),
			}

			prizarenaAPI := prizarenaApiClientFactory(c)
			var response prizarena_interfaces.PairWithStrangerResponse
			if response, err = prizarenaAPI.PairWithStranger(c, payload); err != nil {
				return
			}

			if response.RivalGameUserID == "" { // current user is stranger

			}
			return
		},
	)

}
