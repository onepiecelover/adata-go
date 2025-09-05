// Package adata provides A股量化数据SDK for Go
package adata

import (
	"github.com/onepiecelover/adata-go/pkg/bond"
	"github.com/onepiecelover/adata-go/pkg/fund"
	"github.com/onepiecelover/adata-go/pkg/sentiment"
	"github.com/onepiecelover/adata-go/pkg/stock"
)

const (
	// Version 当前版本号
	Version = "1.0.0"

	// Title 项目标题
	Title = "adata-go"

	// Description 项目描述
	Description = "A Data,A Stock,ETF,Bond,Quant,Stock Market,K Line for Go"

	// URL 项目地址
	URL = "https://github.com/onepiecelover/adata-go"

	// Author 作者
	Author = "1nchaos"

	// AuthorEmail 作者邮箱
	AuthorEmail = "9527@1nchaos.com"

	// License 许可证
	License = "Apache License 2.0"
)

var (
	// Stock 股票模块
	Stock = stock.New()

	// Fund 基金模块
	Fund = fund.New()

	// Bond 债券模块
	Bond = bond.New()

	// Sentiment 情感指标模块
	Sentiment = sentiment.New()
)

// GetVersion 获取版本信息
func GetVersion() string {
	return Version
}

// SetProxy 设置代理
func SetProxy(enable bool, proxyURL string) {
	// 设置全局代理配置
	stock.SetProxy(enable, proxyURL)
	fund.SetProxy(enable, proxyURL)
	bond.SetProxy(enable, proxyURL)
	sentiment.SetProxy(enable, proxyURL)
}
