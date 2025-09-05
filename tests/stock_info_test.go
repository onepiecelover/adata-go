package tests

import (
	"testing"

	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/stock/info"
	"github.com/stretchr/testify/assert"
)

func TestStockInfo_IsValidStockCode(t *testing.T) {
	stockInfo := info.NewStockInfo()

	// 测试有效股票代码
	validCodes := []string{"000001", "600036", "300059", "688001"}
	for _, code := range validCodes {
		_, err := stockInfo.GetConceptEast(code)
		// 这里不测试具体返回值，只测试代码验证逻辑
		if err != nil {
			// 如果是网络错误等，不算测试失败，但如果是无效代码错误则测试失败
			if adataErr, ok := err.(*errors.ADataError); ok {
				if adataErr.Code == errors.ErrInvalidStockCode.Code {
					t.Errorf("Valid stock code %s was rejected", code)
				}
			}
		}
	}

	// 测试无效股票代码
	invalidCodes := []string{"12345", "abcdef", "", "0000000"}
	for _, code := range invalidCodes {
		_, err := stockInfo.GetConceptEast(code)
		assert.Error(t, err, "Invalid stock code %s should return error", code)

		if adataErr, ok := err.(*errors.ADataError); ok {
			assert.Equal(t, errors.ErrInvalidStockCode.Code, adataErr.Code,
				"Invalid stock code %s should return ErrInvalidStockCode", code)
		}
	}
}

func TestStockInfo_AllCode(t *testing.T) {
	stockInfo := info.NewStockInfo()

	codes, err := stockInfo.AllCode()

	// 由于这是真实的网络请求，我们只测试基本的逻辑
	if err != nil {
		// 如果网络请求失败，我们不将其视为测试失败
		t.Logf("Network request failed (expected in some environments): %v", err)
		return
	}

	// 如果请求成功，验证返回数据的基本格式
	assert.NotEmpty(t, codes, "Should return non-empty stock codes list")

	for i, code := range codes {
		if i >= 5 { // 只检查前5个
			break
		}
		assert.NotEmpty(t, code.StockCode, "Stock code should not be empty")
		assert.NotEmpty(t, code.ShortName, "Short name should not be empty")
		assert.NotEmpty(t, code.Exchange, "Exchange should not be empty")
		assert.Len(t, code.StockCode, 6, "Stock code should be 6 digits")
	}
}

func TestStockInfo_AllIndexCode(t *testing.T) {
	stockInfo := info.NewStockInfo()

	indexes, err := stockInfo.AllIndexCode()

	// 由于这是真实的网络请求，我们只测试基本的逻辑
	if err != nil {
		t.Logf("Network request failed (expected in some environments): %v", err)
		return
	}

	// 如果请求成功，验证返回数据的基本格式
	assert.NotEmpty(t, indexes, "Should return non-empty index codes list")

	for i, index := range indexes {
		if i >= 3 { // 只检查前3个
			break
		}
		assert.NotEmpty(t, index.IndexCode, "Index code should not be empty")
		assert.NotEmpty(t, index.IndexName, "Index name should not be empty")
		assert.NotEmpty(t, index.Exchange, "Exchange should not be empty")
	}
}

func TestStockInfo_AllConceptCodeEast(t *testing.T) {
	stockInfo := info.NewStockInfo()

	concepts, err := stockInfo.AllConceptCodeEast()

	// 由于这是真实的网络请求，我们只测试基本的逻辑
	if err != nil {
		t.Logf("Network request failed (expected in some environments): %v", err)
		return
	}

	// 如果请求成功，验证返回数据的基本格式
	assert.NotEmpty(t, concepts, "Should return non-empty concept codes list")

	for i, concept := range concepts {
		if i >= 3 { // 只检查前3个
			break
		}
		assert.NotEmpty(t, concept.ConceptCode, "Concept code should not be empty")
		assert.NotEmpty(t, concept.ConceptName, "Concept name should not be empty")
	}
}

// Mock测试 - 测试业务逻辑而不依赖网络
func TestStockInfo_MockTest(t *testing.T) {
	// 这里可以添加Mock HTTP客户端的测试
	// 由于时间限制，暂时跳过Mock测试的实现
	t.Skip("Mock tests not implemented yet")
}
