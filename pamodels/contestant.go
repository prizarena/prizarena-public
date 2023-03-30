package pamodels

import (
	"fmt"
	"github.com/prizarena/arena/models4arena"
	"github.com/strongo/dalgo/record"
	"strings"
	"time"
)

var ContestantKind = "Contestant"

type ContestantEntity struct {
	// Keep in private repository?
	TimeJoined      time.Time
	TournamentID    string
	GameUserID      string    `datastore:",omitempty"` // can be empty when joining exclusive tournament from @prizarena_bot inline callback
	PrizarenaUserID string    `datastore:",omitempty"`
	StrangerCreated time.Time `datastore:",omitempty"`
	StrangerPairing time.Time `datastore:",omitempty"`
	models4arena.ContestantStats
}

type Contestant struct {
	record.WithID[string]
	*ContestantEntity
}

//var _ db.EntityHolder = (*Contestant)(nil)

type ContestantID string

const contestantIdSeparator = ":"

func (id ContestantID) UserID() string {
	s := string(id)
	if i := strings.Index(s, contestantIdSeparator); i > 0 {
		return s[:i]
	}
	return s
}

func NewContestantID(tournamentID, userID string) string {
	if tournamentID == "" {
		panic("tournamentID is required")
	}
	if count := strings.Count(tournamentID, ":"); count != 1 {
		panic(fmt.Sprintf("tournamentID should contains exactly one ':' charater, got: %v", count))
	}
	v := strings.Split(tournamentID, ":")
	if v[0] == "" {
		panic("gameID is not specified")
	}
	if v[1] == "" {
		panic("gameTournamentID is not specified")
	}
	return tournamentID + contestantIdSeparator + userID
}

func NewContestant(tournamentID, userID string) Contestant {
	return Contestant{
		//StringID: db.NewStrID(NewContestantID(tournamentID, userID)),
	}
}
