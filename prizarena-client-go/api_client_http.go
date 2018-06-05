package prizarena

import (
	"net/http"
	"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"strings"
	"io"
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
	)

const (
	contentTypeJson            = "application/json"
	ApiEndpointNewTournament   = "/api/new-tournament"
	ApiEndpointPlayCompleted   = "/api/play-completed"
	ApiEndpointUserTournaments = "/api/user/tournaments"
	ApiEndpointTournamentInfo  = "/api/tournament/info"
)

func NewHttpApiClient(httpClient *http.Client, server string, token string) prizarena_interfaces.ApiClient {
	if server == "" {
		server = "https://prizarena.com/api"
	}
	return httpApiClient{httpClient: httpClient, server: server, token: token}
}

type httpApiClient struct {
	httpClient *http.Client
	server     string
	token      string
}

func (apiClient httpApiClient) post(endpoint string, body io.Reader, response interface{}) (err error) {
	var resp *http.Response
	if resp, err = apiClient.httpClient.Post(apiClient.server + endpoint, contentTypeJson, body); err != nil {
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
			err = errors.WithMessage(err, "failed to decode response with OK status code as JSON")
			return
		}
	case http.StatusUnauthorized:
		err = ErrUnauthorized;
		return
	case http.StatusForbidden:
		err = ErrForbidden;
		return
	case http.StatusInternalServerError, http.StatusBadRequest:
		errResp := ErrorResponse{}
		if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			err = errors.WithMessage(err, fmt.Sprintf("failed to decode response with status code %v as JSON", resp.StatusCode))
			return
		}
		switch resp.StatusCode {
		case http.StatusBadRequest:
			err = ErrBadRequest{apiError: apiError{Code: errResp.Code, Message: errResp.Message}}
			return
		case http.StatusInternalServerError:
			err = ErrInternalServerError{apiError: apiError{Code: errResp.Code, Message: errResp.Message}}
			return
		default:
			err = apiError{Code: errResp.Code, Message: errResp.Message}
		}
	}
	return
}

func (apiClient httpApiClient) NewTournament(c context.Context, newTournament prizarena_interfaces.NewTournament) (tournament prizarena_interfaces.TournamentDto, err error) {
	body := strings.Reader{}
	err = apiClient.post(ApiEndpointNewTournament, &body, &tournament)
	return
}

func (apiClient httpApiClient) PlayCompleted(c context.Context, e prizarena_interfaces.PlayCompletedEvent) (err error) {
	body := strings.Reader{}
	err = apiClient.post(ApiEndpointPlayCompleted, &body, &PlayCompletedResponse{})
	return
}
