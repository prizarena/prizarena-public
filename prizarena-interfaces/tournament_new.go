package prizarena_interfaces

import (
	"time"
	"github.com/strongo/decimal"
	"net/url"
)

type NewTournament struct {
	GameID string
	Name string
	Starts time.Time
	Ends time.Time
}

type Prize struct {
	Name string
	Medium string
	Currency string
	Value decimal.Decimal64p2
}

type Sponsor struct {
	Name string
	Url url.URL
	Text string
	Prize Prize
}

