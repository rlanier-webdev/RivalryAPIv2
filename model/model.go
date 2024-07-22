package model

import "github.com/rlanier-webdev/RivalryAPIv2/utilities"

type Game struct {
	ID            uint `gorm:"primaryKey"`
	HomeTeam      string
	AwayTeam      string
	Date          utilities.CustomDate
	HomeTeamScore int
	AwayTeamScore int
	Notes         string
}
