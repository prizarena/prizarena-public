package pabot

import (
	// "bitbucket.org/asterus/prizarena-private/prizarena-server/papfacade"
	"bytes"
	"context"
	"fmt"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/strongo/bots-api-telegram"
	"github.com/strongo/bots-framework/botsfw"
	"time"
)

type TournamentCardMode int

const (
	TournamentCardModeNewMessage TournamentCardMode = iota
	TournamentCardModeEditCallbackMessage
	TournamentCardModeInlineQuery
)

func RenderTournamentCard(c context.Context, cardMode TournamentCardMode, tournament pamodels.Tournament) (m botsfw.MessageFromBot, err error) {
	m.IsEdit = cardMode == TournamentCardModeEditCallbackMessage
	m.Format = botsfw.MessageFormatHTML
	m.DisableWebPagePreview = true
	text := new(bytes.Buffer)
	if cardMode == TournamentCardModeInlineQuery {
		fmt.Fprintf(text, `⚔ <b>Tournament</b>: <a href="https://t.me/prizarena_bot?start=%v">%v</a>`, tournament.ID, tournament.Data.Name)
		fmt.Fprint(text, "\n")
	} else {
		fmt.Fprintf(text, "⚔ <b>Tournament</b>: %v\n", tournament.Data.Name)
	}

	{
		gameName := tournament.Data.GameName
		if gameName == "" {
			gameName = tournament.Data.GameID
		}
		fmt.Fprintf(text, `🎮 <b>Game</b>: <a href="https://t.me/prizarena_bot?start=%v">%v</a>`, tournament.Data.GameID, gameName)
		fmt.Fprintln(text, "")
	}

	if tournament.Data.Note != "" {
		fmt.Fprintln(text, tournament.Data.Note)
	}

	if (tournament.Data.Status == "" || tournament.Data.Status == "active") && !tournament.Data.Ends.IsZero() && tournament.Data.Ends.Before(time.Now()) {
		panic("not implemented yet")
		//if tournament, err = papfacade.Tournaments.CloseTournament(c, tournament.ID); err != nil {
		//	return
		//}
	}

	fmt.Fprintln(text, "\n<b>Status</b>:", tournament.Data.Status)
	if len(tournament.Data.ExclusiveTo) > 0 {
		fmt.Fprintf(text, "Exclusive to: %v\n", tournament.Data.ExclusiveTo)
	}

	if tournament.Data.DurationDays > 0 {
		fmt.Fprintf(text, "<b>Duration</b>: %d days", tournament.Data.DurationDays)
		const dtFormat = "2006-01-02 15:04"
		now := time.Now()
		if !tournament.Data.Starts.IsZero() && !tournament.Data.Ends.IsZero() {

			if tournament.Data.Starts.After(now) {
				fmt.Fprintf(text, ", starts %v ", tournament.Data.Starts.Format(dtFormat))
			} else if tournament.Data.Ends.After(now) {
				fmt.Fprintf(text, ", ends %v ", tournament.Data.Ends.Format(dtFormat))
			} else if tournament.Data.Ends.Before(now) {
				fmt.Fprintf(text, ", was hold from %v till %v", tournament.Data.Starts.Format(dtFormat), tournament.Data.Ends.Format(dtFormat))
			}
		} else if !tournament.Data.Starts.IsZero() {
			fmt.Fprintf(text, ", starts</b>: %v", tournament.Data.Starts.Format(dtFormat))
		} else if !tournament.Data.Ends.IsZero() {
			fmt.Fprintf(text, ", ends</b>: %v", tournament.Data.Ends.Format(dtFormat))
		}
		fmt.Fprintln(text, "")
	}

	if tournament.Data.Sponsorship != "" {
		sponsorship := tournament.Data.GetSponsorship()
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

	//if tournament.Status != "draft" {
	//	fmt.Fprintln(text, "Contestants:", tournament.CountOfContestants)
	//	fmt.Fprintln(text, "Games played:", tournament.CountOfPlaysCompleted)
	//}

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
			{Text: "👽 Play against stranger", CallbackData: getPlayStrangerCallbackData(shortTournamentID)},
		},
		[]tgbotapi.InlineKeyboardButton{
			{Text: "✈ Share in Telegram", SwitchInlineQuery: &switchInlineQueryTournament},
		},
	)

	if !tournament.Data.IsSponsored {
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
