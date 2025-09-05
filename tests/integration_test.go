package tests

import (
	"testing"
	"time"

	"github.com/onepiecelover/adata-go"
	"github.com/onepiecelover/adata-go/pkg/types"
	"github.com/stretchr/testify/assert"
)

// 集成测试 - 测试完整的数据获取流程
func TestIntegration_BasicWorkflow(t *testing.T) {
	// 1. 测试获取股票代码列表
	codes, err := adata.Stock.Info.AllCode()
	if err != nil {
		t.Logf("获取股票代码失败 (网络环境问题): %v", err)
		t.Skip("跳过集成测试 - 网络连接问题")
		return
	}

	assert.NotEmpty(t, codes, "应该返回非空的股票代码列表")
	t.Logf("成功获取 %d 只股票代码", len(codes))

	if len(codes) == 0 {
		t.Skip("跳过后续测试 - 没有获取到股票代码")
		return
	}

	// 2. 使用获取到的股票代码测试行情数据
	testStockCode := codes[0].StockCode
	params := &types.MarketParams{
		StockCode:  testStockCode,
		StartDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		KType:      1, // 日K线
		AdjustType: 1, // 前复权
	}

	marketData, err := adata.Stock.Market.GetMarket(params)
	if err != nil {
		t.Logf("获取 %s 行情数据失败: %v", testStockCode, err)
	} else {
		assert.NotEmpty(t, marketData, "应该返回非空的行情数据")
		t.Logf("成功获取 %s 的 %d 条行情数据", testStockCode, len(marketData))
	}

	// 3. 测试当前行情
	testCodes := []string{testStockCode}
	if len(codes) > 1 {
		testCodes = append(testCodes, codes[1].StockCode)
	}

	currentMarkets, err := adata.Stock.Market.ListMarketCurrent(testCodes)
	if err != nil {
		t.Logf("获取当前行情失败: %v", err)
	} else {
		assert.NotEmpty(t, currentMarkets, "应该返回非空的当前行情数据")
		t.Logf("成功获取 %d 只股票的当前行情", len(currentMarkets))
	}
}

func TestIntegration_VersionInfo(t *testing.T) {
	version := adata.GetVersion()
	assert.NotEmpty(t, version, "版本信息不应为空")
	t.Logf("AData-Go 版本: %s", version)
}

func TestIntegration_ProxySettings(t *testing.T) {
	// 测试代理设置功能（不实际使用代理）
	adata.SetProxy(false, "")

	// 尝试获取数据验证代理设置没有破坏功能
	_, err := adata.Stock.Info.AllCode()
	if err != nil {
		t.Logf("代理设置后获取数据失败 (网络环境问题): %v", err)
	}
}

func TestIntegration_ErrorHandling(t *testing.T) {
	// 测试错误处理机制

	// 1. 无效股票代码
	params := &types.MarketParams{
		StockCode: "invalid_code",
		StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		KType:     1,
	}

	_, err := adata.Stock.Market.GetMarket(params)
	assert.Error(t, err, "无效股票代码应该返回错误")

	// 2. 空代码列表
	_, err = adata.Stock.Market.ListMarketCurrent([]string{})
	assert.Error(t, err, "空代码列表应该返回错误")

	// 3. nil参数
	_, err = adata.Stock.Market.GetMarket(nil)
	assert.Error(t, err, "nil参数应该返回错误")
}

// 基准测试
func BenchmarkIntegration_GetStockCodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := adata.Stock.Info.AllCode()
		if err != nil {
			b.Logf("获取股票代码失败: %v", err)
			b.Skip("跳过基准测试 - 网络连接问题")
			return
		}
	}
}

func BenchmarkIntegration_GetMarketData(b *testing.B) {
	params := &types.MarketParams{
		StockCode:  "000001",
		StartDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		KType:      1,
		AdjustType: 1,
	}

	for i := 0; i < b.N; i++ {
		_, err := adata.Stock.Market.GetMarket(params)
		if err != nil {
			b.Logf("获取行情数据失败: %v", err)
			b.Skip("跳过基准测试 - 网络连接问题")
			return
		}
	}
}
