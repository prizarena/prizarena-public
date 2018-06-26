package prizarena

import (
	"net/http"
	"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"strings"
						"github.com/prizarena/prizarena-public/pamodels"
	"net/url"
)

const (
	contentTypeJson               = "application/json"
	ApiEndpointNewTournament      = "/api/new-tournament"
	ApiEndpointPlayCompleted      = "/api/play-completed"
	ApiEndpointPairWithStranger   = "/api/stranger/pair"
	ApiEndpointPairedWithStranger = "/api/stranger/paired"
	ApiEndpointUserTournaments    = "/api/user/tournaments"
	ApiEndpointLeaveTournament    = "/api/leave/tournament"
	ApiEndpointTournamentInfo     = "/api/tournament/info"
)

func NewHttpApiClient(httpClient *http.Client, server, gameID, token string) httpApiClient {
	if httpClient == nil {
		panic("httpClient == nil")
	}
	if server == "" {
		server = "https://prizarena.com"
	}
	return httpApiClient{httpClient: httpClient, server: server, gameID: gameID, token: token}
}

type httpApiClient struct {
	httpClient *http.Client
	server     string
	gameID     string
	token      string
}

var _ prizarena_interfaces.ApiClient = (*httpApiClient)(nil)

func (apiClient httpApiClient) GetTournament(c context.Context, tournamentID string) (tournament pamodels.Tournament, err error) {
	if err = pamodels.VerifyIsFullTournamentID(tournamentID); err != nil {
		return
	}
	err = apiClient.get(ApiEndpointTournamentInfo, url.Values{"id": []string{tournamentID}}, &tournament)
	return
}

func (apiClient httpApiClient) NewTournament(c context.Context, newTournament prizarena_interfaces.NewTournamentPayload) (response prizarena_interfaces.NewTournamentResponse, err error) {
	body := strings.Reader{}
	err = apiClient.post(ApiEndpointNewTournament, &body, &response)
	return
}

func (apiClient httpApiClient) PlayCompleted(c context.Context, e prizarena_interfaces.PlayCompletedPayload) (response prizarena_interfaces.PlayCompletedResponse, err error) {
	err = apiClient.post(ApiEndpointPlayCompleted, &e, &response)
	return
}

func (apiClient httpApiClient) PairWithStranger(c context.Context, payload prizarena_interfaces.PairWithStrangerRequest) (response prizarena_interfaces.PairWithStrangerResponse, err error) {
	err = apiClient.post(ApiEndpointPairWithStranger, &payload, &response)
	return
}

func (apiClient httpApiClient) PairedWithStranger(c context.Context, payload prizarena_interfaces.PairedWithStrangerPayload) (response prizarena_interfaces.PairedWithStrangerResponse, err error) {
	err = apiClient.post(ApiEndpointPairedWithStranger, &payload, &response)
	return
}

func (apiClient httpApiClient) LeaveTournament(c context.Context, battleID string) error {
	return apiClient.post(ApiEndpointLeaveTournament, battleID, nil)
}
