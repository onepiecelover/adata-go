# adata-go

[![Go Report Card](https://goreportcard.com/badge/github.com/onepiecelover/adata-go)](https://goreportcard.com/report/github.com/onepiecelover/adata-go)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/onepiecelover/adata-go)](https://pkg.go.dev/github.com/onepiecelover/adata-go)

A股量化数据SDK for Go，提供全面的A股、ETF、债券等金融数据接口。

## 项目介绍

adata-go 是一个专为Go开发者设计的A股量化数据SDK，提供了丰富的金融数据接口，包括股票基本信息、行情数据、财务数据等。该项目是从Python版本的[adata](https://github.com/onepiecelover/adata)转换而来，保持了原有的功能特性，并充分利用了Go语言的并发特性和静态类型检查优势。

## 功能特性

### 股票信息 (Stock Info)

- 获取所有股票代码
- 获取指数代码
- 获取概念板块代码
- 获取股票所属概念信息
- 获取股票股本信息
- 获取申万行业信息
- 获取交易日历

### 股票行情 (Stock Market)

- K线行情数据（日线、周线、月线）
- 分时行情数据
- 五档行情数据
- 实时行情数据
- 资金流向数据（分时和历史）

### 财务数据 (Stock Finance)

- 核心财务指标
- 资产负债表
- 现金流量表
- 利润表

## 安装

```bash
go get github.com/onepiecelover/adata-go
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/onepiecelover/adata-go"
)

func main() {
    // 获取所有股票代码
    codes, err := adata.Stock.Info.AllCode()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("获取到 %d 只股票\n", len(codes))
    
    // 获取平安银行(000001)的K线数据
    params := &types.MarketParams{
        StockCode: "000001",
        StartDate: "2023-01-01",
        EndDate:   "2023-12-31",
        Period:    "daily",
        Adjust:    "qfq",
    }
    
    klines, err := adata.Stock.Market.GetMarket(params)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, kline := range klines {
        fmt.Printf("日期: %s, 开盘价: %.2f, 收盘价: %.2f\n", 
            kline.TradeDate, kline.Open, kline.Close)
    }
}
```

### 设置代理

```go
// 设置全局代理
adata.SetProxy(true, "http://127.0.0.1:8080")

// 或者单独为某个模块设置代理
adata.Stock.Info.SetProxy(true, "http://127.0.0.1:8080")
adata.Stock.Market.SetProxy(true, "http://127.0.0.1:8080")
adata.Stock.Finance.SetProxy(true, "http://127.0.0.1:8080")
```

## API文档

详细的API文档请参考 [GoDoc](https://pkg.go.dev/github.com/onepiecelover/adata-go)

## 数据源

本项目数据来源于以下公开数据源：

| 数据源 | 描述 |
|--------|------|
| 东方财富 | [数据中心](https://data.eastmoney.com/)，[行情中心](http://quote.eastmoney.com/center/) |
| 百度股市通 | [股市通](https://gushitong.baidu.com/) |
| 腾讯理财 | [行情中心](https://stockapp.finance.qq.com/mstats/#) |
| 新浪财经 | [新浪财经](https://finance.sina.com.cn/stock/) |

## 项目结构

```
adata-go/
├── pkg/
│   ├── stock/
│   │   ├── info/       # 股票信息模块
│   │   ├── market/     # 股票行情模块
│   │   └── finance/    # 财务数据模块
│   ├── fund/           # 基金模块
│   ├── bond/           # 债券模块
│   ├── sentiment/      # 情感指标模块
│   └── types/          # 数据类型定义
├── examples/           # 使用示例
├── tests/              # 单元测试
└── docs/               # 文档
```

## 使用示例

更多使用示例请查看 [examples](examples/) 目录：

- [股票信息示例](examples/stock_info/main.go)
- [股票行情示例](examples/stock_market/main.go)
- [财务数据示例](examples/stock_finance/main.go)
- [并发获取数据示例](examples/concurrent/main.go)

## 测试

运行单元测试：

```bash
go test ./...
```

运行特定模块测试：

```bash
go test ./tests/stock_info_test.go
go test ./tests/stock_market_test.go
go test ./tests/utils_test.go
```

## License

本项目采用 Apache License 2.0 许可证，详情请见 [LICENSE](LICENSE) 文件。

## 致谢

感谢以下数据源提供商：

- [东方财富](https://data.eastmoney.com/)
- [百度股市通](https://gushitong.baidu.com/)
- [腾讯理财](https://stockapp.finance.qq.com/)
- [新浪财经](https://finance.sina.com.cn/)

## 贡献

欢迎提交 Issue 和 Pull Request 来改进本项目。

## 联系方式

- 作者: angle
- 项目地址: <https://github.com/onepiecelover/adata-go>
