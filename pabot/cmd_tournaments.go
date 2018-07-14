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
	"github.com/strongo/emoji/go/emoji"
)

var TournamentsCommandCode = "tournaments"

func tournamentsCommand(prizarenaGameID string) bots.Command {
	return bots.Command{
		Code:     TournamentsCommandCode,
		Commands: []string{"/" + TournamentsCommandCode},
		Action: func(whc bots.WebhookContext) (m bots.MessageFromBot, err error) {
			return tournamentsAction(whc, prizarenaGameID)
		},
		CallbackAction: func(whc bots.WebhookContext, callbackUrl *url.URL) (m bots.MessageFromBot, err error) {
			return tournamentsAction(whc, prizarenaGameID)
		},
	}
}

var languageOptions = []struct {
	flag string
	lang string
}{
	{emoji.UnitedKingdom, "en-US"},
	{emoji.Germany, "de-DE"},
	{emoji.Spain, "es-ES"},
	{emoji.Italy, "it-IT"},
	{emoji.France, "fr-FR"},
	{emoji.Russia, "ru-RU"},
	//{emoji.Uzbekistan, "uz-UZ"},
	{emoji.Iran, "fa-IR"},
	//{emoji.Brazil, "pt-BR"},
}

func GetLangButtons(command, currentLang string) (row []tgbotapi.InlineKeyboardButton) {
	row = make([]tgbotapi.InlineKeyboardButton, 0, 8)
	for _, lo := range languageOptions {
		if lo.lang == currentLang {
			continue
		}
		row = append(row, tgbotapi.InlineKeyboardButton{Text: lo.flag, CallbackData: command + "?l=" + lo.lang})
		//if len(languageButtons) == 7 {
		//	languageButtons = append(languageButtons, tgbotapi.InlineKeyboardButton{Text: "...", CallbackData: TournamentsCommandCode + "?l=more"})
		//}
	}
	return
}

func tournamentsAction(whc bots.WebhookContext, prizarenaGameID string) (m bots.MessageFromBot, err error) {
	s := new(bytes.Buffer)
	s.WriteString(whc.Translate(patrans.TournamentsIntro))

	m.Format = bots.MessageFormatHTML
	m.Text = s.String()
	m.IsEdit = true
	m.Keyboard = tgbotapi.NewInlineKeyboardMarkup(
		GetLangButtons(TournamentsCommandCode, whc.Locale().Code5),
		[]tgbotapi.InlineKeyboardButton{
			{Text: whc.Translate(patrans.MainMenuButton), CallbackData: "start"},
		},
		NewTournamentTelegramInlineButton(whc, prizarenaGameID),
		ListTournamentsTelegramInlineButton(whc, prizarenaGameID),
	)
	return
}

const (
	PrizarenaBotStartCommandListTournaments = "list"
	PrizarenaBotStartCommandNewTournament   = "new"
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
	if lang != "" && !strings.HasPrefix(lang, "en") {
		s += "_" + lang[:2]
	}
	return s
}

func NewTournamentTelegramInlineButton(t strongo.SingleLocaleTranslator, prizarenaGameID string) []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{{Text: t.Translate(patrans.NewTournamentButton), URL: prizarenaTelegramBotStartURL(prizarenaGameID, PrizarenaBotStartCommandNewTournament, t.Locale().Code5)}}
}

func ListTournamentsTelegramInlineButton(t strongo.SingleLocaleTranslator, prizarenaGameID string) []tgbotapi.InlineKeyboardButton {
	return []tgbotapi.InlineKeyboardButton{{Text: t.Translate(patrans.TournamentsButton), URL: prizarenaTelegramBotStartURL(prizarenaGameID, PrizarenaBotStartCommandListTournaments, t.Locale().Code5)}}
}
