package pabot

import (
	"github.com/prizarena/prizarena-public/pamodels"
	"fmt"
	"github.com/strongo/bots-framework/core"
	"bytes"
	"time"
	"bitbucket.org/asterus/prizarena-private/prizarena-server/papfacade"
	"github.com/strongo/bots-api-telegram"
	"context"
)

type TournamentCardMode int

const (
	TournamentCardModeNewMessage TournamentCardMode = iota
	TournamentCardModeEditCallbackMessage
	TournamentCardModeInlineQuery
)

func RenderTournamentCard(c context.Context, cardMode TournamentCardMode, tournament pamodels.Tournament) (m bots.MessageFromBot, err error) {
	m.IsEdit = cardMode == TournamentCardModeEditCallbackMessage
	m.Format = bots.MessageFormatHTML
	m.DisableWebPagePreview = true
	text := new(bytes.Buffer)
	if cardMode == TournamentCardModeInlineQuery {
		fmt.Fprintf(text, `⚔ <b>Tournament</b>: <a href="https://t.me/prizarena_bot?start=%v">%v</a>`, tournament.ID, tournament.Name)
		fmt.Fprint(text, "\n")
	} else {
		fmt.Fprintf(text, "⚔ <b>Tournament</b>: %v\n", tournament.Name)
	}

	{
		gameName := tournament.GameName
		if gameName == "" {
			gameName = tournament.GameID
		}
		fmt.Fprintf(text, `🎮 <b>Game</b>: <a href="https://t.me/prizarena_bot?start=%v">%v</a>`, tournament.GameID, gameName)
		fmt.Fprintln(text, "")
	}

	if tournament.Note != "" {
		fmt.Fprintln(text, tournament.Note)
	}

	if (tournament.Status == "" || tournament.Status == "active") && !tournament.Ends.IsZero() && tournament.Ends.Before(time.Now()) {
		if tournament, err = papfacade.Tournaments.CloseTournament(c, tournament.ID); err != nil {
			return
		}
	}

	fmt.Fprintln(text, "\n<b>Status</b>:", tournament.Status)
	if tournament.IsListed {
		fmt.Fprintln(text, "Publicly listed at https://prizarena.com/tournaments")
	}

	if tournament.DurationDays > 0 {
		fmt.Fprintf(text, "<b>Duration</b>: %d days\n", tournament.DurationDays)
	}
	const dtFormat = "2006-01-02 15:04"
	if !tournament.Starts.IsZero() && !tournament.Ends.IsZero() {
		fmt.Fprintf(text, "<b>Takes place</b>: from %v till %v\n", tournament.Starts.Format(dtFormat), tournament.Ends.Format(dtFormat))
	} else if !tournament.Starts.IsZero() {
		fmt.Fprintf(text, "<b>Takes place</b>: from %v\n", tournament.Starts.Format(dtFormat))
	} else if !tournament.Ends.IsZero() {
		fmt.Fprintf(text, "<b>Takes place</b>: till %v\n", tournament.Ends.Format(dtFormat))
	}

	if tournament.Sponsorship != "" {
		sponsorship := tournament.GetSponsorship()
		if sponsorship.Sponsor.Name != "" {
			fmt.Fprintln(text, "<b>Sponsored by</b>:", sponsorship.Sponsor.Name)
		}
		if sponsorship.Sponsor.Url != "" {
			fmt.Fprintln(text, sponsorship.Sponsor.Url)
		}
		if sponsorship.Sponsor.About != "" {
			fmt.Fprintln(text, sponsorship.Sponsor.About)
		}
		if sponsorship.Prize != nil {
			fmt.Fprintln(text, "<b>Prizes by place:</b>")
			for i, prize := range sponsorship.Prize.ByPlace {
				fmt.Fprintf(text, "  #%d - %v %v\n", i+1, prize.Value, prize.Currency)
			}
			if sponsorship.Prize.RandomsCount > 0 {
				fmt.Fprintf(text, "<b>Prize to %v random contestant(s):</b> %v %v\n", sponsorship.Prize.RandomsCount, sponsorship.Prize.RandomAmount.Value, sponsorship.Prize.RandomAmount.Currency)
			}
		}
	}

	if tournament.Status != "draft" {
		fmt.Fprintln(text, "Contestants:", tournament.CountOfContestants)
		fmt.Fprintln(text, "Games played:", tournament.CountOfPlaysCompleted)
	}
	m.Text = text.String()
	if m.Keyboard, err = getTournamentInGameTelegramKeyboard(c, cardMode, tournament); err != nil {
		return
	}
	return
}

func getTournamentInGameTelegramKeyboard(c context.Context, _ TournamentCardMode, tournament pamodels.Tournament) (keyboard *tgbotapi.InlineKeyboardMarkup, err error) {
	shortTournamentID := tournament.ShortTournamentID()
	switchInlineQueryTournament := "tournament?id=" + shortTournamentID
	switchInlineQueryPlay := "play?tournament=" + shortTournamentID
	keyboard = tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{
			{Text: "⚔ Play against friend", SwitchInlineQuery: &switchInlineQueryPlay},
		},
		[]tgbotapi.InlineKeyboardButton{
			{Text: "👽 Play against stranger", CallbackData: "play-stranger?t=" + shortTournamentID},
		},
		[]tgbotapi.InlineKeyboardButton{
			{Text: "✈ Share in Telegram", SwitchInlineQuery: &switchInlineQueryTournament},
		},
	)

	if !tournament.IsSponsored {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
			[]tgbotapi.InlineKeyboardButton{
				{Text: "💵 Become a Sponsor", URL: "https://t.me/prizarena_bot?start=sponsor__t-" + tournament.ID},
			},
		)
	}

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
		[]tgbotapi.InlineKeyboardButton{
			{Text: "🔄 Refresh", CallbackData: refreshTournamentCallbackData(tournament.ID)},
		},
	)
	return
}
