// 主页

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Index godoc
func Index(c *gin.Context) {
	data := gin.H{
		"Env":       viper.GetString("env"),
		"PageTitle": "Marketool | 主页",
	}
	c.HTML(http.StatusOK, "index.html", data)
	return
}
