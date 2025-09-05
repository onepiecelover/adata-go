package main

import (
	"fmt"
	"log"

	"github.com/onepiecelover/adata-go"
)

// 股票信息查询示例
func main() {
	fmt.Println("=== 股票信息查询示例 ===")

	// 1. 获取所有股票代码
	fmt.Println("1. 获取所有股票代码...")
	allCodes, err := adata.Stock.Info.AllCode()
	if err != nil {
		log.Fatalf("获取股票代码失败: %v", err)
	}

	fmt.Printf("总共获取到 %d 只股票\n\n", len(allCodes))

	// 2. 按交易所分类统计
	fmt.Println("2. 按交易所分类统计:")
	exchangeCount := make(map[string]int)
	for _, code := range allCodes {
		exchangeCount[code.Exchange]++
	}

	for exchange, count := range exchangeCount {
		fmt.Printf("  %s: %d 只\n", exchange, count)
	}
	fmt.Println()

	// 3. 显示部分股票信息
	fmt.Println("3. 部分股票信息:")
	for i, code := range allCodes {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s - %s (%s) 上市日期: %s\n",
			code.StockCode, code.ShortName, code.Exchange, code.ListDate)
	}
	fmt.Println()

	// 4. 获取指数代码
	fmt.Println("4. 获取指数代码...")
	indexCodes, err := adata.Stock.Info.AllIndexCode()
	if err != nil {
		log.Printf("获取指数代码失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个指数\n", len(indexCodes))
		for i, index := range indexCodes {
			if i >= 5 {
				break
			}
			fmt.Printf("  %s - %s (%s)\n", index.IndexCode, index.IndexName, index.Exchange)
		}
	}
	fmt.Println()

	// 5. 获取东方财富概念代码
	fmt.Println("5. 获取东方财富概念代码...")
	conceptCodes, err := adata.Stock.Info.AllConceptCodeEast()
	if err != nil {
		log.Printf("获取概念代码失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个概念\n", len(conceptCodes))
		for i, concept := range conceptCodes {
			if i >= 5 {
				break
			}
			fmt.Printf("  %s - %s\n", concept.ConceptCode, concept.ConceptName)
		}
	}
	fmt.Println()

	// 6. 查询特定股票的概念信息
	fmt.Println("6. 查询平安银行(000001)的概念信息...")
	concepts, err := adata.Stock.Info.GetConceptEast("000001")
	if err != nil {
		log.Printf("获取概念信息失败: %v", err)
	} else {
		fmt.Printf("平安银行所属 %d 个概念:\n", len(concepts))
		for _, concept := range concepts {
			fmt.Printf("  %s - %s\n", concept.ConceptCode, concept.ConceptName)
		}
	}

	fmt.Println("\n=== 示例结束 ===")
}
