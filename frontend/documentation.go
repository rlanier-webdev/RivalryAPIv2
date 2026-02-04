package frontend

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func DocumentationPageHandler(c *gin.Context) {
	mdContent, err := os.ReadFile("README.md")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read documentation")
		return
	}

	// Convert markdown to HTML
	unsafeHTML := blackfriday.MarkdownCommon(mdContent)

	// Sanitize HTML to prevent XSS
	policy := bluemonday.UGCPolicy()
	safeHTML := policy.SanitizeBytes(unsafeHTML)

	c.HTML(http.StatusOK, "documentation.html", gin.H{
		"Content": template.HTML(safeHTML),
	})
}
