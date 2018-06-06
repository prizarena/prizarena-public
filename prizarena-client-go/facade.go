package prizarena

import (
					"context"
	"github.com/prizarena/prizarena-public/prizarena-interfaces"
)

type facade struct {
	apiClient prizarena_interfaces.ApiClient
}

func NewFacade(apiClient prizarena_interfaces.ApiClient) facade {
	return facade{apiClient: apiClient}
}



func (facade facade) MakeMoveAgainstStranger(
	c context.Context,
	// now time.Time,
	tournamentID string,
	gameUserID string,
	move *prizarena_interfaces.MoveDto,
	onRivalFound func(rivalUserID string, rivalMove *prizarena_interfaces.MoveDto) error,
	onStranger func() error,
) (err error) {
	if move == nil && onStranger == nil {
		panic("Either 'move' or 'onStranger' should be defined.")
	}
	pairPayload := prizarena_interfaces.PairWithStrangerPayload{
		TournamentID: tournamentID,
		GameUserID: gameUserID,
	}

	response ,err := facade.apiClient.PairWithStranger(c, pairPayload)
	if err != nil {
		return err
	}

	if response.RivalGameUserID == "" {
		if onStranger != nil {
			onStranger()
		}
	} else {
		if err = onRivalFound(response.RivalGameUserID, response.RivalMove); err != nil {
			return err
		}
		pairedPayload := prizarena_interfaces.PairedWithStrangerPayload{
			GameUserID: gameUserID,
			RivalGameUserID: response.RivalGameUserID,
		}

		_, err := facade.apiClient.PairedWithStranger(c, pairedPayload)
		if err != nil {
			return err
		}
	}

	// var rivalUserIDs []string
	//
	// contestant := new(arena.Contestant)
	//
	// contestant.ID = arena.NewContestantID(tournamentID, gameUserID)
	//
	// if err = DB.Get(c, contestant); err != nil {
	// 	if db.IsNotFound(err) {
	// 		if err = DB.Get(c, user); err != nil {
	// 			return
	// 		}
	// 		rivalUserIDs = user.GetRivalUserIDs().Strings()
	// 	} else {
	// 		return
	// 	}
	// } else {
	// 	rivalUserIDs = contestant.RivalGameUserIDs.Strings()
	// }
	//
	// for {
	// 	var rivalUserID string
	// 	if rivalUserID, err = TournamentDAL.FindStranger(c, tournamentID, userID, rivalUserIDs); err != nil {
	// 		err = errors.WithMessage(err, "failed to find stranger")
	// 		return
	// 	}
	// 	log.Debugf(c, "strangerFacade.PlaceBidAgainstStranger() => rivalUserID: %v", rivalUserID)
	//
	// 	switch rivalUserID {
	// 	case userID:
	// 		err = errors.WithMessage(err, "FindStranger returned rivalUserID equal to current userID")
	// 		return
	// 	case "": // no strangers with existing open bids found
	// 		err = onStranger(contestant)
	// 		return
	// 	default: // Link 2 strangers
	//
	// 		if err = onRivalFound(rivalUserID); errors.Cause(err) == arena.ErrRivalUserIsNotBiddingAgainstStranger {
	// 			err = nil
	// 			continue
	// 		}
	// 		return
	// 	}
	// }
	return
}
