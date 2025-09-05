// Package sentiment 提供情感指标相关功能
package sentiment

import (
	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
)

// Sentiment 情感指标模块结构体
type Sentiment struct {
	client *client.Client
}

// New 创建情感指标模块实例
func New() *Sentiment {
	return &Sentiment{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func SetProxy(enabled bool, proxyURL string) {
	sentiment := New()
	sentiment.client.SetProxy(enabled, proxyURL)
}

// SetProxy 设置代理
func (s *Sentiment) SetProxy(enabled bool, proxyURL string) {
	s.client.SetProxy(enabled, proxyURL)
}

// GetHotList 获取热门板块
func (s *Sentiment) GetHotList() (interface{}, error) {
	// 热门板块获取功能待实现
	return nil, errors.NewADataError(90001, "热门板块获取功能待实现", "")
}

// GetNorthFlow 获取北向资金流向
func (s *Sentiment) GetNorthFlow() (interface{}, error) {
	// 北向资金流向获取功能待实现
	return nil, errors.NewADataError(90002, "北向资金流向获取功能待实现", "")
}

// GetSecuritiesMargin 获取融资融券数据
func (s *Sentiment) GetSecuritiesMargin() (interface{}, error) {
	// 融资融券数据获取功能待实现
	return nil, errors.NewADataError(90003, "融资融券数据获取功能待实现", "")
}

// GetStockLifting 获取限售解禁数据
func (s *Sentiment) GetStockLifting() (interface{}, error) {
	// 限售解禁数据获取功能待实现
	return nil, errors.NewADataError(90004, "限售解禁数据获取功能待实现", "")
}

// GetMineClearance 获取雷暴预警数据
func (s *Sentiment) GetMineClearance() (interface{}, error) {
	// 雷暴预警数据获取功能待实现
	return nil, errors.NewADataError(90005, "雷暴预警数据获取功能待实现", "")
}
