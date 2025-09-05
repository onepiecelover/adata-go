// Package headers 定义各数据源的请求头
package headers

import (
	"math/rand"
	"time"
)

// UserAgents 用户代理列表
var UserAgents = []string{
	"Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; SM-G981B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Mobile Safari/537.36",
	"Mozilla/5.0 (iPad; CPU OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Mobile Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
}

// GetRandomUserAgent 获取随机用户代理
func GetRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return UserAgents[rand.Intn(len(UserAgents))]
}

// SinaHeaders 新浪财经请求头
var SinaHeaders = map[string]string{
	"Host":            "hq.sinajs.cn",
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0",
	"Accept":          "*/*",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Accept-Encoding": "gzip, deflate, br",
	"Referer":         "http://vip.stock.finance.sina.com.cn/",
	"Connection":      "keep-alive",
	"Sec-Fetch-Dest":  "script",
	"Sec-Fetch-Mode":  "no-cors",
	"Sec-Fetch-Site":  "cross-site",
}

// EastMoneyHeaders 东方财富请求头
var EastMoneyHeaders = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/110.0",
	"Accept":          "*/*",
	"Content-Type":    "text/plain;charset=UTF-8",
	"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	"Accept-Encoding": "gzip, deflate, br",
	"Connection":      "keep-alive",
	"Referer":         "https://data.eastmoney.com/rzrq/total.html",
}

// BaiduHeaders 百度股市通请求头
func GetBaiduHeaders() map[string]string {
	return map[string]string{
		"Host":            "finance.pae.baidu.com",
		"User-Agent":      GetRandomUserAgent(),
		"Accept":          "application/vnd.finance-web.v1+json",
		"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"Accept-Encoding": "gzip, deflate, br",
		"Content-Type":    "application/json",
		"Origin":          "https://gushitong.baidu.com",
		"Connection":      "keep-alive",
		"Referer":         "https://gushitong.baidu.com/",
	}
}

// THSHeaders 同花顺请求头
func GetTHSHeaders() map[string]string {
	return map[string]string{
		"Host":                      "q.10jqka.com.cn",
		"Referer":                   "http://q.10jqka.com.cn/",
		"User-Agent":                GetRandomUserAgent(),
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language":           "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"Accept-Encoding":           "gzip, deflate",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
	}
}

// GetTencentHeaders 腾讯财经请求头
func GetTencentHeaders() map[string]string {
	return map[string]string{
		"User-Agent":      GetRandomUserAgent(),
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Accept-Encoding": "gzip, deflate",
		"Connection":      "keep-alive",
		"Referer":         "https://stockapp.finance.qq.com/",
	}
}

// GetCommonHeaders 获取通用请求头
func GetCommonHeaders() map[string]string {
	return map[string]string{
		"User-Agent":      GetRandomUserAgent(),
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Cache-Control":   "no-cache",
	}
}
