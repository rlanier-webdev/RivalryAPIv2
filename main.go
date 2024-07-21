package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

type Team struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"unique;not null"`
}

type Game struct {
	ID            uint `gorm:"primaryKey"`
	HomeTeam      string
	AwayTeam      string
	Date          CustomDate
	HomeTeamScore int
	AwayTeamScore int
	Notes         string
}

func initDB() {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database: ", err)
		}

		err = db.AutoMigrate(&Team{}, &Game{})
		if err != nil {
			log.Fatal("failed to migrate database: ", err)
		}
	})
}

func main() {
	initDB()

	r := gin.Default()

	r.POST("/api/teams", createTeamHandler)
	r.GET("/api/teams", getTeamsHandler)

	r.GET("/api/games", getGamesHandler)
	r.GET("/api/games/:id", getGameByIDHandler)
	r.GET("/api/games/year/:year", getGamesByYearHandler)

	if err := r.Run(":1889"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
