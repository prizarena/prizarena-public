package pamodels

import (
	"github.com/pkg/errors"
	"github.com/strongo/app"
	"github.com/strongo/app/user"
	"github.com/strongo/dalgo/record"
	"time"
)

const (
	UserKind = "User"
	//strangerRivalBidKey = "$tranger"
)

type UserEntity struct {
	strongo.AppUserBase
	user.AccountsOfUser

	Name        string `datastore:",noindex,omitempty"`
	Created     time.Time
	AvatarURL   string `datastore:",noindex,omitempty"`
	FirebaseUID string `datastore:",omitempty"`
	Tokens      int
	//
	//
	TournamentIDs []string `datastore:",noindex"`
	//BattlesHandler
}

type User struct {
	record.WithID[string]
	*UserEntity
}

func (u *UserEntity) SetBotUserID(platform, botID, botUserID string) {
	u.AddAccount(user.Account{
		Provider: platform,
		App:      botID,
		ID:       botUserID,
	})
}

var ErrNotEnoughTokens = errors.New("not enough tokens")
