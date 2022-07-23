// 关于

package routes

import (
	"net/http"

	"github.com/chenhu1001/marketool/version"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// About godoc
func About(c *gin.Context) {
	data := gin.H{
		"Env":       viper.GetString("env"),
		"Version":   version.Version,
		"PageTitle": "Marketool | 关于",
	}
	c.HTML(http.StatusOK, "about.html", data)
	return
}
