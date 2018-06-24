package pabot

import (
	"github.com/strongo/bots-framework/core"
	"bytes"
	"github.com/strongo/bots-api-telegram"
	"github.com/strongo/app"
	"net/url"
)

var TournamentsCommandCode = "tournaments"

func tournamentsCommand(prizarenaGameID string) bots.Command{
	return bots.Command{
		Code: TournamentsCommandCode,
		Commands: []string{"/tournaments"},
		Action: func(whc bots.WebhookContext) (m bots.MessageFromBot, err error) {
			return tournamentsAction(whc, prizarenaGameID)
		},
		CallbackAction: func(whc bots.WebhookContext, callbackUrl *url.URL) (m bots.MessageFromBot, err error) {
			return tournamentsAction(whc, prizarenaGameID)
		},
	}
}

func tournamentsAction(whc bots.WebhookContext, prizarenaGameID string) (m bots.MessageFromBot, err error) {
	s := new(bytes.Buffer)
	s.WriteString("<b>Tournaments</b>")

	m.Format = bots.MessageFormatHTML
	m.Text = s.String()
	m.Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		NewTournamentTelegramInlineButton(whc, prizarenaGameID),
	)
	return
}

func NewTournamentTelegramInlineButton(t strongo.SingleLocaleTranslator, prizarenaGameID string) []tgbotapi.InlineKeyboardButton{
	return []tgbotapi.InlineKeyboardButton{{Text: "⚔ New tournament", URL: "https://t.me/prizarena_bot?start="+prizarenaGameID}}
}
