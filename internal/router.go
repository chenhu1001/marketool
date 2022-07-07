package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func Initialize() {
	Router = gin.Default()
	Router.Static("/static/", "./assets/static")
	Router.LoadHTMLGlob("./views/templates/*")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
