package models

import "github.com/rlanier-webdev/RivalryAPIv2/utils"

type Game struct {
	ID            uint `gorm:"primaryKey"`
	HomeTeam      string
	AwayTeam      string
	Date          utils.CustomDate
	HomeTeamScore int
	AwayTeamScore int
	Notes         string
}
