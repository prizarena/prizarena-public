package prizarena

import (
	"testing"
	"net/http"
	"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
	"net/http/httptest"
	"encoding/json"
	"fmt"
	"encoding/base64"
	"strings"
	"bytes"
)

func TestNewHttpApiClientWithNilHttpClient(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("should raise panic if nil passed as httpClient")
		}
	}()
	NewHttpApiClient(nil, "", "test-game", "123")
}

func TestNewHttpApiClient(t *testing.T) {
	var apiClient prizarena_interfaces.ApiClient = NewHttpApiClient(http.DefaultClient, "", "test-game", "123")
	switch apiClient.(type) {
	case httpApiClient: // OK
	default:
		t.Fatalf("Expected instance of httpApiClient got %T", apiClient)
	}
}

const secureToken = "t123" //t0ken

func TestHttpApiClient_NewTournament(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		gameID := getFromAuthHeaderGameID(t, r, secureToken)
		w.Write([]byte(fmt.Sprintf(`{"Tournament": {"ID": "a1b2", "GameID": "%v"}}`, gameID)))
	}))
	defer ts.Close()

	const gameID = "test-game"
	client := NewHttpApiClient(http.DefaultClient, ts.URL, gameID, secureToken)
	c := context.Background()
	response, err := client.NewTournament(c, prizarena_interfaces.NewTournamentPayload{})
	if err != nil {
		t.Fatal(err)
	}
	if response.Tournament.ID != "a1b2" {
		t.Fatalf("Unexpected response.Tournament.ID: [%v]", response.Tournament.ID)
	}
	if response.Tournament.GameID != gameID {
		t.Fatalf("Unexpected response.Tournament.GameID: [%v]", response.Tournament.GameID)
	}
}

func getFromAuthHeaderGameID(t *testing.T, r *http.Request, expectedToken string) (gameID string) {
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Basic ") {
		return
	}
	authHeader = authHeader[6:]
	buf := bytes.NewBuffer([]byte{})
	buf.ReadFrom(base64.NewDecoder(base64.URLEncoding, strings.NewReader(authHeader)))
	authHeader = buf.String()
	v := strings.Split(authHeader, ":")
	if v[1] != expectedToken {
		t.Fatalf("token != expectedToken: %v != %v", v[1], expectedToken)
	}
	return v[0]
}

func TestHttpApiClient_PlayCompleted(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		gameID := getFromAuthHeaderGameID(t, r, secureToken)
		if gameID != "test-game" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		e := prizarena_interfaces.PlayCompletedPayload{}
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if e.TournamentID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(fmt.Sprintf(`{"Tournament": {"ID": "%v", "GameID": "%v", "ContestantsCount": 2}}`, e.TournamentID, gameID)))
	}))
	defer ts.Close()

	client := NewHttpApiClient(http.DefaultClient, ts.URL, "test-game", secureToken)
	c := context.Background()
	var (
		response prizarena_interfaces.PlayCompletedResponse
		err      error
	)
	response, err = client.PlayCompleted(c, prizarena_interfaces.PlayCompletedPayload{TournamentID: "a1b2", PlayID: "play123"})
	if err != nil {
		t.Fatal(err)
	}
	if response.Tournament.ID != "a1b2" {
		t.Fatalf("Unexpected response.Tournament.ID: [%v]", response.Tournament.ID)
	}
	if response.Tournament.ContestantsCount != 2 {
		t.Errorf("Unexpected response.ContestantsCount: %v", response.Tournament.ContestantsCount)
	}
	if response.Tournament.GameID != "test-game" {
		t.Errorf("Unexpected response.Tournament.GameID: [%v]", response.Tournament.GameID)
	}
}
