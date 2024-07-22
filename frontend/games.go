package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GamesPageHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Games",
	}
	c.HTML(http.StatusOK, "games.html", data)
}
