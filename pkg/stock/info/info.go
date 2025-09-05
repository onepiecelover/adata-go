// Package info 提供股票基础信息相关功能
package info

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/common/headers"
	"github.com/onepiecelover/adata-go/pkg/common/utils"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// StockInfo 股票信息结构体
type StockInfo struct {
	client *client.Client
}

// NewStockInfo 创建股票信息实例
func NewStockInfo() *StockInfo {
	return &StockInfo{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func (s *StockInfo) SetProxy(enabled bool, proxyURL string) {
	s.client.SetProxy(enabled, proxyURL)
}

// AllCode 获取所有股票代码
func (s *StockInfo) AllCode() ([]types.StockCode, error) {
	// 优先使用百度数据源
	codes, err := s.getAllCodeFromBaidu()
	if err == nil && len(codes) >= 5000 {
		return codes, nil
	}

	// 备用东方财富数据源
	codes, err = s.getAllCodeFromEast()
	if err == nil && len(codes) >= 5000 {
		return codes, nil
	}

	// 备用新浪数据源
	codes, err = s.getAllCodeFromSina()
	if err != nil {
		return nil, err
	}

	return codes, nil
}

// getAllCodeFromBaidu 从百度获取股票代码
func (s *StockInfo) getAllCodeFromBaidu() ([]types.StockCode, error) {
	baseURL := "https://finance.pae.baidu.com/selfselect/getmarketrank"
	params := map[string]string{
		"sort_type":     "1",
		"sort_key":      "14",
		"from_mid":      "1",
		"group":         "pclist",
		"type":          "ab",
		"finClientType": "pc",
	}

	var allCodes []types.StockCode
	maxPageSize := 200

	for pageNo := 0; pageNo < 50; pageNo++ {
		params["pn"] = strconv.Itoa(pageNo * maxPageSize)
		params["rn"] = strconv.Itoa(maxPageSize)

		var result struct {
			ResultCode string `json:"ResultCode"`
			Result     struct {
				Result []struct {
					DisplayData struct {
						ResultData struct {
							TplData struct {
								Result struct {
									Rank []struct {
										Code     string `json:"code"`
										Name     string `json:"name"`
										Exchange string `json:"exchange"`
									} `json:"rank"`
								} `json:"result"`
							} `json:"tplData"`
						} `json:"resultData"`
					} `json:"DisplayData"`
				} `json:"Result"`
			} `json:"Result"`
		}

		err := s.client.GetJSON(baseURL, params, headers.GetBaiduHeaders(), &result)
		if err != nil {
			continue
		}

		if result.ResultCode != "0" || len(result.Result.Result) == 0 {
			break
		}

		rankData := result.Result.Result[0].DisplayData.ResultData.TplData.Result.Rank
		if len(rankData) == 0 {
			break
		}

		for _, item := range rankData {
			code := types.StockCode{
				StockCode: utils.FormatStockCode(item.Code),
				ShortName: utils.CleanString(item.Name),
				Exchange:  item.Exchange,
			}
			allCodes = append(allCodes, code)
		}

		// 添加请求间隔
		time.Sleep(100 * time.Millisecond)
	}

	return allCodes, nil
}

// getAllCodeFromEast 从东方财富获取股票代码
func (s *StockInfo) getAllCodeFromEast() ([]types.StockCode, error) {
	baseURL := "https://82.push2.eastmoney.com/api/qt/clist/get"

	var allCodes []types.StockCode
	currPage := 1
	pageSize := 50

	for currPage < 200 {
		params := map[string]string{
			"pn":     strconv.Itoa(currPage),
			"pz":     strconv.Itoa(pageSize),
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"fid":    "f3",
			"fs":     "m:0 t:6,m:0 t:80,m:1 t:2,m:1 t:23,m:0 t:81 s:2048",
			"fields": "f12,f14",
			"_":      strconv.FormatInt(time.Now().UnixMilli(), 10),
		}

		var result struct {
			Data struct {
				Diff []struct {
					F12 string `json:"f12"` // 股票代码
					F14 string `json:"f14"` // 股票简称
				} `json:"diff"`
			} `json:"data"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			break
		}

		if len(result.Data.Diff) == 0 {
			break
		}

		for _, item := range result.Data.Diff {
			code := types.StockCode{
				StockCode: utils.FormatStockCode(item.F12),
				ShortName: utils.CleanString(item.F14),
				Exchange:  utils.GetExchangeByStockCode(item.F12),
			}
			allCodes = append(allCodes, code)
		}

		if len(result.Data.Diff) < pageSize {
			break
		}

		currPage++
		time.Sleep(100 * time.Millisecond)
	}

	return allCodes, nil
}

// getAllCodeFromSina 从新浪获取股票代码
func (s *StockInfo) getAllCodeFromSina() ([]types.StockCode, error) {
	baseURL := "https://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"

	var allCodes []types.StockCode
	currPage := 1

	for currPage < 200 {
		params := map[string]string{
			"page":   strconv.Itoa(currPage),
			"num":    "80",
			"sort":   "changepercent",
			"asc":    "0",
			"node":   "hs_a",
			"symbol": "",
			"_s_r_a": "page",
		}

		var result []struct {
			Symbol string `json:"symbol"` // 股票代码
			Name   string `json:"name"`   // 股票简称
		}

		err := s.client.GetJSON(baseURL, params, headers.SinaHeaders, &result)
		if err != nil || len(result) == 0 {
			break
		}

		for _, item := range result {
			code := types.StockCode{
				StockCode: utils.FormatStockCode(item.Symbol),
				ShortName: utils.CleanString(item.Name),
				Exchange:  utils.GetExchangeByStockCode(item.Symbol),
			}
			allCodes = append(allCodes, code)
		}

		if len(result) < 80 {
			break
		}

		currPage++
		time.Sleep(100 * time.Millisecond)
	}

	return allCodes, nil
}

// AllConceptCodeTHS 获取同花顺概念代码列表
func (s *StockInfo) AllConceptCodeTHS() ([]types.ConceptCode, error) {
	// 这里需要实现同花顺概念代码获取逻辑
	// 由于同花顺接口较复杂，这里提供基础框架
	return nil, errors.NewADataError(50001, "同花顺概念代码获取功能待实现", "")
}

// AllConceptCodeEast 获取东方财富概念代码列表
func (s *StockInfo) AllConceptCodeEast() ([]types.ConceptCode, error) {
	baseURL := "https://push2.eastmoney.com/api/qt/clist/get"

	var allConcepts []types.ConceptCode
	currPage := 1
	pageSize := 100

	for currPage < 50 {
		params := map[string]string{
			"pn":     strconv.Itoa(currPage),
			"pz":     strconv.Itoa(pageSize),
			"po":     "1",
			"np":     "1",
			"fields": "f12,f13,f14,f62",
			"fid":    "f62",
			"fs":     "m:90+t:3",
		}

		var result struct {
			Data struct {
				Diff []struct {
					F12 string `json:"f12"` // 概念代码
					F14 string `json:"f14"` // 概念名称
				} `json:"diff"`
			} `json:"data"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			break
		}

		if len(result.Data.Diff) == 0 {
			break
		}

		for _, item := range result.Data.Diff {
			concept := types.ConceptCode{
				ConceptCode: item.F12,
				ConceptName: utils.CleanString(item.F14),
			}
			allConcepts = append(allConcepts, concept)
		}

		if len(result.Data.Diff) < pageSize {
			break
		}

		currPage++
		time.Sleep(100 * time.Millisecond)
	}

	return allConcepts, nil
}

// AllIndexCode 获取所有指数代码
func (s *StockInfo) AllIndexCode() ([]types.IndexCode, error) {
	return s.getAllIndexCodeFromEast()
}

// getAllIndexCodeFromEast 从东方财富获取指数代码
func (s *StockInfo) getAllIndexCodeFromEast() ([]types.IndexCode, error) {
	var allIndexes []types.IndexCode

	// 上海指数
	shIndexes, err := s.getIndexCodeByMarket("sh")
	if err == nil {
		allIndexes = append(allIndexes, shIndexes...)
	}

	// 深圳指数
	szIndexes, err := s.getIndexCodeByMarket("sz")
	if err == nil {
		allIndexes = append(allIndexes, szIndexes...)
	}

	return allIndexes, nil
}

// getIndexCodeByMarket 根据市场获取指数代码
func (s *StockInfo) getIndexCodeByMarket(market string) ([]types.IndexCode, error) {
	baseURL := "https://31.push2.eastmoney.com/api/qt/clist/get"

	var fs string
	if market == "sh" {
		fs = "m:1+s:2"
	} else {
		fs = "m:0+t:5"
	}

	var indexes []types.IndexCode
	currPage := 1

	for currPage < 10 {
		params := map[string]string{
			"pn":     strconv.Itoa(currPage),
			"pz":     "20",
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"dect":   "1",
			"wbp2u":  "|0|0|0|web",
			"fid":    "f3",
			"fs":     fs,
			"fields": "f12,f13,f14",
			"_":      strconv.FormatInt(time.Now().UnixMilli(), 10),
		}

		var result struct {
			Data struct {
				Diff []struct {
					F12 string `json:"f12"` // 指数代码
					F13 string `json:"f13"` // 市场代码
					F14 string `json:"f14"` // 指数名称
				} `json:"diff"`
			} `json:"data"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			break
		}

		if result.Data.Diff == nil || len(result.Data.Diff) == 0 {
			break
		}

		for _, item := range result.Data.Diff {
			index := types.IndexCode{
				IndexCode: item.F12,
				IndexName: utils.CleanString(item.F14),
				Exchange:  strings.ToUpper(market),
			}
			indexes = append(indexes, index)
		}

		currPage++
		time.Sleep(100 * time.Millisecond)
	}

	return indexes, nil
}

// GetConceptEast 根据股票代码获取东方财富概念信息
func (s *StockInfo) GetConceptEast(stockCode string) ([]types.ConceptCode, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"
	params := map[string]string{
		"reportName":   "RPT_F10_CORETHEME_BOARDTYPE",
		"columns":      "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,NEW_BOARD_CODE,BOARD_NAME,SELECTED_BOARD_REASON,IS_PRECISE,BOARD_RANK,BOARD_YIELD,DERIVE_BOARD_CODE",
		"quoteColumns": "f3~05~NEW_BOARD_CODE~BOARD_YIELD",
		"filter":       fmt.Sprintf("(SECUCODE=\"%s\")(IS_PRECISE=\"1\")", url.QueryEscape(stockCodeWithExchange)),
		"pageNumber":   "1",
		"pageSize":     "50",
		"sortTypes":    "1",
		"sortColumns":  "BOARD_RANK",
		"source":       "HSF10",
		"client":       "PC",
	}

	var result struct {
		Success bool `json:"success"`
		Result  struct {
			Data []struct {
				NewBoardCode        string `json:"NEW_BOARD_CODE"`
				BoardName           string `json:"BOARD_NAME"`
				SelectedBoardReason string `json:"SELECTED_BOARD_REASON"`
			} `json:"data"`
		} `json:"result"`
	}

	err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到概念数据", "")
	}

	var concepts []types.ConceptCode
	for _, item := range result.Result.Data {
		concept := types.ConceptCode{
			ConceptCode: item.NewBoardCode,
			ConceptName: utils.CleanString(item.BoardName),
		}
		concepts = append(concepts, concept)
	}

	return concepts, nil
}

// GetStockShares 获取股票股本信息
func (s *StockInfo) GetStockShares(stockCode string, isHistory bool) ([]types.StockShares, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	baseURL := "https://datacenter.eastmoney.com/securities/api/data/v1/get"
	params := map[string]string{
		"reportName":   "RPT_F10_EH_EQUITY",
		"columns":      "SECUCODE,SECURITY_CODE,END_DATE,TOTAL_SHARES,LIMITED_SHARES,LIMITED_OTHARS,LIMITED_DOMESTIC_NATURAL,LIMITED_STATE_LEGAL,LIMITED_OVERSEAS_NOSTATE,LIMITED_OVERSEAS_NATURAL,UNLIMITED_SHARES,LISTED_A_SHARES,B_FREE_SHARE,H_FREE_SHARE,FREE_SHARES,LIMITED_A_SHARES,NON_FREE_SHARES,LIMITED_B_SHARES,OTHER_FREE_SHARES,LIMITED_STATE_SHARES,LIMITED_DOMESTIC_NOSTATE,LOCK_SHARES,LIMITED_FOREIGN_SHARES,LIMITED_H_SHARES,SPONSOR_SHARES,STATE_SPONSOR_SHARES,SPONSOR_SOCIAL_SHARES,RAISE_SHARES,RAISE_STATE_SHARES,RAISE_DOMESTIC_SHARES,RAISE_OVERSEAS_SHARES,CHANGE_REASON",
		"quoteColumns": "",
		"filter":       fmt.Sprintf(`(SECUCODE="%s")`, url.QueryEscape(stockCodeWithExchange)),
		"pageNumber":   "1",
		"pageSize":     "200",
		"sortTypes":    "-1",
		"sortColumns":  "END_DATE",
		"source":       "HSF10",
		"client":       "PC",
	}

	var result struct {
		Success bool `json:"success"`
		Result  struct {
			Data []struct {
				SecurityCode  string  `json:"SECURITY_CODE"`
				EndDate       string  `json:"END_DATE"`
				TotalShares   float64 `json:"TOTAL_SHARES"`
				LimitedShares float64 `json:"LIMITED_SHARES"`
				ListedAShares float64 `json:"LISTED_A_SHARES"`
				ChangeReason  string  `json:"CHANGE_REASON"`
			} `json:"data"`
		} `json:"result"`
	}

	err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到股本数据", "")
	}

	var shares []types.StockShares
	for _, item := range result.Result.Data {
		share := types.StockShares{
			StockCode:    item.SecurityCode,
			ChangeDate:   item.EndDate,
			TotalShares:  item.TotalShares,
			LimitShares:  item.LimitedShares,
			ListAShares:  item.ListedAShares,
			ChangeReason: item.ChangeReason,
		}
		shares = append(shares, share)
	}

	// 如果不需要历史数据，只返回最新的
	if !isHistory && len(shares) > 0 {
		shares = shares[:1]
	}

	return shares, nil
}

// GetIndustrySW 获取申万行业信息
func (s *StockInfo) GetIndustrySW(stockCode string) ([]types.IndustrySW, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	// 构建请求参数
	codeList := []map[string]string{{"code": stockCode, "market": "ab", "type": "stock"}}
	codeListJSON, _ := json.Marshal(codeList)

	baseURL := "https://finance.pae.baidu.com/api/getrelatedblock"
	params := map[string]string{
		"stock":         string(codeListJSON),
		"finClientType": "pc",
	}

	var result struct {
		Result map[string][]struct {
			Name string `json:"name"`
			List []struct {
				Name     string `json:"name"`
				Describe string `json:"describe"`
				XcxQuery string `json:"xcx_query"`
			} `json:"list"`
		} `json:"Result"`
	}

	err := s.client.GetJSON(baseURL, params, headers.GetBaiduHeaders(), &result)
	if err != nil {
		return nil, err
	}

	if result.Result == nil {
		return nil, errors.NewADataError(errors.ErrNoDataFound.Code, "未找到行业数据", "")
	}

	var industries []types.IndustrySW
	for key, value := range result.Result {
		for _, industryType := range value {
			if industryType.Name == "行业" {
				for _, item := range industryType.List {
					// 解析xcx_query参数获取code
					u, _ := url.ParseQuery(item.XcxQuery)
					code := ""
					if codes := u["code"]; len(codes) > 0 {
						code = codes[0]
					}

					industry := types.IndustrySW{
						StockCode:    key,
						SWCode:       code,
						IndustryName: item.Name,
						IndustryType: item.Describe,
						Source:       "百度股市通",
					}
					industries = append(industries, industry)
				}
			}
		}
	}

	return industries, nil
}

// TradeCalendar 获取交易日历
func (s *StockInfo) TradeCalendar(year int) ([]types.TradeCalendar, error) {
	// 如果没有指定年份，默认使用当前年份
	if year == 0 {
		year = time.Now().Year()
	}

	// 获取深交所交易日历
	return s.getCalendarFromSZSE(year)
}

// getCalendarFromSZSE 从深交所获取交易日历
func (s *StockInfo) getCalendarFromSZSE(year int) ([]types.TradeCalendar, error) {
	var allData []types.TradeCalendar

	// 遍历12个月
	for i := 1; i <= 12; i++ {
		apiURL := fmt.Sprintf("http://www.szse.cn/api/report/exchange/onepersistenthour/monthList?month=%d-%02d", year, i)

		var result struct {
			Data []types.TradeCalendar `json:"data"`
		}

		err := s.client.GetJSON(apiURL, nil, nil, &result)
		if err != nil {
			continue
		}

		// 结果为空跳出循环
		if len(result.Data) == 0 {
			break
		}

		allData = append(allData, result.Data...)
	}

	return allData, nil
}
