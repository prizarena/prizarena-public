package pamodels

import (
	"github.com/strongo/db"
	"github.com/strongo/decimal"
	"strings"
	"fmt"
)

const GameKind = "G"

type Game struct {
	db.StringID
	*GameEntity
}

var _ db.EntityHolder = (*Game)(nil)

type GameEntity struct {
	Name                     string              `datastore:",omitempty"`
	URL                      string              `datastore:",omitempty"`
	PWA                      string              `datastore:",omitempty"`
	TelegramBot              string              `datastore:",omitempty"`
	Token                    string              `datastore:",omitempty"`
	CountOfActiveTournaments int                 `datastore:",omitempty"`
	CountOfClosedTournaments int                 `datastore:",omitempty"`
	PotSizeInUsdCents        decimal.Decimal64p2 `datastore:",omitempty"`
}

func (Game) Kind() string {
	return GameKind
}

func (g Game) Entity() interface{} {
	return g.GameEntity
}

func (Game) NewEntity() interface{} {
	return new(GameEntity)
}

func (g *Game) SetEntity(entity interface{}) {
	g.GameEntity = entity.(*GameEntity)
}

func (ge GameEntity) GameURL() string {
	switch {
	case ge.TelegramBot != "" && !strings.HasPrefix(ge.URL, "https://t.me/"):
		return fmt.Sprintf("https://telegram.me/%v?start=ref-prizarena_bot", ge.TelegramBot)
	case ge.URL != "":
		return ge.URL
	}
	return ""
}