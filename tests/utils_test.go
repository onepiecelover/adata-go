package tests

import (
	"testing"

	"github.com/onepiecelover/adata-go/pkg/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangeByStockCode(t *testing.T) {
	tests := []struct {
		stockCode string
		expected  string
	}{
		{"000001", "SZ"}, // 深圳主板
		{"300059", "SZ"}, // 创业板
		{"600036", "SH"}, // 上海主板
		{"688001", "SH"}, // 科创板
		{"430001", "BJ"}, // 北交所
		{"invalid", "UNKNOWN"},
		{"", "UNKNOWN"},
	}

	for _, test := range tests {
		result := utils.GetExchangeByStockCode(test.stockCode)
		assert.Equal(t, test.expected, result, "Failed for stock code: %s", test.stockCode)
	}
}

func TestCompileExchangeByStockCode(t *testing.T) {
	tests := []struct {
		stockCode string
		expected  string
	}{
		{"000001", "000001.SZ"},
		{"300059", "300059.SZ"},
		{"600036", "600036.SH"},
		{"688001", "688001.SH"},
		{"430001", "430001.BJ"},
		{"invalid", "invalid"},
	}

	for _, test := range tests {
		result := utils.CompileExchangeByStockCode(test.stockCode)
		assert.Equal(t, test.expected, result, "Failed for stock code: %s", test.stockCode)
	}
}

func TestIsValidStockCode(t *testing.T) {
	tests := []struct {
		stockCode string
		expected  bool
	}{
		{"000001", true},
		{"300059", true},
		{"600036", true},
		{"688001", true},
		{"430001", true},
		{"12345", false},   // 5位数字
		{"1234567", false}, // 7位数字
		{"abcdef", false},  // 非数字
		{"000abc", false},  // 包含字母
		{"", false},        // 空字符串
	}

	for _, test := range tests {
		result := utils.IsValidStockCode(test.stockCode)
		assert.Equal(t, test.expected, result, "Failed for stock code: %s", test.stockCode)
	}
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"2024-01-01", "2024-01-01", false},
		{"20240101", "2024-01-01", false},
		{"2024/01/01", "2024-01-01", false},
		{"2024.01.01", "2024-01-01", false},
		{"", "", false},
		{"invalid", "", true},
		{"2024-13-01", "", true}, // 无效月份
	}

	for _, test := range tests {
		result, err := utils.FormatDate(test.input)
		if test.hasError {
			assert.Error(t, err, "Expected error for input: %s", test.input)
		} else {
			assert.NoError(t, err, "Unexpected error for input: %s", test.input)
			assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
		}
	}
}

func TestFormatDateForAPI(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"2024-01-01", "20240101", false},
		{"20240101", "20240101", false},
		{"2024/01/01", "20240101", false},
		{"", "", false},
		{"invalid", "", true},
	}

	for _, test := range tests {
		result, err := utils.FormatDateForAPI(test.input)
		if test.hasError {
			assert.Error(t, err, "Expected error for input: %s", test.input)
		} else {
			assert.NoError(t, err, "Unexpected error for input: %s", test.input)
			assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
		}
	}
}

func TestParseFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"123.45", 123.45},
		{"123", 123.0},
		{"0", 0.0},
		{"", 0.0},
		{"-", 0.0},
		{"--", 0.0},
		{"1,234.56", 1234.56}, // 带千分位分隔符
		{"invalid", 0.0},
	}

	for _, test := range tests {
		result := utils.ParseFloat(test.input)
		assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
	}
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"123", 123},
		{"0", 0},
		{"", 0},
		{"-", 0},
		{"--", 0},
		{"1,234", 1234}, // 带千分位分隔符
		{"invalid", 0},
	}

	for _, test := range tests {
		result := utils.ParseInt(test.input)
		assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
	}
}

func TestCleanString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"  hello world  ", "hello world"},
		{"hello    world", "hello world"},
		{"hello\tworld", "hello world"},
		{"hello\nworld", "hello world"},
		{"  ", ""},
		{"", ""},
	}

	for _, test := range tests {
		result := utils.CleanString(test.input)
		assert.Equal(t, test.expected, result, "Failed for input: '%s'", test.input)
	}
}

func TestFormatStockCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "000001"},
		{"123", "000123"},
		{"123456", "123456"},
		{"1234567", "123456"}, // 超过6位会截断
		{"  123  ", "000123"},
		{"sz000001", "000001"},
		{"", ""},
		{"abc", ""},
	}

	for _, test := range tests {
		result := utils.FormatStockCode(test.input)
		assert.Equal(t, test.expected, result, "Failed for input: %s", test.input)
	}
}

func TestConvertUnits(t *testing.T) {
	tests := []struct {
		value    float64
		unit     string
		expected float64
	}{
		{123.45, "万", 1234500},
		{123.45, "亿", 12345000000},
		{123.45, "", 123.45},
		{123.45, "其他", 123.45},
	}

	for _, test := range tests {
		result := utils.ConvertUnits(test.value, test.unit)
		assert.Equal(t, test.expected, result, "Failed for value: %f, unit: %s", test.value, test.unit)
	}
}
