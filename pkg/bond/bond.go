// Package bond 提供债券相关功能
package bond

import (
	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// Bond 债券模块结构体
type Bond struct {
	client *client.Client
}

// New 创建债券模块实例
func New() *Bond {
	return &Bond{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func SetProxy(enabled bool, proxyURL string) {
	bond := New()
	bond.client.SetProxy(enabled, proxyURL)
}

// SetProxy 设置代理
func (b *Bond) SetProxy(enabled bool, proxyURL string) {
	b.client.SetProxy(enabled, proxyURL)
}

// AllBondCode 获取所有债券代码
func (b *Bond) AllBondCode() ([]types.BondInfo, error) {
	// 债券代码获取功能待实现
	return nil, errors.NewADataError(80001, "债券代码获取功能待实现", "")
}

// GetBondMarket 获取债券行情数据
func (b *Bond) GetBondMarket(bondCode, startDate, endDate string, kType int) ([]types.MarketData, error) {
	// 债券行情数据获取功能待实现
	return nil, errors.NewADataError(80002, "债券行情数据获取功能待实现", "")
}

// GetBondMarketCurrent 获取债券当前行情
func (b *Bond) GetBondMarketCurrent(bondCodes []string) ([]types.CurrentMarket, error) {
	// 债券当前行情获取功能待实现
	return nil, errors.NewADataError(80003, "债券当前行情获取功能待实现", "")
}
