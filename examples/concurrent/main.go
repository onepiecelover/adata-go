package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/onepiecelover/adata-go"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// 并发获取多只股票的行情数据示例
func main() {
	fmt.Println("=== 并发获取股票行情示例 ===")

	// 股票代码列表
	stockCodes := []string{"000001", "000002", "600036", "600519", "000858", "002594"}

	// 使用通道收集结果
	resultChan := make(chan MarketResult, len(stockCodes))
	var wg sync.WaitGroup

	// 并发获取每只股票的行情数据
	for _, code := range stockCodes {
		wg.Add(1)
		go func(stockCode string) {
			defer wg.Done()

			startDate, _ := time.Parse("2006-01-02", "2024-01-01")
			endDate, _ := time.Parse("2006-01-02", "2024-01-10")
			params := &types.MarketParams{
				StockCode:  stockCode,
				StartDate:  startDate,
				EndDate:    endDate,
				KType:      1, // 日K线
				AdjustType: 1, // 前复权
			}

			data, err := adata.Stock.Market.GetMarket(params)
			resultChan <- MarketResult{
				StockCode: stockCode,
				Data:      data,
				Error:     err,
			}
		}(code)
	}

	// 关闭结果通道
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集并显示结果
	fmt.Println("并发获取结果:")
	for result := range resultChan {
		if result.Error != nil {
			log.Printf("股票 %s 获取失败: %v", result.StockCode, result.Error)
		} else {
			fmt.Printf("股票 %s: 获取到 %d 条数据\n", result.StockCode, len(result.Data))
			if len(result.Data) > 0 {
				latest := result.Data[len(result.Data)-1]
				fmt.Printf("  最新数据 (%s): 开盘=%.2f, 收盘=%.2f, 涨跌幅=%.2f%%\n",
					latest.TradeDate, latest.Open, latest.Close, latest.ChangePct)
			}
		}
	}

	fmt.Println("\n=== 批量获取当前行情 ===")

	// 批量获取当前行情
	currentMarkets, err := adata.Stock.Market.ListMarketCurrent(stockCodes)
	if err != nil {
		log.Printf("批量获取当前行情失败: %v", err)
	} else {
		fmt.Printf("成功获取 %d 只股票的当前行情:\n", len(currentMarkets))
		for _, market := range currentMarkets {
			fmt.Printf("  %s(%s): %.2f (%.2f%%)\n",
				market.StockCode, market.ShortName, market.Price, market.ChangePct)
		}
	}

	fmt.Println("\n=== 示例结束 ===")
}

// MarketResult 行情数据结果
type MarketResult struct {
	StockCode string
	Data      []types.MarketData
	Error     error
}
