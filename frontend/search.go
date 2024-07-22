package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchPageHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Search Games",
	}
	c.HTML(http.StatusOK, "search.html", data)
}
