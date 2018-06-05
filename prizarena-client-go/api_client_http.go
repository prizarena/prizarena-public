package prizarena

import (
	"net/http"
	"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"strings"
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
	"bytes"
	"encoding/base64"
)

const (
	contentTypeJson            = "application/json"
	ApiEndpointNewTournament   = "/api/new-tournament"
	ApiEndpointPlayCompleted   = "/api/play-completed"
	ApiEndpointUserTournaments = "/api/user/tournaments"
	ApiEndpointTournamentInfo  = "/api/tournament/info"
)

func NewHttpApiClient(httpClient *http.Client, server, gameID, token string) httpApiClient {
	if httpClient == nil {
		panic("httpClient == nil")
	}
	if server == "" {
		server = "https://prizarena.com/api"
	}
	return httpApiClient{httpClient: httpClient, server: server, gameID: gameID, token: token}
}

type httpApiClient struct {
	httpClient *http.Client
	server     string
	gameID     string
	token      string
}

func (apiClient httpApiClient) post(endpoint string, reqData interface{}, response interface{}) (err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	var reqBody bytes.Buffer
	if err = json.NewEncoder(&reqBody).Encode(reqData); err != nil {
		return errors.WithMessage(err, "failed to encode request data")
	}
	req, err = http.NewRequest("POST", apiClient.server+endpoint, &reqBody)
	req.Header.Add("Content-Type", contentTypeJson)
	{
		var buf bytes.Buffer
		encoder := base64.NewEncoder(base64.URLEncoding, &buf)
		if _, err = encoder.Write([]byte(fmt.Sprintf("%v:%v", apiClient.gameID, apiClient.token))); err != nil {
			return errors.WithMessage(err, "failed to encode Authorization header to base54")
		}
		encoder.Close()
		req.Header.Add("Authorization", "Basic " + buf.String())
	}

	if resp, err = apiClient.httpClient.Do(req); err != nil {
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
		if resp.ContentLength > 0 {
			if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
				err = errors.WithMessage(err, fmt.Sprintf("failed to decode response with status code %v as JSON", resp.StatusCode))
				return
			}
		} else {
			errResp.Code = fmt.Sprintf("HTTP_STATUS=%v", resp.StatusCode)
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
	default:
		err := apiError{Code: fmt.Sprintf("HTTP_STATUS=%v", resp.StatusCode)}
		if resp.ContentLength > 0 {
			var buf bytes.Buffer
			buf.ReadFrom(resp.Body)
			err.Message = buf.String()
		}
		return err
	}
	return
}

func (apiClient httpApiClient) NewTournament(c context.Context, newTournament prizarena_interfaces.NewTournament) (response prizarena_interfaces.NewTournamentResponseDto, err error) {
	body := strings.Reader{}
	err = apiClient.post(ApiEndpointNewTournament, &body, &response)
	return
}

func (apiClient httpApiClient) PlayCompleted(c context.Context, e prizarena_interfaces.PlayCompletedEvent) (response prizarena_interfaces.PlayCompletedResponse, err error) {
	err = apiClient.post(ApiEndpointPlayCompleted, &e, &response)
	return
}
