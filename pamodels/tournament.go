package pamodels

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/strongo/dalgo/record"
	"github.com/strongo/decimal"
	"strings"
	"time"
)

var TournamentKind = "Tournament"

const TournamentIDSeparator = ":"

type TournamentSponsorshipEntity struct {
	Sponsorship string `datastore:",noindex"`
	IsSponsored bool   `datastore:",omitempty"`
	PotSizeUSD  decimal.Decimal64p2
}

func (entity TournamentSponsorshipEntity) GetSponsorship() (v TournamentSponsorshipJson) {
	if entity.Sponsorship == "" {
		return v
	}
	if err := json.Unmarshal([]byte(entity.Sponsorship), &v); err != nil {
		panic(err)
	}
	return
}

type TournamentWithID struct {
	ID string
	TournamentEntity
}

type TournamentEntity struct {
	TournamentSponsorshipEntity
	CreatorGameUserID     string `datastore:",omitempty"`
	CreatorUserID         string
	GameID                string
	GameName              string `datastore:",noindex"`
	TgBotChatID           int64  `datastore:",noindex,omitempty"`
	TgBotMsgID            int64  `datastore:",noindex,omitempty"`
	Status                string
	Name                  string    `datastore:",noindex,omitempty"`
	Note                  string    `datastore:",noindex,omitempty"`
	Cached                time.Time `datastore:",omitempty"`
	Created               time.Time
	Starts                time.Time
	Ends                  time.Time `datastore:",omitempty"`
	DurationDays          int       `datastore:",noindex,omitempty"`
	MinGamesToScore       int       `datastore:",noindex,omitempty"`
	ExclusiveTo           []string  `datastore:",noindex"`
	CountOfContestants    int       `datastore:",noindex,omitempty"`
	CountOfPlaysCompleted int       `datastore:",noindex,omitempty"`
	LastPlayIDs           string    `datastore:",noindex,omitempty"`
}

func (j TournamentEntity) IsExclusive() bool {
	return len(j.ExclusiveTo) > 0
}

//const TournamentStarID = "*"

type Tournament struct {
	record.WithID[string]
	Data *TournamentEntity
}

//var _ db.EntityHolder = (*Tournament)(nil)

//func (Tournament) Kind() string {
//	return TournamentKind
//}

//func (Tournament) NewEntity() interface{} {
//	return new(TournamentEntity)
//}
//
//func (t Tournament) Entity() interface{} {
//	return t.Data
//}
//
//func (t *Tournament) SetEntity(v interface{}) {
//	if v == nil {
//		t.TournamentEntity = nil
//	} else {
//		t.TournamentEntity = v.(*TournamentEntity)
//	}
//}

const specialCharacter = "/\\.'\"\\<>"

func (v TournamentEntity) Validate() error {
	switch v.Status {
	case "":
		return errors.New("tournament.Status is a required field")
	case "draft", "active", "closed":
		return errors.New("tournament has unknown status: " + v.Status)
	}
	if v.Name == "" {
		return errors.New("tournament.Name is a required field")
	} else if len(v.Name) > 50 {
		return errors.New("tournament.Name is too long, max 50 characters allowed")
	} else if strings.ContainsAny(v.Name, specialCharacter) {
		return errors.New("tournament.Name is not allowed to contain special characters")
	}
	//if strings.Contains(v.ID, TournamentIDSeparator) {
	//	if t.Data.GameID == "" {
	//		return errors.New("tournament.GameID is a required field")
	//	} else if t.Data.GameID != strings.Split(t.ID, TournamentIDSeparator)[0] {
	//		return errors.New("tournament.GameID does not match 1st part of tournament.ID")
	//	}
	//} else if t.Data.GameID != "" {
	//	return errors.New("tournament.GameID must be empty if no GameID in 1st part of tournament")
	//}

	if v.GameName == "" {
		return errors.New("tournament.GameName is a required field")
	}
	if v.CreatorUserID == "" {
		return errors.New("tournament.CreatorUserID is a required field")
	}
	return nil
}

var ErrInvalidTournamentID = errors.New("invalid tournament ID")

func VerifyIsFullTournamentID(v string) error {
	if i := strings.Index(v, TournamentIDSeparator); i < 0 {
		return errors.WithMessage(ErrInvalidTournamentID, "tournament ID should have ':' character.")
	} else if i == 0 {
		return errors.WithMessage(ErrInvalidTournamentID, "tournament ID should have game ID before ':' character.")
	} else if i == len(v)-1 {
		return errors.WithMessage(ErrInvalidTournamentID, "tournament ID should have game tournament ID after ':' character.")
	}
	return nil
}

func (t Tournament) ShortTournamentID() string {
	if sepIndex := strings.Index(t.ID, TournamentIDSeparator); sepIndex >= 0 {
		return t.ID[sepIndex+1:]
	}
	return t.ID
}
