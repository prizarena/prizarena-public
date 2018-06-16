package pabot

import (
	"net/http"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"github.com/strongo/bots-framework/core"
)

var newPrizarenaApiClient func(httpClient *http.Client) prizarena_interfaces.ApiClient

func InitPrizarenaBot(router bots.WebhooksRouter, prizarenaApiClientFactory func(httpClient *http.Client) prizarena_interfaces.ApiClient) {
	if prizarenaApiClientFactory == nil {
		panic("prizarenaApiClientFactory is required parameter")
	}
	newPrizarenaApiClient = prizarenaApiClientFactory
	router.RegisterCommands(
		[]bots.Command{
			refreshTournamentCommand,
		},
	)
}

func NewPrizarenaApiClient(httpClient *http.Client) prizarena_interfaces.ApiClient {
	if httpClient == nil {
		panic("httpClient == nil")
	}
	return newPrizarenaApiClient(httpClient)
}