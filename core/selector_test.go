package core

import (
	"testing"

	"github.com/chenhu1001/marketool/datacenter/eastmoney"
	"github.com/chenhu1001/marketool/logging"
	"github.com/stretchr/testify/require"
)

func TestAutoFilterStocks(t *testing.T) {
	logging.SetLevel("error")
	checker := NewChecker(_ctx, DefaultCheckerOptions)
	s := NewSelector(_ctx, eastmoney.DefaultFilter, checker)
	_, err := s.AutoFilterStocks(_ctx)
	require.Nil(t, err)
	// t.Log(result)
}
