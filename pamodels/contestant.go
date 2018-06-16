package pamodels

import (
	"github.com/strongo/db"
	"strings"
	"time"
	"fmt"
	"github.com/prizarena/arena/arena-go"
)

var ContestantKind = "Contestant"

type ContestantEntity struct {
	// Keep in private repository?
	TimeJoined      time.Time
	TournamentID    string
	GameUserID      string
	PrizarenaUserID string    `datastore:",omitempty"`
	StrangerCreated time.Time `datastore:",omitempty"`
	StrangerPairing time.Time `datastore:",omitempty"`
	arena.ContestantStats
}

type Contestant struct {
	db.StringID
	*ContestantEntity
}

var _ db.EntityHolder = (*Contestant)(nil)

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
	vals := strings.Split(tournamentID, ":")
	if vals[0] == "" {
		panic("gameID is not specified")
	}
	if vals[1] == "" {
		panic("gameTournamentID is not specified")
	}
	return tournamentID + contestantIdSeparator + userID
}

func NewContestant(tournamentID, userID string) Contestant {
	return Contestant{
		StringID: db.NewStrID(NewContestantID(tournamentID, userID)),
	}
}

var _ db.EntityHolder = (*Contestant)(nil)

func (Contestant) Kind() string {
	return ContestantKind
}

func (Contestant) NewEntity() interface{} {
	return new(ContestantEntity)
}

func (t Contestant) Entity() interface{} {
	return t.ContestantEntity
}

func (t *Contestant) SetEntity(v interface{}) {
	if v == nil {
		t.ContestantEntity = nil
	} else {
		t.ContestantEntity = v.(*ContestantEntity)
	}
}
