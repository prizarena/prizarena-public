package pabot

import (
	"github.com/strongo/bots-framework/core"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/pacached"
	"net/url"
	"strings"
	"github.com/strongo/log"
	"github.com/strongo/db"
)

type InlineQueryContext struct {
	ID   string
	Text string
}

type InlineQueryMessageBuilder func(tournament pamodels.Tournament) (m bots.MessageFromBot, err error)

func ProcessInlineQueryTournament(whc bots.WebhookContext, inlineQuery InlineQueryContext, gameID, tournamentParamName string, reply InlineQueryMessageBuilder) (m bots.MessageFromBot, err error) {
	c := whc.Context()
	var tournament pamodels.Tournament
	if tournament.ID, err = GetQueryValueFromUrl(inlineQuery.Text, tournamentParamName); err != nil {
		return
	}
	if tournament.ID == "" {
		return reply(tournament)
	}
	tournament.ID = gameID + pamodels.TournamentIDSeparator + tournament.ID
	log.Debugf(c, "tournament.ID: %v", tournament.ID)
	cached := NewCachedApi(whc)
	if tournament, err = cached.GetTournamentByID(c, tournament.ID); err != nil {
		if db.IsNotFound(err) {
			log.Debugf(c, err.Error())
			err = nil
			return
		}
		return
	}
	return reply(tournament)
}

func GetQueryValueFromUrl(inlineQueryText, paramName string) (v string, err error) {
	if values, err := GetUrlValues(inlineQueryText); err != nil {
		return "", err
	} else if values != nil {
		v = values.Get(paramName)
	}
	return
}

func GetUrlValues(s string) (values url.Values, err error) {
	if qIndex := strings.Index(s, "?"); qIndex >= 0 {
		s = s[qIndex+1:]
	}
	return url.ParseQuery(s)
}

func NewCachedApi(whc bots.WebhookContext) pacached.Cached {
	c := whc.Context()
	httpClient := whc.BotContext().BotHost.GetHTTPClient(c)
	apiClient := NewPrizarenaApiClient(httpClient)
	return pacached.NewCached(apiClient)
}