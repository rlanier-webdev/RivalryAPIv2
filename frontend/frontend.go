package frontend

import (
	"github.com/rlanier-webdev/RivalryAPIv2/model"
	"gorm.io/gorm"
)

var db *gorm.DB

var games []model.Game

func SetDB(database *gorm.DB) {
	db = database
}

func SetGames(g []model.Game) {
	games = g
}