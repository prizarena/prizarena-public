package padal

import (
	"context"
	"github.com/prizarena/prizarena-public/pamodels"
	"github.com/strongo/db"
)

func GetGameByID(c context.Context, gameID string) (game pamodels.Game, err error) {
	game.ID = gameID
	if err = DB.Get(c, &game); db.IsNotFound(err) {
		var found bool
		if game.GameEntity, found = games[gameID]; found {
			err = nil
		}
	}
	return
}

// func GetPopularGames(c context.Context) (games []pamodels.Game, err error) {
// 	for _, gameID := range []string{"rockpaperscissors"}
// 	games = []pamodels.Game{
// 		games[],
// 	}
// 	return
// }

var games = map[string]*pamodels.GameEntity{ // Just temporary
	"rockpaperscissors": {
		Name:  "Rock-Paper-Scissors 💎📄✂️",
		TelegramBot: "playRockPaperScissorsBot",
		Token: "r-p-s",
	},
	"biddingtictactoe": {
		Name:  "Bidding Tic-Tac-Toe ❌💰⭕",
		TelegramBot: "BiddingTicTacToeBot",
		Token: "b-ttt",
	},
	"matchingpennies": {
		Name:  "Matching Pennies",
		Token: "m-p",
		TelegramBot: "MatchingPenniesBot",
	},
	"greedgame": {
		Name:  "Greed Game",
		Token: "g-g",
		TelegramBot: "GreedGameBot",
	},
}
