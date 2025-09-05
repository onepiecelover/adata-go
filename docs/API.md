# AData-Go API 文档

## 简介

AData-Go 是 AData Python 版本的 Go 语言移植版本，提供专业的A股量化数据获取功能。

## 安装

```bash
go get github.com/onepiecelover/adata-go
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/onepiecelover/adata-go"
)

func main() {
    // 获取所有股票代码
    codes, err := adata.Stock.Info.AllCode()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("获取到 %d 只股票\n", len(codes))
}
```

## 模块说明

### 1. 股票模块 (Stock)

#### 1.1 股票信息 (Stock.Info)

##### AllCode() - 获取所有股票代码

```go
codes, err := adata.Stock.Info.AllCode()
if err != nil {
    // 处理错误
}

for _, code := range codes {
    fmt.Printf("%s - %s (%s)\n", code.StockCode, code.ShortName, code.Exchange)
}
```

**返回字段说明：**

- `StockCode`: 股票代码
- `ShortName`: 股票简称  
- `Exchange`: 交易所代码 (SH/SZ/BJ)
- `ListDate`: 上市日期

##### AllIndexCode() - 获取所有指数代码

```go
indexes, err := adata.Stock.Info.AllIndexCode()
```

##### AllConceptCodeEast() - 获取东方财富概念代码

```go
concepts, err := adata.Stock.Info.AllConceptCodeEast()
```

##### GetConceptEast(stockCode) - 获取股票所属概念

```go
concepts, err := adata.Stock.Info.GetConceptEast("000001")
```

#### 1.2 股票行情 (Stock.Market)

##### GetMarket(params) - 获取K线行情数据

```go
params := &types.MarketParams{
    StockCode:  "000001",
    StartDate:  "2024-01-01", 
    EndDate:    "2024-01-31",
    KType:      1, // 1:日K, 2:周K, 3:月K
    AdjustType: 1, // 0:不复权, 1:前复权, 2:后复权
}

data, err := adata.Stock.Market.GetMarket(params)
```

**返回字段说明：**

- `TradeTime`: 交易时间
- `TradeDate`: 交易日期
- `Open`: 开盘价
- `High`: 最高价
- `Low`: 最低价
- `Close`: 收盘价
- `Volume`: 成交量
- `Amount`: 成交额
- `Change`: 涨跌额
- `ChangePct`: 涨跌幅
- `Turnover`: 换手率
- `PreClose`: 昨收价

##### ListMarketCurrent(codes) - 获取当前行情

```go
codes := []string{"000001", "600036", "600519"}
markets, err := adata.Stock.Market.ListMarketCurrent(codes)
```

##### GetMarketMin(stockCode) - 获取分时行情

```go
data, err := adata.Stock.Market.GetMarketMin("000001")
```

##### GetMarketFive(stockCode) - 获取五档行情

```go
five, err := adata.Stock.Market.GetMarketFive("000001")
```

#### 1.3 财务数据 (Stock.Finance)

##### GetCoreIndex(stockCode) - 获取核心财务指标

```go
finance, err := adata.Stock.Finance.GetCoreIndex("000001")
```

### 2. 基金模块 (Fund)

```go
// 获取ETF信息
etfs, err := adata.Fund.AllETFExchangeTradedInfo()

// 获取ETF行情
marketData, err := adata.Fund.GetETFMarket("510300", "2024-01-01", "2024-01-31", 1)
```

### 3. 债券模块 (Bond)

```go
// 获取债券代码
bonds, err := adata.Bond.AllBondCode()

// 获取债券行情
marketData, err := adata.Bond.GetBondMarket("110001", "2024-01-01", "2024-01-31", 1)
```

### 4. 情感指标模块 (Sentiment)

```go
// 获取热门板块
hot, err := adata.Sentiment.GetHotList()

// 获取北向资金
north, err := adata.Sentiment.GetNorthFlow()
```

## 配置

### 设置代理

```go
// 全局设置代理
adata.SetProxy(true, "http://proxy-server:8080")

// 单独模块设置代理  
adata.Stock.Info.SetProxy(true, "http://proxy-server:8080")
```

### 版本信息

```go
version := adata.GetVersion()
fmt.Println("AData-Go 版本:", version)
```

## 错误处理

AData-Go 使用自定义错误类型：

```go
import "github.com/onepiecelover/adata-go/pkg/common/errors"

data, err := adata.Stock.Info.AllCode()
if err != nil {
    if adataErr, ok := err.(*errors.ADataError); ok {
        fmt.Printf("错误代码: %d, 消息: %s\n", adataErr.Code, adataErr.Message)
    }
}
```

**常见错误代码：**

- `10001`: 无效股票代码
- `10002`: 无效日期格式
- `20001`: 请求失败
- `20002`: 解析响应失败
- `30001`: 未找到数据
- `30002`: 数据源不可用

## 数据类型

### MarketParams - 行情查询参数

```go
type MarketParams struct {
    StockCode   string // 股票代码
    StartDate   string // 开始日期 (YYYY-MM-DD)
    EndDate     string // 结束日期 (YYYY-MM-DD)
    KType       int    // K线类型
    AdjustType  int    // 复权类型
}
```

### StockCode - 股票代码信息

```go
type StockCode struct {
    StockCode string // 股票代码
    ShortName string // 股票简称
    Exchange  string // 交易所
    ListDate  string // 上市日期
}
```

### MarketData - 行情数据

```go
type MarketData struct {
    TradeTime  time.Time // 交易时间
    TradeDate  string    // 交易日期
    Open       float64   // 开盘价
    High       float64   // 最高价
    Low        float64   // 最低价
    Close      float64   // 收盘价
    Volume     int64     // 成交量
    Amount     float64   // 成交额
    Change     float64   // 涨跌额
    ChangePct  float64   // 涨跌幅
    Turnover   float64   // 换手率
    PreClose   float64   // 昨收价
    StockCode  string    // 股票代码
}
```

## 并发使用

AData-Go 支持并发调用：

```go
import "sync"

var wg sync.WaitGroup
codes := []string{"000001", "000002", "600036"}

for _, code := range codes {
    wg.Add(1)
    go func(stockCode string) {
        defer wg.Done()
        
        params := &types.MarketParams{
            StockCode: stockCode,
            StartDate: "2024-01-01",
            KType:     1,
        }
        
        data, err := adata.Stock.Market.GetMarket(params)
        // 处理数据...
    }(code)
}

wg.Wait()
```

## 注意事项

1. **请求频率控制**: 建议在请求间添加适当延时，避免被数据源限制
2. **错误重试**: 网络请求可能失败，建议实现重试机制
3. **数据验证**: 使用数据前请验证数据完整性
4. **代理设置**: 如遇访问限制，可设置代理服务器

## 更新日志

### v1.0.0 (2024-01-01)

- 初始版本发布
- 实现基础股票信息和行情数据获取
- 支持多数据源切换
- 提供并发安全的API接口

## 技术支持

- GitHub: <https://github.com/onepiecelover/adata-go>
- 文档: <https://adata.30006124.xyz/>
- 原Python版本: <https://github.com/onepiecelover/adata>

## 许可证

Apache License 2.0
