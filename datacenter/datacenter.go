// Package datacenter 数据来源
package datacenter

import (
	"github.com/chenhu1001/marketool/datacenter/chinabond"
	"github.com/chenhu1001/marketool/datacenter/eastmoney"
	"github.com/chenhu1001/marketool/datacenter/eniu"
	"github.com/chenhu1001/marketool/datacenter/sina"
	"github.com/chenhu1001/marketool/datacenter/zszx"
)

var (
	// EastMoney 东方财富
	EastMoney eastmoney.EastMoney
	// Eniu 亿牛网
	Eniu eniu.Eniu
	// Sina 新浪财经
	Sina sina.Sina
	// Zszx 招商证券
	Zszx zszx.Zszx
	// ChinaBond 中国债券信息网
	ChinaBond chinabond.ChinaBond
)

func init() {
	EastMoney = eastmoney.NewEastMoney()
	Eniu = eniu.NewEniu()
	Sina = sina.NewSina()
	Zszx = zszx.NewZszx()
	ChinaBond = chinabond.NewChinaBond()
}
