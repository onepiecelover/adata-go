// Package market 提供股票行情相关功能
package market

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/common/headers"
	"github.com/onepiecelover/adata-go/pkg/common/utils"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// StockMarket 股票行情结构体
type StockMarket struct {
	client *client.Client
}

// NewStockMarket 创建股票行情实例
func NewStockMarket() *StockMarket {
	return &StockMarket{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func (s *StockMarket) SetProxy(enabled bool, proxyURL string) {
	s.client.SetProxy(enabled, proxyURL)
}

// GetMarket 获取股票K线行情数据
func (s *StockMarket) GetMarket(params *types.MarketParams) ([]types.MarketData, error) {
	if params == nil {
		return nil, errors.NewADataError(errors.ErrInvalidStockCode.Code, "参数不能为空", "")
	}

	if !utils.IsValidStockCode(params.StockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 优先使用东方财富数据源
	data, err := s.getMarketFromEast(params)
	if err == nil && len(data) > 0 {
		return data, nil
	}

	// 备用百度数据源
	data, err = s.getMarketFromBaidu(params)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getMarketFromEast 从东方财富获取K线数据
func (s *StockMarket) getMarketFromEast(params *types.MarketParams) ([]types.MarketData, error) {
	baseURL := "http://push2his.eastmoney.com/api/qt/stock/kline/get"

	// 参数处理
	secID := "0"
	if strings.HasPrefix(params.StockCode, "6") {
		secID = "1"
	}

	startDate := "19900101"
	if !params.StartDate.IsZero() {
		startDate = params.StartDate.Format("20060102")
	}

	endDate := utils.GetCurrentDateForAPI()
	if !params.EndDate.IsZero() {
		endDate = params.EndDate.Format("20060102")
	}

	kType := strconv.Itoa(params.KType)
	if params.KType < 5 {
		kType = "10" + kType
	}

	queryParams := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116",
		"ut":      "7eea3edcaed734bea9cbfc24409ed989",
		"klt":     kType,
		"fqt":     strconv.Itoa(params.AdjustType),
		"secid":   fmt.Sprintf("%s.%s", secID, params.StockCode),
		"beg":     startDate,
		"end":     endDate,
		"_":       strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	var result struct {
		Data struct {
			Code   string   `json:"code"`
			Market int      `json:"market"`
			Name   string   `json:"name"`
			Klines []string `json:"klines"`
		} `json:"data"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Data.Klines) == 0 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到行情数据", "")
	}

	var marketData []types.MarketData
	for _, kline := range result.Data.Klines {
		data, err := s.parseEastKlineData(kline, params.StockCode)
		if err != nil {
			continue
		}
		marketData = append(marketData, *data)
	}

	return marketData, nil
}

// parseEastKlineData 解析东方财富K线数据
func (s *StockMarket) parseEastKlineData(kline, stockCode string) (*types.MarketData, error) {
	parts := strings.Split(kline, ",")
	if len(parts) < 11 {
		return nil, fmt.Errorf("invalid kline data format")
	}

	return &types.MarketData{
		TradeDate: parts[0],
		Open:      utils.ParseFloat(parts[1]),
		Close:     utils.ParseFloat(parts[2]),
		High:      utils.ParseFloat(parts[3]),
		Low:       utils.ParseFloat(parts[4]),
		Volume:    utils.ParseInt(parts[5]),
		Amount:    utils.ParseFloat(parts[6]),
		Change:    utils.ParseFloat(parts[7]),
		ChangePct: utils.ParseFloat(parts[8]),
		Turnover:  utils.ParseFloat(parts[9]),
		PreClose:  utils.ParseFloat(parts[10]), // 修正为正确的字段索引
		StockCode: stockCode,
	}, nil
}

// getMarketFromBaidu 从百度获取K线数据
func (s *StockMarket) getMarketFromBaidu(params *types.MarketParams) ([]types.MarketData, error) {
	baseURL := "https://finance.pae.baidu.com/selfselect/getstockquotation"

	startTime := ""
	if !params.StartDate.IsZero() {
		startTime = params.StartDate.Format("2006-01-02")
	}

	queryParams := map[string]string{
		"all":           "1",
		"isIndex":       "false",
		"isBk":          "false",
		"isBlock":       "false",
		"isFutures":     "false",
		"isStock":       "true",
		"newFormat":     "1",
		"group":         "quotation_kline_ab",
		"finClientType": "pc",
		"code":          params.StockCode,
		"start_time":    startTime,
		"ktype":         strconv.Itoa(params.KType),
	}

	var result struct {
		ResultCode string `json:"ResultCode"`
		Result     struct {
			NewMarketData struct {
				Keys       []string `json:"keys"`
				MarketData string   `json:"marketData"`
			} `json:"newMarketData"`
		} `json:"Result"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.GetBaiduHeaders(), &result)
	if err != nil {
		return nil, err
	}

	if result.ResultCode != "0" || result.Result.NewMarketData.MarketData == "" {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到行情数据", "")
	}

	return s.parseBaiduMarketData(result.Result.NewMarketData.MarketData, params.StockCode)
}

// parseBaiduMarketData 解析百度行情数据
func (s *StockMarket) parseBaiduMarketData(marketData, stockCode string) ([]types.MarketData, error) {
	lines := strings.Split(marketData, ";")
	var result []types.MarketData

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 11 {
			continue
		}

		// 时间戳转换
		timestamp, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}

		tradeTime := time.Unix(timestamp, 0)

		data := types.MarketData{
			TradeDate: tradeTime.Format("2006-01-02"),
			Open:      utils.ParseFloat(parts[1]),
			Close:     utils.ParseFloat(parts[2]),
			High:      utils.ParseFloat(parts[3]),
			Low:       utils.ParseFloat(parts[4]),
			Volume:    utils.ParseInt(parts[5]),
			Amount:    utils.ParseFloat(parts[6]),
			Change:    utils.ParseFloat(parts[7]),
			ChangePct: utils.ParseFloat(parts[8]),
			Turnover:  utils.ParseFloat(parts[9]),
			PreClose:  utils.ParseFloat(parts[10]),
			StockCode: stockCode,
		}

		result = append(result, data)
	}

	return result, nil
}

// GetMarketMin 获取股票当日分时行情
func (s *StockMarket) GetMarketMin(stockCode string) ([]types.MarketMin, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 优先使用东方财富数据源
	data, err := s.getMarketMinFromEast(stockCode)
	if err == nil && len(data) > 0 {
		return data, nil
	}

	// 备用百度数据源
	data, err = s.getMarketMinFromBaidu(stockCode)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getMarketMinFromEast 从东方财富获取分时数据
func (s *StockMarket) getMarketMinFromEast(stockCode string) ([]types.MarketMin, error) {
	baseURL := "https://push2.eastmoney.com/api/qt/stock/trends2/get"

	// 参数处理
	secID := "0"
	if strings.HasPrefix(stockCode, "6") {
		secID = "1"
	}

	queryParams := map[string]string{
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58",
		"ut":      "fa5fd1943c7b386f172d6893dbfba10b",
		"ndays":   "1",
		"iscr":    "1",
		"iscca":   "0",
		"secid":   fmt.Sprintf("%s.%s", secID, stockCode),
		"_":       strconv.FormatInt(time.Now().UnixMilli(), 10),
	}

	var result struct {
		Data struct {
			PreClose float64  `json:"preClose"`
			Trends   []string `json:"trends"`
		} `json:"data"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if result.Data.Trends == nil || len(result.Data.Trends) == 0 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到分时数据", "")
	}

	var marketMinData []types.MarketMin
	for _, trend := range result.Data.Trends {
		data, err := s.parseEastTrendData(trend, stockCode, result.Data.PreClose)
		if err != nil {
			continue
		}
		marketMinData = append(marketMinData, *data)
	}

	return marketMinData, nil
}

// parseEastTrendData 解析东方财富分时数据
func (s *StockMarket) parseEastTrendData(trend, stockCode string, preClose float64) (*types.MarketMin, error) {
	parts := strings.Split(trend, ",")
	if len(parts) < 8 {
		return nil, fmt.Errorf("invalid trend data format")
	}

	tradeTime, err := time.Parse("2006-01-02 15:04", parts[0])
	if err != nil {
		return nil, err
	}

	volume := utils.ParseInt(parts[5]) * 100 // 换算成股
	amount := utils.ParseFloat(parts[6])
	price := utils.ParseFloat(parts[2])
	avgPrice := price

	change := price - preClose
	changePct := 0.0
	if preClose != 0 {
		changePct = change / preClose * 100
	}

	return &types.MarketMin{
		StockCode: stockCode,
		TradeTime: tradeTime,
		Price:     price,
		Change:    change,
		ChangePct: changePct,
		AvgPrice:  avgPrice,
		Volume:    volume,
		Amount:    amount,
	}, nil
}

// getMarketMinFromBaidu 从百度获取分时数据
func (s *StockMarket) getMarketMinFromBaidu(stockCode string) ([]types.MarketMin, error) {
	baseURL := "https://finance.pae.baidu.com/selfselect/getstockquotation"

	queryParams := map[string]string{
		"all":           "1",
		"code":          stockCode,
		"isIndex":       "false",
		"isBk":          "false",
		"isBlock":       "false",
		"isFutures":     "false",
		"isStock":       "true",
		"newFormat":     "1",
		"group":         "quotation_minute_ab",
		"finClientType": "pc",
	}

	var result struct {
		ResultCode string `json:"ResultCode"`
		Result     struct {
			// 百度分时数据结构
		} `json:"Result"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.GetBaiduHeaders(), &result)
	if err != nil {
		return nil, err
	}

	if result.ResultCode != "0" {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到分时数据", "")
	}

	// 解析分时数据
	return nil, errors.NewADataError(50102, "百度分时数据解析待实现", "")
}

// ListMarketCurrent 获取多个股票的当前行情
func (s *StockMarket) ListMarketCurrent(codes []string) ([]types.CurrentMarket, error) {
	if len(codes) == 0 {
		return nil, errors.NewADataError(errors.ErrInvalidStockCode.Code, "股票代码列表不能为空", "")
	}

	// 优先使用新浪数据源
	data, err := s.getCurrentMarketFromSina(codes)
	if err == nil && len(data) > 0 {
		return data, nil
	}

	// 备用腾讯数据源
	data, err = s.getCurrentMarketFromTencent(codes)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getCurrentMarketFromSina 从新浪获取当前行情
func (s *StockMarket) getCurrentMarketFromSina(codes []string) ([]types.CurrentMarket, error) {
	baseURL := "https://hq.sinajs.cn/list="

	// 构建请求URL
	var urlCodes []string
	for _, code := range codes {
		if !utils.IsValidStockCode(code) {
			continue
		}

		prefix := code[:1]
		switch prefix {
		case "0", "3":
			urlCodes = append(urlCodes, "s_sz"+code)
		case "6", "9":
			urlCodes = append(urlCodes, "s_sh"+code)
		case "4", "8":
			urlCodes = append(urlCodes, "s_bj"+code)
		}
	}

	if len(urlCodes) == 0 {
		return nil, errors.NewADataError(errors.ErrInvalidStockCode.Code, "没有有效的股票代码", "")
	}

	url := baseURL + strings.Join(urlCodes, ",")

	response, err := s.client.GetText(url, nil, headers.SinaHeaders)
	if err != nil {
		return nil, err
	}

	return s.parseSinaCurrentMarket(response)
}

// parseSinaCurrentMarket 解析新浪当前行情数据
func (s *StockMarket) parseSinaCurrentMarket(response string) ([]types.CurrentMarket, error) {
	lines := strings.Split(response, "\n")
	var results []types.CurrentMarket

	for _, line := range lines {
		if line == "" || !strings.Contains(line, "=") {
			continue
		}

		// 解析格式: var hq_str_s_sz000001="平安银行,000001,11.50,0.10,0.88,227841,261792";
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}

		dataStr := strings.Trim(parts[1], `";`)
		dataParts := strings.Split(dataStr, ",")

		if len(dataParts) < 7 {
			continue
		}

		result := types.CurrentMarket{
			ShortName: utils.CleanString(dataParts[0]),
			StockCode: utils.FormatStockCode(dataParts[1]),
			Price:     utils.ParseFloat(dataParts[2]),
			Change:    utils.ParseFloat(dataParts[3]),
			ChangePct: utils.ParseFloat(dataParts[4]),
			Volume:    utils.ParseInt(dataParts[5]) * 100,     // 新浪返回的是手，需要转换为股
			Amount:    utils.ParseFloat(dataParts[6]) * 10000, // 新浪返回的是万元，需要转换为元
		}

		results = append(results, result)
	}

	return results, nil
}

// getCurrentMarketFromTencent 从腾讯获取当前行情
func (s *StockMarket) getCurrentMarketFromTencent(codes []string) ([]types.CurrentMarket, error) {
	baseURL := "https://qt.gtimg.cn/r=0.5979076524724433&q="

	// 构建请求URL
	var urlCodes []string
	for _, code := range codes {
		if !utils.IsValidStockCode(code) {
			continue
		}

		prefix := code[:1]
		switch prefix {
		case "0", "3":
			urlCodes = append(urlCodes, "s_sz"+code)
		case "6", "9":
			urlCodes = append(urlCodes, "s_sh"+code)
		case "4", "8":
			urlCodes = append(urlCodes, "s_bj"+code)
		}
	}

	if len(urlCodes) == 0 {
		return nil, errors.NewADataError(errors.ErrInvalidStockCode.Code, "没有有效的股票代码", "")
	}

	url := baseURL + strings.Join(urlCodes, ",")

	response, err := s.client.GetText(url, nil, headers.GetTencentHeaders())
	if err != nil {
		return nil, err
	}

	return s.parseTencentCurrentMarket(response)
}

// parseTencentCurrentMarket 解析腾讯当前行情数据
func (s *StockMarket) parseTencentCurrentMarket(response string) ([]types.CurrentMarket, error) {
	// 解析腾讯行情数据格式
	// v_s_sz000936="51~华西股份~000936~12.60~1.15~10.04~69137~8711~~111.64~GP-A";
	lines := strings.Split(response, ";")
	var results []types.CurrentMarket

	for _, line := range lines {
		if len(line) < 8 {
			continue
		}

		parts := strings.Split(line, "~")
		if len(parts) < 11 {
			continue
		}

		result := types.CurrentMarket{
			ShortName: utils.CleanString(parts[1]),
			StockCode: utils.FormatStockCode(parts[2]),
			Price:     utils.ParseFloat(parts[3]),
			Change:    utils.ParseFloat(parts[4]),
			ChangePct: utils.ParseFloat(parts[5]),
			Volume:    utils.ParseInt(parts[6]) * 100,     // 腾讯返回的是手，需要转换为股
			Amount:    utils.ParseFloat(parts[7]) * 10000, // 腾讯返回的是万元，需要转换为元
		}

		results = append(results, result)
	}

	return results, nil
}

// GetMarketFive 获取股票五档行情
func (s *StockMarket) GetMarketFive(stockCode string) (*types.MarketFive, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 优先使用腾讯数据源
	data, err := s.getFiveMarketFromTencent(stockCode)
	if err == nil {
		return data, nil
	}

	// 备用百度数据源
	data, err = s.getFiveMarketFromBaidu(stockCode)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getFiveMarketFromTencent 从腾讯获取五档行情
func (s *StockMarket) getFiveMarketFromTencent(stockCode string) (*types.MarketFive, error) {
	baseURL := "https://web.sqt.gtimg.cn/q="

	// 构建请求URL
	var urlCode string
	prefix := stockCode[:1]
	switch prefix {
	case "0", "3":
		urlCode = "sz" + stockCode
	case "6", "9":
		urlCode = "sh" + stockCode
	case "4", "8":
		urlCode = "bj" + stockCode
	}

	url := baseURL + urlCode

	response, err := s.client.GetText(url, nil, nil)
	if err != nil {
		return nil, err
	}

	return s.parseTencentFiveMarket(response, stockCode)
}

// parseTencentFiveMarket 解析腾讯五档行情数据
func (s *StockMarket) parseTencentFiveMarket(response, stockCode string) (*types.MarketFive, error) {
	lines := strings.Split(response, ";")
	if len(lines) == 0 || len(lines[0]) < 8 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到五档行情数据", "")
	}

	dataStr := lines[0]
	parts := strings.Split(dataStr, "~")
	if len(parts) < 85 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "五档行情数据格式不正确", "")
	}

	// 解析五档行情数据
	marketFive := &types.MarketFive{
		StockCode: stockCode,
		Price:     utils.ParseFloat(parts[3]),
		Change:    utils.ParseFloat(parts[31]),
		ChangePct: utils.ParseFloat(parts[32]),
		Volume:    utils.ParseInt(parts[6]) * 100,     // 腾讯返回的是手，需要转换为股
		Amount:    utils.ParseFloat(parts[7]) * 10000, // 腾讯返回的是万元，需要转换为元
	}

	// 解析五档买盘
	for i := 0; i < 5; i++ {
		marketFive.BuyPrices[i] = utils.ParseFloat(parts[9+i*2])
		marketFive.BuyVolumes[i] = utils.ParseInt(parts[10+i*2]) * 100 // 转换为股
	}

	// 解析五档卖盘
	for i := 0; i < 5; i++ {
		marketFive.SellPrices[i] = utils.ParseFloat(parts[19+i*2])
		marketFive.SellVolumes[i] = utils.ParseInt(parts[20+i*2]) * 100 // 转换为股
	}

	return marketFive, nil
}

// getFiveMarketFromBaidu 从百度获取五档行情
func (s *StockMarket) getFiveMarketFromBaidu(stockCode string) (*types.MarketFive, error) {
	// 百度五档行情接口实现
	return nil, errors.NewADataError(50202, "百度五档行情获取待实现", "")
}

// GetCapitalFlowMin 获取股票当日分时资金流向
func (s *StockMarket) GetCapitalFlowMin(stockCode string) ([]types.CapitalFlow, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 使用东方财富数据源
	return s.getCapitalFlowMinFromEast(stockCode)
}

// getCapitalFlowMinFromEast 从东方财富获取分时资金流向
func (s *StockMarket) getCapitalFlowMinFromEast(stockCode string) ([]types.CapitalFlow, error) {
	baseURL := "https://push2.eastmoney.com/api/qt/stock/fflow/kline/get"

	// 参数处理
	secID := "0"
	if strings.HasPrefix(stockCode, "6") {
		secID = "1"
	}

	queryParams := map[string]string{
		"lmt":     "0",
		"klt":     "1",
		"fields1": "f1,f2,f3,f7",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65",
		"secid":   fmt.Sprintf("%s.%s", secID, stockCode),
	}

	var result struct {
		Data struct {
			Klines []string `json:"klines"`
		} `json:"data"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if result.Data.Klines == nil || len(result.Data.Klines) == 0 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到分时资金流向数据", "")
	}

	var capitalFlows []types.CapitalFlow
	for _, kline := range result.Data.Klines {
		flow, err := s.parseEastCapitalFlowMinData(kline, stockCode)
		if err != nil {
			continue
		}
		capitalFlows = append(capitalFlows, *flow)
	}

	return capitalFlows, nil
}

// parseEastCapitalFlowMinData 解析东方财富分时资金流向数据
func (s *StockMarket) parseEastCapitalFlowMinData(kline, stockCode string) (*types.CapitalFlow, error) {
	parts := strings.Split(kline, ",")
	if len(parts) < 15 {
		return nil, fmt.Errorf("invalid capital flow min data format")
	}

	tradeTime, err := time.Parse("2006-01-02 15:04", parts[0])
	if err != nil {
		return nil, err
	}

	return &types.CapitalFlow{
		StockCode:    stockCode,
		TradeDate:    tradeTime,
		MainInflow:   utils.ParseFloat(parts[1]),
		SuperInflow:  utils.ParseFloat(parts[2]),
		LargeInflow:  utils.ParseFloat(parts[3]),
		MediumInflow: utils.ParseFloat(parts[4]),
		SmallInflow:  utils.ParseFloat(parts[5]),
	}, nil
}

// GetCapitalFlow 获取股票历史资金流向
func (s *StockMarket) GetCapitalFlow(stockCode, startDate, endDate string) ([]types.CapitalFlow, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 使用东方财富数据源
	return s.getCapitalFlowFromEast(stockCode, startDate, endDate)
}

// getCapitalFlowFromEast 从东方财富获取历史资金流向
func (s *StockMarket) getCapitalFlowFromEast(stockCode, startDate, endDate string) ([]types.CapitalFlow, error) {
	baseURL := "https://push2his.eastmoney.com/api/qt/stock/fflow/daykline/get"

	// 参数处理
	secID := "0"
	if strings.HasPrefix(stockCode, "6") {
		secID = "1"
	}

	queryParams := map[string]string{
		"lmt":     "0",
		"klt":     "101",
		"fields1": "f1,f2,f3,f7",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61",
		"secid":   fmt.Sprintf("%s.%s", secID, stockCode),
	}

	var result struct {
		Data struct {
			Klines []string `json:"klines"`
		} `json:"data"`
	}

	err := s.client.GetJSON(baseURL, queryParams, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if result.Data.Klines == nil || len(result.Data.Klines) == 0 {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到历史资金流向数据", "")
	}

	var capitalFlows []types.CapitalFlow
	for _, kline := range result.Data.Klines {
		flow, err := s.parseEastCapitalFlowData(kline, stockCode)
		if err != nil {
			continue
		}

		// 日期范围筛选
		if startDate != "" || endDate != "" {
			startDateParsed, _ := time.Parse("2006-01-02", startDate)
			endDateParsed, _ := time.Parse("2006-01-02", endDate)

			if startDate != "" && flow.TradeDate.Before(startDateParsed) {
				continue
			}

			if endDate != "" && flow.TradeDate.After(endDateParsed) {
				continue
			}
		}

		capitalFlows = append(capitalFlows, *flow)
	}

	return capitalFlows, nil
}

// parseEastCapitalFlowData 解析东方财富历史资金流向数据
func (s *StockMarket) parseEastCapitalFlowData(kline, stockCode string) (*types.CapitalFlow, error) {
	parts := strings.Split(kline, ",")
	if len(parts) < 11 {
		return nil, fmt.Errorf("invalid capital flow data format")
	}

	tradeDate, err := time.Parse("2006-01-02", parts[0])
	if err != nil {
		return nil, err
	}

	return &types.CapitalFlow{
		StockCode:    stockCode,
		TradeDate:    tradeDate,
		MainInflow:   utils.ParseFloat(parts[1]),
		SuperInflow:  utils.ParseFloat(parts[2]),
		LargeInflow:  utils.ParseFloat(parts[3]),
		MediumInflow: utils.ParseFloat(parts[4]),
		SmallInflow:  utils.ParseFloat(parts[5]),
	}, nil
}
