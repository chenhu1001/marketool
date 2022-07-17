// Package cron 定时任务
package cron

import (
	"context"

	"github.com/chenhu1001/marketool/datacenter"
	"github.com/chenhu1001/marketool/goutils"
	"github.com/chenhu1001/marketool/models"
)

// SyncBond 同步债券
func SyncBond() {
	if !goutils.IsTradingDay() {
		return
	}
	ctx := context.Background()
	syl := datacenter.ChinaBond.QueryAAACompanyBondSyl(ctx)
	if syl != 0 {
		models.AAACompanyBondSyl = syl
	}
}
