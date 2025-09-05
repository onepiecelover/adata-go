// Package fund 提供基金相关功能
package fund

import (
	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// Fund 基金模块结构体
type Fund struct {
	client *client.Client
}

// New 创建基金模块实例
func New() *Fund {
	return &Fund{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func SetProxy(enabled bool, proxyURL string) {
	fund := New()
	fund.client.SetProxy(enabled, proxyURL)
}

// SetProxy 设置代理
func (f *Fund) SetProxy(enabled bool, proxyURL string) {
	f.client.SetProxy(enabled, proxyURL)
}

// AllETFExchangeTradedInfo 获取所有场内ETF信息
func (f *Fund) AllETFExchangeTradedInfo() ([]types.ETFInfo, error) {
	// ETF信息获取功能待实现
	return nil, errors.NewADataError(70001, "ETF信息获取功能待实现", "")
}

// GetETFMarket 获取ETF行情数据
func (f *Fund) GetETFMarket(etfCode, startDate, endDate string, kType int) ([]types.MarketData, error) {
	// ETF行情数据获取功能待实现
	return nil, errors.NewADataError(70002, "ETF行情数据获取功能待实现", "")
}

// GetETFMarketCurrent 获取ETF当前行情
func (f *Fund) GetETFMarketCurrent(etfCodes []string) ([]types.CurrentMarket, error) {
	// ETF当前行情获取功能待实现
	return nil, errors.NewADataError(70003, "ETF当前行情获取功能待实现", "")
}
