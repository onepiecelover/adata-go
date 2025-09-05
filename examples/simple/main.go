package main

import (
	"fmt"
	"log"
)

// 简化的示例程序，展示基本概念
func main() {
	fmt.Println("=== AData-Go 示例程序 ===")
	fmt.Println("版本: 1.0.0")
	fmt.Println()

	fmt.Println("🚀 AData-Go 功能特性:")
	fmt.Println("✅ 股票基础信息查询")
	fmt.Println("✅ K线行情数据获取")
	fmt.Println("✅ 实时行情数据查询")
	fmt.Println("✅ 五档行情信息")
	fmt.Println("✅ 概念板块信息")
	fmt.Println("✅ 财务数据获取")
	fmt.Println("✅ 基金ETF信息")
	fmt.Println("✅ 债券数据查询")
	fmt.Println("✅ 市场情感指标")
	fmt.Println()

	fmt.Println("📖 主要模块:")
	fmt.Println("  📊 Stock  - 股票数据模块")
	fmt.Println("    ├── Info     - 基础信息 (代码、概念、指数)")
	fmt.Println("    ├── Market   - 行情数据 (K线、实时、五档)")
	fmt.Println("    └── Finance  - 财务数据 (核心指标、三大报表)")
	fmt.Println("  📈 Fund   - 基金模块 (ETF信息和行情)")
	fmt.Println("  📋 Bond   - 债券模块 (债券信息和行情)")
	fmt.Println("  📝 Sentiment - 情感指标 (热点、资金流向)")
	fmt.Println()

	fmt.Println("🔧 技术特性:")
	fmt.Println("  ⚡ 高性能并发处理")
	fmt.Println("  🔄 多数据源自动切换")
	fmt.Println("  🔒 类型安全保证")
	fmt.Println("  🌐 代理支持")
	fmt.Println("  ⚙️  灵活配置")
	fmt.Println()

	fmt.Println("📝 使用示例:")
	fmt.Println(`
  // 获取股票代码
  codes, err := stock.Info.AllCode()
  
  // 获取K线数据
  data, err := stock.Market.GetMarket(&types.MarketParams{
      StockCode: "000001",
      StartDate: "2024-01-01",
      KType:     1, // 日K线
  })
  
  // 获取实时行情
  current, err := stock.Market.ListMarketCurrent([]string{"000001", "600036"})
	`)

	fmt.Println("🎯 快速开始:")
	fmt.Println("  go get github.com/onepiecelover/adata-go")
	fmt.Println("  import \"github.com/onepiecelover/adata-go\"")
	fmt.Println()

	fmt.Println("📚 更多信息:")
	fmt.Println("  文档: https://adata.30006124.xyz/")
	fmt.Println("  源码: https://github.com/onepiecelover/adata-go")
	fmt.Println("  Python版本: https://github.com/onepiecelover/adata")

	fmt.Println()
	fmt.Println("=== 示例程序结束 ===")

	log.Println("注: 完整功能演示需要网络连接和有效的数据源接口")
}
