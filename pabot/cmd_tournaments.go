package pabot

import (
	"github.com/strongo/bots-framework/core"
	"bytes"
	"github.com/strongo/bots-api-telegram"
	"github.com/strongo/app"
	"net/url"
	"fmt"
	"strings"
	"github.com/prizarena/prizarena-public/patrans"
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
	s.WriteString(whc.Translate(patrans.TournamentsIntro))

	m.Format = bots.MessageFormatHTML
	m.Text = s.String()
	m.Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		NewTournamentTelegramInlineButton(whc, prizarenaGameID),
		ListTournamentsTelegramInlineButton(whc, prizarenaGameID),
	)
	return
}

const (
	PrizarenaBotStartCommandListTournaments = "list"
	PrizarenaBotStartCommandNewTournament = "new"
)

func prizarenaTelegramBotStartURL(prizarenaGameID, command, lang string) string {
	s := fmt.Sprintf("https://t.me/prizarena_bot?start=%v", prizarenaGameID)
	switch command {
	case "":
		// OK, nothing to add
	case "list", "new":
		s += "__" + command
	default:
		panic("Unknown prizarena bot start command")
	}
	if lang != "" && !strings.HasPrefix(lang,"en") {
		s += "_" + lang[:2]
	}
	return s
}

func NewTournamentTelegramInlineButton(t strongo.SingleLocaleTranslator, prizarenaGameID string) []tgbotapi.InlineKeyboardButton{
	return []tgbotapi.InlineKeyboardButton{{Text: t.Translate(patrans.NewTournamentButton), URL: prizarenaTelegramBotStartURL(prizarenaGameID, PrizarenaBotStartCommandNewTournament, t.Locale().Code5)}}
}

func ListTournamentsTelegramInlineButton(t strongo.SingleLocaleTranslator, prizarenaGameID string) []tgbotapi.InlineKeyboardButton{
	return []tgbotapi.InlineKeyboardButton{{Text: t.Translate(patrans.TournamentsButton), URL: prizarenaTelegramBotStartURL(prizarenaGameID, PrizarenaBotStartCommandListTournaments, t.Locale().Code5)}}
}
