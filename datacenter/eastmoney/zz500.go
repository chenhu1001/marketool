// 中证500成分股

package eastmoney

import (
	"context"
	"fmt"

	"github.com/chenhu1001/marketool/goutils"
	"github.com/corpix/uarand"
)

// ZZ500Item 中证500成分股信息
type ZZ500Item struct {
	Secucode         string      `json:"SECUCODE"`           // 股票代码.XX
	SecurityCode     string      `json:"SECURITY_CODE"`      // 股票代码
	SecurityNameAbbr string      `json:"SECURITY_NAME_ABBR"` // 股票简称
	ClosePrice       float64     `json:"CLOSE_PRICE"`        // 最新价格
	Industry         string      `json:"INDUSTRY"`           // 主营行业
	Region           string      `json:"REGION"`             // 地区
	Weight           float64     `json:"WEIGHT"`             // 持仓比例（%）
	Eps              float64     `json:"EPS"`                // 每股收益
	Bps              float64     `json:"BPS"`                // 每股净资产
	Roe              float64     `json:"ROE"`                // 净资产收益率
	TotalShares      float64     `json:"TOTAL_SHARES"`       // 总股本（亿股）
	FreeShares       float64     `json:"FREE_SHARES"`        // 流通股本（亿股）
	FreeCap          float64     `json:"FREE_CAP"`           // 流通市值（亿元）
	Type             string      `json:"TYPE"`
	F2               interface{} `json:"f2"`
	F3               interface{} `json:"f3"`
}

// RspZZ500 ZZ500接口返回结构
type RspZZ500 struct {
	Version string `json:"version"`
	Result  struct {
		Pages int         `json:"pages"`
		Data  []ZZ500Item `json:"data"`
		Count int         `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ZZ500 返回中证500成分股列表
func (e EastMoney) ZZ500(ctx context.Context) (results []ZZ500Item, err error) {
	apiurl := "https://datacenter-web.eastmoney.com/api/data/v1/get?sortColumns=ROE&sortTypes=-1&pageSize=500&pageNumber=1&reportName=RPT_INDEX_TS_COMPONENT&columns=SECUCODE%2CSECURITY_CODE%2CTYPE%2CSECURITY_NAME_ABBR%2CCLOSE_PRICE%2CINDUSTRY%2CREGION%2CWEIGHT%2CEPS%2CBPS%2CROE%2CTOTAL_SHARES%2CFREE_SHARES%2CFREE_CAP&quoteColumns=f2%2Cf3&quoteType=0&source=WEB&client=WEB&filter=(TYPE%3D%223%22)"
	header := map[string]string{
		"user-agent": uarand.GetRandom(),
	}
	rsp := RspZZ500{}
	if err := goutils.HTTPGET(ctx, e.HTTPClient, apiurl, header, &rsp); err != nil {
		return nil, err
	}
	if rsp.Code != 0 {
		return nil, fmt.Errorf("ZZ500 rsp code error, rsp:%+v", rsp)
	}
	if len(rsp.Result.Data) != 500 {
		return nil, fmt.Errorf("ZZ500 rsp data len != 500, len=%d", len(rsp.Result.Data))
	}
	return rsp.Result.Data, nil
}
