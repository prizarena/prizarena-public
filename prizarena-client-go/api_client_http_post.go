package prizarena

import (
	"net/http"
	"encoding/json"
	"github.com/pkg/errors"
	"fmt"
	"bytes"
	"encoding/base64"
	"io"
	"net/url"
)

func (apiClient httpApiClient) get(endpoint string, params url.Values, response interface{}) (err error) {
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}
	return apiClient.rpc("GET", endpoint, nil, response)
}

func (apiClient httpApiClient) post(endpoint string, reqData interface{}, response interface{}) (err error) {
	var reqBody bytes.Buffer
	if err = json.NewEncoder(&reqBody).Encode(reqData); err != nil {
		return errors.WithMessage(err, "failed to encode request data")
	}
	return apiClient.rpc("POST", endpoint, &reqBody, response)
}

func (apiClient httpApiClient) rpc(method, endpoint string, requestBody io.Reader, response interface{}) (err error) {
	var (
		req  *http.Request
		resp *http.Response
	)
	req, err = http.NewRequest(method, apiClient.server+endpoint, requestBody)
	req.Header.Add("Content-Type", contentTypeJson)
	{
		var buf bytes.Buffer
		encoder := base64.NewEncoder(base64.URLEncoding, &buf)
		if _, err = encoder.Write([]byte(fmt.Sprintf("%v:%v", apiClient.gameID, apiClient.token))); err != nil {
			return errors.WithMessage(err, "failed to encode Authorization header to base54")
		}
		encoder.Close()
		req.Header.Add("Authorization", "Basic "+buf.String())
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
