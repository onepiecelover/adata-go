package main

import (
	"fmt"
	"log"
	"time"

	"github.com/onepiecelover/adata-go/pkg/stock"
	"github.com/onepiecelover/adata-go/pkg/types"
)

func main() {
	fmt.Println("=== AData-Go 示例程序 ===")
	fmt.Printf("版本: %s\n\n", "1.0.0")

	// 创建股票实例
	stockInstance := stock.New()

	// 1. 获取所有股票代码
	fmt.Println("1. 获取股票代码列表...")
	codes, err := stockInstance.Info.AllCode()
	if err != nil {
		log.Printf("获取股票代码失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 只股票\n", len(codes))
		// 显示前5只股票
		for i, code := range codes {
			if i >= 5 {
				break
			}
			fmt.Printf("  %s - %s (%s)\n", code.StockCode, code.ShortName, code.Exchange)
		}
	}

	fmt.Println()

	// 2. 获取股票K线行情数据
	fmt.Println("2. 获取股票K线行情数据...")
	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-01-31")
	params := &types.MarketParams{
		StockCode:  "000001",
		StartDate:  startDate,
		EndDate:    endDate,
		KType:      1, // 日K线
		AdjustType: 1, // 前复权
	}

	marketData, err := stockInstance.Market.GetMarket(params)
	if err != nil {
		log.Printf("获取行情数据失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条行情数据\n", len(marketData))
		// 显示前3条数据
		for i, data := range marketData {
			if i >= 3 {
				break
			}
			fmt.Printf("  %s: 开盘=%.2f, 收盘=%.2f, 最高=%.2f, 最低=%.2f, 成交量=%d\n",
				data.TradeDate, data.Open, data.Close, data.High, data.Low, data.Volume)
		}
	}

	fmt.Println()

	// 3. 获取当前股票行情
	fmt.Println("3. 获取当前股票行情...")
	codes_list := []string{"000001", "000002", "600036", "600519"}
	currentMarket, err := stockInstance.Market.ListMarketCurrent(codes_list)
	if err != nil {
		log.Printf("获取当前行情失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 只股票的当前行情\n", len(currentMarket))
		for _, market := range currentMarket {
			fmt.Printf("  %s(%s): 价格=%.2f, 涨跌=%.2f(%.2f%%), 成交量=%d\n",
				market.StockCode, market.ShortName, market.Price,
				market.Change, market.ChangePct, market.Volume)
		}
	}

	fmt.Println()

	// 4. 获取东方财富概念信息
	fmt.Println("4. 获取股票概念信息...")
	concepts, err := stockInstance.Info.GetConceptEast("000001")
	if err != nil {
		log.Printf("获取概念信息失败: %v", err)
	} else {
		fmt.Printf("000001 所属概念数量: %d\n", len(concepts))
		for i, concept := range concepts {
			if i >= 3 {
				break
			}
			fmt.Printf("  %s - %s\n", concept.ConceptCode, concept.ConceptName)
		}
	}

	fmt.Println()

	// 5. 设置代理示例（如果需要）
	fmt.Println("5. 代理设置示例...")
	// adata.SetProxy(true, "http://proxy-server:8080")
	fmt.Println("  代理设置已跳过（示例中未实际启用）")

	fmt.Println()

	// 6. 展示错误处理
	fmt.Println("6. 错误处理示例...")
	startDate, _ = time.Parse("2006-01-02", "2024-01-01")
	_, err = stockInstance.Market.GetMarket(&types.MarketParams{
		StockCode: "invalid_code", // 无效的股票代码
		StartDate: startDate,
	})
	if err != nil {
		fmt.Printf("  预期错误: %v\n", err)
	}

	fmt.Println("\n=== 示例程序结束 ===")
}
