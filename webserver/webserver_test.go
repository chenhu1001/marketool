package webserver

import (
	"testing"

	"github.com/chenhu1001/marketool/goutils"
	"github.com/spf13/viper"
)

func TestViperConfig(t *testing.T) {
	InitWithConfigFile("../config.toml")
	defer viper.Reset()
	if !goutils.IsInitedViper() {
		t.Error("init viper failed")
	}
}
