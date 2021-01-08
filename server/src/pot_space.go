package main

import "strconv"

type PotSpace struct {
	PointValue int
	CoinValue  int
	HasRuby    bool
}

func NewPotSpace(pointValue int, coinValue int, hasRuby bool) *PotSpace {
	return &PotSpace{pointValue, coinValue, hasRuby}
}

func (ps *PotSpace) Name() string {
	ruby := ""
	if ps.HasRuby {
		ruby = ", Ruby!"
	}
	return "Points:" + strconv.Itoa(ps.PointValue) + ", Coins:" + strconv.Itoa(ps.CoinValue) + ruby
}
