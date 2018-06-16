package pamodels

//go:generate ffjson $GOFILE

import (
	"github.com/crediterra/money"
)

type TournamentSponsorJson struct {
	Name  string `json:",omitempty"`
	Url   string `json:",omitempty"`
	About string `json:",omitempty"`
}

type TournamentPrizeJson struct {
	Medium       string         `json:",omitempty"`
	ByPlace      []money.Amount `json:",omitempty"`
	RandomAmount money.Amount   `json:",omitempty"`
	RandomsCount int            `json:",omitempty"`
}

type TournamentSponsorshipJson struct {
	Sponsor *TournamentSponsorJson
	Prize   *TournamentPrizeJson
}
