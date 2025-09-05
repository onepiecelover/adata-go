// Package utils 提供工具函数
package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ExchangeSuffix 交易所后缀映射
var ExchangeSuffix = map[string]string{
	"00": ".SZ", // 深圳主板
	"20": ".SZ", // 深圳B股
	"30": ".SZ", // 创业板
	"43": ".BJ", // 北交所
	"60": ".SH", // 上海主板
	"68": ".SH", // 科创板
	"83": ".BJ", // 北交所
	"87": ".BJ", // 北交所
	"90": ".SH", // 上海B股
	"92": ".BJ", // 北交所
}

// GetExchangeByStockCode 根据股票代码获取交易所
func GetExchangeByStockCode(stockCode string) string {
	if len(stockCode) < 2 {
		return "UNKNOWN"
	}

	prefix := stockCode[:2]
	if suffix, exists := ExchangeSuffix[prefix]; exists {
		return suffix[1:] // 去掉点号，只返回交易所代码
	}

	return "UNKNOWN"
}

// CompileExchangeByStockCode 根据股票代码补全市场后缀
func CompileExchangeByStockCode(stockCode string) string {
	if len(stockCode) < 2 {
		return stockCode
	}

	prefix := stockCode[:2]
	if suffix, exists := ExchangeSuffix[prefix]; exists {
		return stockCode + suffix
	}

	return stockCode
}

// IsValidStockCode 验证股票代码是否有效
func IsValidStockCode(stockCode string) bool {
	if len(stockCode) != 6 {
		return false
	}

	// 检查是否全为数字
	if _, err := strconv.Atoi(stockCode); err != nil {
		return false
	}

	// 检查前缀是否有效
	prefix := stockCode[:2]
	_, exists := ExchangeSuffix[prefix]
	return exists
}

// FormatDate 格式化日期字符串
func FormatDate(date string) (string, error) {
	if date == "" {
		return "", nil
	}

	// 支持多种日期格式
	formats := []string{
		"2006-01-02",
		"20060102",
		"2006/01/02",
		"2006.01.02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, date); err == nil {
			return t.Format("2006-01-02"), nil
		}
	}

	return "", fmt.Errorf("unsupported date format: %s", date)
}

// FormatDateForAPI 将日期格式化为API所需格式（YYYYMMDD）
func FormatDateForAPI(date string) (string, error) {
	if date == "" {
		return "", nil
	}

	formatted, err := FormatDate(date)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(formatted, "-", ""), nil
}

// GetCurrentDate 获取当前日期字符串
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// GetCurrentDateForAPI 获取当前日期的API格式
func GetCurrentDateForAPI() string {
	return time.Now().Format("20060102")
}

// ParseFloat 安全解析浮点数
func ParseFloat(s string) float64 {
	if s == "" || s == "-" || s == "--" {
		return 0.0
	}

	// 移除可能的千分位分隔符
	s = strings.ReplaceAll(s, ",", "")

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return 0.0
}

// ParseInt 安全解析整数
func ParseInt(s string) int64 {
	if s == "" || s == "-" || s == "--" {
		return 0
	}

	// 移除可能的千分位分隔符
	s = strings.ReplaceAll(s, ",", "")

	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}

	return 0
}

// CleanString 清理字符串中的空白字符
func CleanString(s string) string {
	// 移除前后空白字符
	s = strings.TrimSpace(s)

	// 移除内部多余的空白字符
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	return s
}

// FormatStockCode 格式化股票代码为6位数字
func FormatStockCode(code string) string {
	// 移除可能的前缀和后缀
	code = strings.TrimSpace(code)

	// 提取数字部分
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(code, -1)

	if len(matches) == 0 {
		return ""
	}

	// 取第一个匹配的数字序列
	numStr := matches[0]

	// 补齐到6位
	if len(numStr) < 6 {
		numStr = fmt.Sprintf("%06s", numStr)
	} else if len(numStr) > 6 {
		numStr = numStr[:6]
	}

	return numStr
}

// ConvertUnits 转换单位（万、亿）
func ConvertUnits(value float64, unit string) float64 {
	switch strings.TrimSpace(unit) {
	case "万":
		return value * 10000
	case "亿":
		return value * 100000000
	default:
		return value
	}
}

// IsMarketOpen 检查当前是否为交易时间
func IsMarketOpen() bool {
	now := time.Now()

	// 检查是否为周末
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}

	// 检查时间范围
	hour := now.Hour()
	minute := now.Minute()
	timeInMin := hour*60 + minute

	// 上午 9:30-11:30
	morningStart := 9*60 + 30 // 9:30
	morningEnd := 11*60 + 30  // 11:30

	// 下午 13:00-15:00
	afternoonStart := 13 * 60 // 13:00
	afternoonEnd := 15 * 60   // 15:00

	return (timeInMin >= morningStart && timeInMin <= morningEnd) ||
		(timeInMin >= afternoonStart && timeInMin <= afternoonEnd)
}
