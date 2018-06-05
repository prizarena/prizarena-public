package prizarena

import (
	"net/http"
	"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"strings"
	"encoding/json"
)

const (
	contentTypeJson = "application/json"
	ApiEndpointNewTournament = "/api/new-tournament"
	ApiEndpointPlayCompleted = "/api/play-completed"
)

func NewHttpApiClient(httpClient *http.Client, token string) prizarena_interfaces.ApiClient {
	return httpApiClient{httpClient: httpClient, token: token}
}

type httpApiClient struct {
	httpClient *http.Client
	token string
}

func (apiClient httpApiClient) NewTournament(c context.Context, newTournament prizarena_interfaces.NewTournament) (tournament prizarena_interfaces.Tournament, err error) {
	body := strings.Reader{}
	var resp *http.Response
	resp, err = apiClient.httpClient.Post(ApiEndpointNewTournament, contentTypeJson, &body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(resp.Body).Decode(&tournament)
	case http.StatusInternalServerError, http.StatusBadRequest:
		errResp := ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errResp)
	}
	return
}

func (apiClient httpApiClient) PlayCompleted(c context.Context, e prizarena_interfaces.PlayCompletedEvent) (err error){
	body := strings.Reader{}
	_, err = apiClient.httpClient.Post(ApiEndpointPlayCompleted, contentTypeJson, &body)
	return
}