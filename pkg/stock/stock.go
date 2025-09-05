// Package stock 提供股票相关功能
package stock

import (
	"github.com/onepiecelover/adata-go/pkg/stock/finance"
	"github.com/onepiecelover/adata-go/pkg/stock/info"
	"github.com/onepiecelover/adata-go/pkg/stock/market"
)

// Stock 股票模块结构体
type Stock struct {
	Info    *info.StockInfo
	Market  *market.StockMarket
	Finance *finance.StockFinance
}

// New 创建股票模块实例
func New() *Stock {
	return &Stock{
		Info:    info.NewStockInfo(),
		Market:  market.NewStockMarket(),
		Finance: finance.NewStockFinance(),
	}
}

// SetProxy 设置代理
func SetProxy(enabled bool, proxyURL string) {
	// 设置所有子模块的代理
	stock := New()
	stock.Info.SetProxy(enabled, proxyURL)
	stock.Market.SetProxy(enabled, proxyURL)
	stock.Finance.SetProxy(enabled, proxyURL)
}
