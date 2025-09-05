# 开发指南

## 项目架构

AData-Go 采用模块化架构设计，主要包含以下层次：

```
adata-go/
├── pkg/
│   ├── common/          # 公共模块
│   │   ├── client/      # HTTP客户端
│   │   ├── errors/      # 错误处理
│   │   ├── headers/     # 请求头管理
│   │   └── utils/       # 工具函数
│   ├── stock/           # 股票模块
│   │   ├── info/        # 基础信息
│   │   ├── market/      # 行情数据
│   │   └── finance/     # 财务数据
│   ├── fund/            # 基金模块
│   ├── bond/            # 债券模块
│   ├── sentiment/       # 情感指标
│   └── types/           # 类型定义
├── examples/            # 示例代码
├── tests/              # 测试代码
└── docs/               # 文档
```

## 设计原则

### 1. 模块化设计

每个功能模块独立，职责单一，便于维护和扩展。

### 2. 接口统一

所有模块提供统一的接口设计，保持API的一致性。

### 3. 错误处理

使用自定义错误类型，提供详细的错误信息和错误代码。

### 4. 并发安全

所有公共接口都是线程安全的，支持并发调用。

### 5. 数据源冗余

支持多个数据源，自动切换，提高数据获取的可靠性。

## 核心组件

### HTTP客户端 (common/client)

提供统一的HTTP请求功能：

```go
type Client struct {
    client      *resty.Client
    proxyConfig *ProxyConfig
    retryTimes  int
    waitTime    time.Duration
}
```

**特性：**

- 自动重试机制
- 代理支持
- 超时控制
- 并发安全

### 错误处理 (common/errors)

自定义错误类型：

```go
type ADataError struct {
    Code    int    // 错误代码
    Message string // 错误消息
    Detail  string // 详细信息
}
```

### 数据类型 (types)

定义所有数据结构：

```go
type MarketData struct {
    TradeTime  time.Time
    TradeDate  string
    Open       float64
    // ...其他字段
}
```

## 开发流程

### 1. 环境准备

```bash
# 克隆项目
git clone https://github.com/onepiecelover/adata-go.git
cd adata-go

# 安装依赖
go mod download

# 运行测试
go test ./...
```

### 2. 新增数据源

以新增一个数据源为例：

```go
// 1. 在对应模块中新增数据源文件
// pkg/stock/market/source_new.go

func (s *StockMarket) getMarketFromNewSource(params *types.MarketParams) ([]types.MarketData, error) {
    // 实现数据获取逻辑
    baseURL := "https://api.newsource.com/data"
    
    // 构建请求参数
    queryParams := map[string]string{
        "code": params.StockCode,
        "start": params.StartDate,
        // ...
    }
    
    // 发送请求
    var result NewSourceResponse
    err := s.client.GetJSON(baseURL, queryParams, headers.NewSourceHeaders, &result)
    if err != nil {
        return nil, err
    }
    
    // 解析数据
    return s.parseNewSourceData(result)
}

// 2. 在主方法中集成新数据源
func (s *StockMarket) GetMarket(params *types.MarketParams) ([]types.MarketData, error) {
    // 优先使用原有数据源
    data, err := s.getMarketFromEast(params)
    if err == nil && len(data) > 0 {
        return data, nil
    }
    
    // 新增数据源作为备用
    data, err = s.getMarketFromNewSource(params)
    if err == nil && len(data) > 0 {
        return data, nil
    }
    
    // 其他备用数据源...
}
```

### 3. 新增功能模块

```go
// 1. 定义数据类型 (pkg/types/types.go)
type NewFeatureData struct {
    Field1 string  `json:"field1"`
    Field2 float64 `json:"field2"`
    // ...
}

// 2. 实现功能模块 (pkg/stock/newfeature/newfeature.go)
type NewFeature struct {
    client *client.Client
}

func NewNewFeature() *NewFeature {
    return &NewFeature{
        client: client.NewClient(),
    }
}

func (n *NewFeature) GetData(param string) ([]types.NewFeatureData, error) {
    // 实现具体逻辑
}

// 3. 集成到主模块 (pkg/stock/stock.go)
type Stock struct {
    Info       *info.StockInfo
    Market     *market.StockMarket
    Finance    *finance.StockFinance
    NewFeature *newfeature.NewFeature // 新增
}
```

## 测试指南

### 单元测试

```go
package market

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGetMarket(t *testing.T) {
    market := NewStockMarket()
    
    params := &types.MarketParams{
        StockCode: "000001",
        StartDate: "2024-01-01",
        EndDate:   "2024-01-10",
        KType:     1,
    }
    
    data, err := market.GetMarket(params)
    
    assert.NoError(t, err)
    assert.NotEmpty(t, data)
    assert.Equal(t, "000001", data[0].StockCode)
}
```

### 集成测试

```go
func TestIntegration(t *testing.T) {
    // 测试完整的数据获取流程
    codes, err := adata.Stock.Info.AllCode()
    assert.NoError(t, err)
    assert.NotEmpty(t, codes)
    
    // 使用获取到的股票代码测试行情数据
    if len(codes) > 0 {
        params := &types.MarketParams{
            StockCode: codes[0].StockCode,
            StartDate: "2024-01-01",
            KType:     1,
        }
        
        data, err := adata.Stock.Market.GetMarket(params)
        assert.NoError(t, err)
    }
}
```

### Mock测试

```go
type MockClient struct {
    responses map[string]interface{}
}

func (m *MockClient) GetJSON(url string, params map[string]string, headers map[string]string, result interface{}) error {
    // 模拟HTTP响应
    if mockResp, exists := m.responses[url]; exists {
        // 将mock数据赋值给result
        return nil
    }
    return errors.New("mock data not found")
}
```

## 性能优化

### 1. 并发控制

```go
import "golang.org/x/sync/semaphore"

// 限制并发数量
const MaxConcurrency = 10
sem := semaphore.NewWeighted(MaxConcurrency)

for _, code := range codes {
    sem.Acquire(ctx, 1)
    go func(stockCode string) {
        defer sem.Release(1)
        // 处理逻辑
    }(code)
}
```

### 2. 缓存机制

```go
import "time"

type Cache struct {
    data   map[string]CacheItem
    mu     sync.RWMutex
}

type CacheItem struct {
    Value     interface{}
    ExpiresAt time.Time
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, exists := c.data[key]
    if !exists || time.Now().After(item.ExpiresAt) {
        return nil, false
    }
    
    return item.Value, true
}
```

### 3. 连接池

```go
// HTTP客户端已内置连接池
client := resty.New()
client.SetTimeout(30 * time.Second)
// resty会自动管理连接池
```

## 部署指南

### 1. 编译

```bash
# 编译为可执行文件
go build -o adata-go ./examples/basic

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o adata-go-linux ./examples/basic
GOOS=windows GOARCH=amd64 go build -o adata-go.exe ./examples/basic
```

### 2. Docker部署

```dockerfile
FROM golang:1.18-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./examples/basic

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### 3. 配置管理

```go
// 使用环境变量配置
import "os"

func init() {
    if proxyURL := os.Getenv("ADATA_PROXY_URL"); proxyURL != "" {
        adata.SetProxy(true, proxyURL)
    }
}
```

## 贡献指南

### 1. Fork项目

### 2. 创建特性分支

```bash
git checkout -b feature/new-feature
```

### 3. 提交代码

```bash
git commit -m "Add new feature"
```

### 4. 提交Pull Request

## 常见问题

### Q: 如何处理数据源限制？

A: 设置适当的请求间隔，使用代理，或实现请求队列。

### Q: 如何提高数据获取成功率？

A: 实现多数据源冗余，添加重试机制，监控数据源状态。

### Q: 如何处理大量数据？

A: 使用流式处理，分批获取，实现数据压缩。

### Q: 如何保证数据准确性？

A: 数据校验，多源比对，异常数据过滤。
