package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPageHandler(c *gin.Context) {
	data := gin.H{
		"Title": "Welcome to the Games Index",
	}
	c.HTML(http.StatusOK, "index.html", data)
}
