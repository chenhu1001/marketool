// 主页

package routes

import (
	"github.com/chenhu1001/marketool/cron"
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

// SyncFund godoc
func SyncFund(c *gin.Context) {
	go cron.SyncFund()

	c.JSON(http.StatusOK, "SyncFund Success.")
	return
}

// SyncIndustryList godoc
func SyncIndustryList(c *gin.Context) {
	go cron.SyncIndustryList()

	c.JSON(http.StatusOK, "SyncIndustryList Success.")
	return
}

// SyncFundManagers godoc
func SyncFundManagers(c *gin.Context) {
	go cron.SyncFundManagers()

	c.JSON(http.StatusOK, "SyncFundManagers Success.")
	return
}
