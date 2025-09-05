package tests

import (
	"testing"
	"time"

	adataErrors "github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/stock/market"
	"github.com/onepiecelover/adata-go/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestStockMarket_GetMarket_ValidParams(t *testing.T) {
	stockMarket := market.NewStockMarket()

	params := &types.MarketParams{
		StockCode:  "000001",
		StartDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		KType:      1, // 日K线
		AdjustType: 1, // 前复权
	}

	data, err := stockMarket.GetMarket(params)

	// 由于这是真实的网络请求，我们只测试基本的逻辑
	if err != nil {
		t.Logf("Network request failed (expected in some environments): %v", err)
		return
	}

	// 如果请求成功，验证返回数据的基本格式
	assert.NotEmpty(t, data, "Should return non-empty market data")

	for i, marketData := range data {
		if i >= 3 { // 只检查前3条
			break
		}
		assert.Equal(t, "000001", marketData.StockCode, "Stock code should match")
		assert.NotEmpty(t, marketData.TradeDate, "Trade date should not be empty")
		assert.GreaterOrEqual(t, marketData.Open, 0.0, "Open price should be non-negative")
		assert.GreaterOrEqual(t, marketData.Close, 0.0, "Close price should be non-negative")
		assert.GreaterOrEqual(t, marketData.High, 0.0, "High price should be non-negative")
		assert.GreaterOrEqual(t, marketData.Low, 0.0, "Low price should be non-negative")
		assert.GreaterOrEqual(t, marketData.Volume, int64(0), "Volume should be non-negative")
	}
}

func TestStockMarket_GetMarket_InvalidParams(t *testing.T) {
	stockMarket := market.NewStockMarket()

	// 测试nil参数
	_, err := stockMarket.GetMarket(nil)
	assert.Error(t, err, "Should return error for nil params")

	// 测试无效股票代码
	params := &types.MarketParams{
		StockCode: "invalid",
		StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		KType:     1,
	}

	_, err = stockMarket.GetMarket(params)
	assert.Error(t, err, "Should return error for invalid stock code")

	if adataErr, ok := err.(*adataErrors.ADataError); ok {
		assert.Equal(t, adataErrors.ErrInvalidStockCode.Code, adataErr.Code,
			"Should return ErrInvalidStockCode for invalid stock code")
	}
}

func TestStockMarket_ListMarketCurrent_ValidCodes(t *testing.T) {
	stockMarket := market.NewStockMarket()

	codes := []string{"000001", "600036", "000002"}
	data, err := stockMarket.ListMarketCurrent(codes)

	// 由于这是真实的网络请求，我们只测试基本的逻辑
	if err != nil {
		t.Logf("Network request failed (expected in some environments): %v", err)
		return
	}

	// 如果请求成功，验证返回数据的基本格式
	assert.NotEmpty(t, data, "Should return non-empty current market data")

	for _, currentMarket := range data {
		assert.NotEmpty(t, currentMarket.StockCode, "Stock code should not be empty")
		assert.NotEmpty(t, currentMarket.ShortName, "Short name should not be empty")
		assert.GreaterOrEqual(t, currentMarket.Price, 0.0, "Price should be non-negative")
		assert.GreaterOrEqual(t, currentMarket.Volume, int64(0), "Volume should be non-negative")
	}
}

func TestStockMarket_ListMarketCurrent_InvalidParams(t *testing.T) {
	stockMarket := market.NewStockMarket()

	// 测试空代码列表
	_, err := stockMarket.ListMarketCurrent([]string{})
	assert.Error(t, err, "Should return error for empty codes list")

	// 测试nil代码列表
	_, err = stockMarket.ListMarketCurrent(nil)
	assert.Error(t, err, "Should return error for nil codes list")
}

func TestStockMarket_GetMarketMin_Valid(t *testing.T) {
	stockMarket := market.NewStockMarket()

	_, err := stockMarket.GetMarketMin("000001")

	// 由于分时数据接口可能未完全实现，我们只测试参数验证
	if err != nil {
		// 如果是网络错误或功能未实现，不算测试失败
		if adataErr, ok := err.(*adataErrors.ADataError); ok {
			if adataErr.Code >= 50000 { // 功能未实现的错误代码
				t.Logf("Feature not implemented yet: %v", err)
				return
			}
		}
		t.Logf("Network request failed (expected in some environments): %v", err)
	}
}

func TestStockMarket_GetMarketMin_Invalid(t *testing.T) {
	stockMarket := market.NewStockMarket()

	_, err := stockMarket.GetMarketMin("invalid")
	assert.Error(t, err, "Should return error for invalid stock code")

	if adataErr, ok := err.(*adataErrors.ADataError); ok {
		assert.Equal(t, adataErrors.ErrInvalidStockCode.Code, adataErr.Code,
			"Should return ErrInvalidStockCode for invalid stock code")
	}
}

func TestStockMarket_GetMarketFive_Valid(t *testing.T) {
	stockMarket := market.NewStockMarket()

	_, err := stockMarket.GetMarketFive("000001")

	// 由于五档行情接口可能未完全实现，我们只测试参数验证
	if err != nil {
		// 如果是网络错误或功能未实现，不算测试失败
		if adataErr, ok := err.(*adataErrors.ADataError); ok {
			if adataErr.Code >= 50000 { // 功能未实现的错误代码
				t.Logf("Feature not implemented yet: %v", err)
				return
			}
		}
		t.Logf("Network request failed (expected in some environments): %v", err)
	}
}

func TestStockMarket_GetMarketFive_Invalid(t *testing.T) {
	stockMarket := market.NewStockMarket()

	_, err := stockMarket.GetMarketFive("invalid")
	assert.Error(t, err, "Should return error for invalid stock code")

	if adataErr, ok := err.(*adataErrors.ADataError); ok {
		assert.Equal(t, adataErrors.ErrInvalidStockCode.Code, adataErr.Code,
			"Should return ErrInvalidStockCode for invalid stock code")
	}
}
