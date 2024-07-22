package frontend

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func DocumentationPageHandler(c *gin.Context) {
	mdContent, err := os.ReadFile("README.md")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read documentation")
		return
	}
	htmlContent := blackfriday.MarkdownCommon(mdContent)

	c.HTML(http.StatusOK, "documentation.html", gin.H{
		"Content": template.HTML(htmlContent),
	})

}
