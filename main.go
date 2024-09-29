package main

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/rlanier-webdev/RivalryAPIv2/frontend"
	"github.com/rlanier-webdev/RivalryAPIv2/models"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	once sync.Once
)

func initDB() {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect to database: ", err)
		}

		err = db.AutoMigrate(&models.Game{})
		if err != nil {
			log.Fatal("failed to migrate database: ", err)
		}

		// Load games from the database
		var games []models.Game
		err = db.Find(&games).Error
		if err != nil {
			log.Fatal("failed to load games from the database: ", err)
		}

		// Set the loaded games in the frontend package
		frontend.SetGames(games)
	})
}

func main() {
	initDB()
	frontend.SetDB(db)

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", frontend.IndexPageHandler)
	r.GET("/search", frontend.SearchPageHandler)
	r.GET("/docs", frontend.DocumentationPageHandler)
	r.GET("/games", frontend.GamesPageHandler)

	r.GET("/api/games", getGamesHandler)
	r.GET("/api/games/:id", getGameByIDHandler)
	r.GET("/api/games/year/:year", getGamesByYearHandler)
	r.GET("/api/games/home/:team", getGamesByHomeHandler)
	r.GET("/api/games/away/:team", getGamesByAwayHandler)

	if err := r.Run(":1889"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
