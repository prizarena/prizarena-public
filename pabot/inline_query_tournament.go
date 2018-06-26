package pabot

import (
	"github.com/strongo/bots-framework/core"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/prizarena/prizarena-public/pacached"
	"net/url"
	"strings"
	"github.com/strongo/log"
	"github.com/strongo/db"
	"context"
)

type InlineQueryContext struct {
	ID   string
	Text string
}

type InlineQueryMessageBuilder func(tournament pamodels.Tournament) (m bots.MessageFromBot, err error)

func ProcessInlineQueryTournament(whc bots.WebhookContext, inlineQuery InlineQueryContext, prizarenaGameID, prizarenaToken, tournamentParamName string, reply InlineQueryMessageBuilder) (m bots.MessageFromBot, err error) {
	c := whc.Context()
	var tournament pamodels.Tournament
	if tournament.ID, err = GetQueryValueFromUrl(inlineQuery.Text, tournamentParamName); err != nil {
		return
	}
	if tournament.ID == "" {
		return reply(tournament)
	}
	tournament.ID = prizarenaGameID + pamodels.TournamentIDSeparator + tournament.ID
	log.Debugf(c, "tournament.ID: %v", tournament.ID)
	cached := NewCachedApi(c, prizarenaGameID, prizarenaToken)
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

func NewCachedApi(c context.Context, prizarenaGameID, prizarenaToken string) pacached.Cached {
	apiClient := newPrizarenaApiUrlfetchClient(c, "", prizarenaGameID, prizarenaToken)
	return pacached.NewCached(apiClient)
}