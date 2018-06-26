package pabot

import (
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"github.com/strongo/bots-framework/core"
	"github.com/prizarena/prizarena-public/padal/pagaedal"
	"context"
	"google.golang.org/appengine/urlfetch"
	"github.com/prizarena/prizarena-public/prizarena-client-go"
)

func InitPrizarenaInGameBot(prizarenaGameID, prizarenaToken string, router bots.WebhooksRouter) {
	newPrizarenaApi := func(c context.Context) prizarena_interfaces.ApiClient {
		return newPrizarenaApiUrlfetchClient(c, "", prizarenaGameID, prizarenaToken)
	}
	router.RegisterCommands(
		[]bots.Command{
			refreshTournamentCommand(newPrizarenaApi),
			tournamentsCommand(prizarenaGameID),
			playStrangerCommand(newPrizarenaApi),
		},
	)
	pagaedal.RegisterDal()
}

var newPrizarenaApiUrlfetchClient = func(c context.Context, server, gameID, token string) prizarena_interfaces.ApiClient {
	httpClient := urlfetch.Client(c)
	return prizarena.NewHttpApiClient(httpClient, server, gameID, token)
}
