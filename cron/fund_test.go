package cron

import (
	"testing"

	"github.com/chenhu1001/marketool/logging"
	"github.com/spf13/viper"
)

func _TestSyncFund(t *testing.T) {
	logging.SetLevel("warn")
	viper.SetDefault("app.chan_size", 500)
	SyncFund()
}

func _TestSyncFundManagers(t *testing.T) {
	logging.SetLevel("warn")
	viper.SetDefault("app.chan_size", 500)
	SyncFundManagers()
}
