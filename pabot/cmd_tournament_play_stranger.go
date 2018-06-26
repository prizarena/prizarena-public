package pabot

import (
	"github.com/strongo/bots-framework/core"
	"net/url"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	)

const playStrangerCommandCode = "play-stranger"

func getPlayStrangerCallbackData(tournamentID string) string {
	return playStrangerCommandCode + "?tournament=" + tournamentID
}

func playStrangerCommand(prizarenaApiClientFactory prizarena_interfaces.ApiClientFactory) bots.Command {
	return bots.NewCallbackCommand(playStrangerCommandCode,
		func(whc bots.WebhookContext, callbackUrl *url.URL) (m bots.MessageFromBot, err error) {
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
