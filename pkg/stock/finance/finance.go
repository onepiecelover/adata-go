// Package finance 提供股票财务数据相关功能
package finance

import (
	"fmt"
	"net/url"

	"github.com/onepiecelover/adata-go/pkg/common/client"
	"github.com/onepiecelover/adata-go/pkg/common/errors"
	"github.com/onepiecelover/adata-go/pkg/common/headers"
	"github.com/onepiecelover/adata-go/pkg/common/utils"
	"github.com/onepiecelover/adata-go/pkg/types"
)

// StockFinance 股票财务数据结构体
type StockFinance struct {
	client *client.Client
}

// NewStockFinance 创建股票财务数据实例
func NewStockFinance() *StockFinance {
	return &StockFinance{
		client: client.NewClient(),
	}
}

// SetProxy 设置代理
func (s *StockFinance) SetProxy(enabled bool, proxyURL string) {
	s.client.SetProxy(enabled, proxyURL)
}

// GetCoreIndex 获取核心财务数据
func (s *StockFinance) GetCoreIndex(stockCode string) ([]types.FinanceCore, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	// 报告类型：年报、中报、三季报、一季报
	reportTypes := []string{"年报", "中报", "三季报", "一季报"}

	var allData []types.FinanceCore

	for _, reportType := range reportTypes {
		baseURL := "https://datacenter.eastmoney.com/securities/api/data/get"
		params := map[string]string{
			"type":         "RPT_F10_FINANCE_MAINFINADATA",
			"sty":          "APP_F10_MAINFINADATA",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")(REPORT_TYPE="%s")`, url.QueryEscape(stockCodeWithExchange), url.QueryEscape(reportType)),
			"p":            "1",
			"ps":           "100",
			"sr":           "-1",
			"st":           "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}

		var result struct {
			Code   int `json:"code"`
			Result struct {
				Data []struct {
					SecurityCode       string  `json:"SECURITY_CODE"`
					SecurityNameAbbr   string  `json:"SECURITY_NAME_ABBR"`
					ReportDate         string  `json:"REPORT_DATE"`
					ReportType         string  `json:"REPORT_TYPE"`
					NoticeDate         string  `json:"NOTICE_DATE"`
					EPSJB              float64 `json:"EPSJB"`
					EPSKCJB            float64 `json:"EPSKCJB"`
					EPSXS              float64 `json:"EPSXS"`
					BPS                float64 `json:"BPS"`
					MGZBGJ             float64 `json:"MGZBGJ"`
					MGWFPLR            float64 `json:"MGWFPLR"`
					MGJYXJJE           float64 `json:"MGJYXJJE"`
					TOTALOPERATEREVE   float64 `json:"TOTALOPERATEREVE"`
					MLR                float64 `json:"MLR"`
					PARENTNETPROFIT    float64 `json:"PARENTNETPROFIT"`
					KCFJCXSYJLR        float64 `json:"KCFJCXSYJLR"`
					TOTALOPERATEREVETZ float64 `json:"TOTALOPERATEREVETZ"`
					PARENTNETPROFITTZ  float64 `json:"PARENTNETPROFITTZ"`
					KCFJCXSYJLRTZ      float64 `json:"KCFJCXSYJLRTZ"`
					YYZSRGDHBZC        float64 `json:"YYZSRGDHBZC"`
					NETPROFITRPHBZC    float64 `json:"NETPROFITRPHBZC"`
					KFJLRGDHBZC        float64 `json:"KFJLRGDHBZC"`
					ROEJQ              float64 `json:"ROEJQ"`
					ROEKCJQ            float64 `json:"ROEKCJQ"`
					ZZCJLL             float64 `json:"ZZCJLL"`
					XSMLL              float64 `json:"XSMLL"`
					XSJLL              float64 `json:"XSJLL"`
					YSZKYYSR           float64 `json:"YSZKYYSR"`
					XSJXLYYSR          float64 `json:"XSJXLYYSR"`
					JYXJLYYSR          float64 `json:"JYXJLYYSR"`
					TAXRATE            float64 `json:"TAXRATE"`
					LD                 float64 `json:"LD"`
					SD                 float64 `json:"SD"`
					XJLLB              float64 `json:"XJLLB"`
					ZCFZL              float64 `json:"ZCFZL"`
					QYCS               float64 `json:"QYCS"`
					CQBL               float64 `json:"CQBL"`
					ZZCZZTS            float64 `json:"ZZCZZTS"`
					CHZZTS             float64 `json:"CHZZTS"`
					YSZKZZTS           float64 `json:"YSZKZZTS"`
					TOAZZL             float64 `json:"TOAZZL"`
					CHZZL              float64 `json:"CHZZL"`
					YSZKZZL            float64 `json:"YSZKZZL"`
				} `json:"data"`
			} `json:"result"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			continue
		}

		if result.Code == 0 && result.Result.Data != nil {
			for _, item := range result.Result.Data {
				financeCore := types.FinanceCore{
					StockCode:             item.SecurityCode,
					ShortName:             item.SecurityNameAbbr,
					ReportDate:            item.ReportDate,
					ReportType:            item.ReportType,
					NoticeDate:            item.NoticeDate,
					BasicEPS:              item.EPSJB,
					DilutedEPS:            item.EPSKCJB,
					NonGaapEPS:            item.EPSXS,
					NetAssetPS:            item.BPS,
					CapReservePS:          item.MGZBGJ,
					UndistProfitPS:        item.MGWFPLR,
					OperCFPS:              item.MGJYXJJE,
					TotalRev:              item.TOTALOPERATEREVE,
					GrossProfit:           item.MLR,
					NetProfitAttrSH:       item.PARENTNETPROFIT,
					NonGaapNetProfit:      item.KCFJCXSYJLR,
					TotalRevYoYGR:         item.TOTALOPERATEREVETZ,
					NetProfitYoYGR:        item.PARENTNETPROFITTZ,
					NonGaapNetProfitYoYGR: item.KCFJCXSYJLRTZ,
					TotalRevQoQGR:         item.YYZSRGDHBZC,
					NetProfitQoQGR:        item.NETPROFITRPHBZC,
					NonGaapNetProfitQoQGR: item.KFJLRGDHBZC,
					ROEWtd:                item.ROEJQ,
					ROENonGaapWtd:         item.ROEKCJQ,
					ROAWtd:                item.ZZCJLL,
					GrossMargin:           item.XSMLL,
					NetMargin:             item.XSJLL,
					AdvReceiptsToRev:      item.YSZKYYSR,
					NetCFSalesToRev:       item.XSJXLYYSR,
					OperCFToRev:           item.JYXJLYYSR,
					EffTaxRate:            item.TAXRATE,
					CurrRatio:             item.LD,
					QuickRatio:            item.SD,
					CashFlowRatio:         item.XJLLB,
					AssetLiabRatio:        item.ZCFZL,
					EquityMultiplier:      item.QYCS,
					EquityRatio:           item.CQBL,
					TotalAssetTurnDays:    item.ZZCZZTS,
					InvTurnDays:           item.CHZZTS,
					AcctRecvTurnDays:      item.YSZKZZTS,
					TotalAssetTurnRate:    item.TOAZZL,
					InvTurnRate:           item.CHZZL,
					AcctRecvTurnRate:      item.YSZKZZL,
				}
				allData = append(allData, financeCore)
			}
		}
	}

	return allData, nil
}

// GetBalance 获取资产负债表数据
func (s *StockFinance) GetBalance(stockCode string) ([]types.BalanceSheet, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	// 报告类型：年报、中报、三季报、一季报
	reportTypes := []string{"年报", "中报", "三季报", "一季报"}

	var allData []types.BalanceSheet

	for _, reportType := range reportTypes {
		baseURL := "https://datacenter.eastmoney.com/securities/api/data/get"
		params := map[string]string{
			"type":         "RPT_F10_FINANCE_BALANCE",
			"sty":          "APP_F10_BALANCE",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")(REPORT_TYPE="%s")`, url.QueryEscape(stockCodeWithExchange), url.QueryEscape(reportType)),
			"p":            "1",
			"ps":           "100",
			"sr":           "-1",
			"st":           "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}

		var result struct {
			Code   int `json:"code"`
			Result struct {
				Data []struct {
					SecurityCode           string  `json:"SECURITY_CODE"`
					ReportDate             string  `json:"REPORT_DATE"`
					ReportType             string  `json:"REPORT_TYPE"`
					NoticeDate             string  `json:"NOTICE_DATE"`
					TotalAssets            float64 `json:"TOTAL_ASSETS"`
					CurrentAssets          float64 `json:"CURRENT_ASSETS"`
					NonCurrentAssets       float64 `json:"NON_CURRENT_ASSETS"`
					CashAndCashEquivalents float64 `json:"CASH_AND_CASH_EQUIVALENTS"`
					AccountsReceivable     float64 `json:"ACCOUNTS_RECEIVABLE"`
					Inventory              float64 `json:"INVENTORY"`
					FixedAssets            float64 `json:"FIXED_ASSETS"`
					IntangibleAssets       float64 `json:"INTANGIBLE_ASSETS"`
					TotalLiabilities       float64 `json:"TOTAL_LIABILITIES"`
					CurrentLiabilities     float64 `json:"CURRENT_LIABILITIES"`
					NonCurrentLiabilities  float64 `json:"NON_CURRENT_LIABILITIES"`
					ShortTermBorrowing     float64 `json:"SHORT_TERM_BORROWING"`
					AccountsPayable        float64 `json:"ACCOUNTS_PAYABLE"`
					LongTermBorrowing      float64 `json:"LONG_TERM_BORROWING"`
					TotalEquity            float64 `json:"TOTAL_EQUITY"`
					ShareCapital           float64 `json:"SHARE_CAPITAL"`
					CapitalReserve         float64 `json:"CAPITAL_RESERVE"`
					RetainedEarnings       float64 `json:"RETAINED_EARNINGS"`
				} `json:"data"`
			} `json:"result"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			continue
		}

		if result.Code == 0 && result.Result.Data != nil {
			for _, item := range result.Result.Data {
				balanceSheet := types.BalanceSheet{
					StockCode:              item.SecurityCode,
					ReportDate:             item.ReportDate,
					ReportType:             item.ReportType,
					NoticeDate:             item.NoticeDate,
					TotalAssets:            item.TotalAssets,
					CurrentAssets:          item.CurrentAssets,
					NonCurrentAssets:       item.NonCurrentAssets,
					CashAndCashEquivalents: item.CashAndCashEquivalents,
					AccountsReceivable:     item.AccountsReceivable,
					Inventory:              item.Inventory,
					FixedAssets:            item.FixedAssets,
					IntangibleAssets:       item.IntangibleAssets,
					TotalLiabilities:       item.TotalLiabilities,
					CurrentLiabilities:     item.CurrentLiabilities,
					NonCurrentLiabilities:  item.NonCurrentLiabilities,
					ShortTermBorrowing:     item.ShortTermBorrowing,
					AccountsPayable:        item.AccountsPayable,
					LongTermBorrowing:      item.LongTermBorrowing,
					TotalEquity:            item.TotalEquity,
					ShareCapital:           item.ShareCapital,
					CapitalReserve:         item.CapitalReserve,
					RetainedEarnings:       item.RetainedEarnings,
				}
				allData = append(allData, balanceSheet)
			}
		}
	}

	return allData, nil
}

// GetCashFlow 获取现金流量表数据
func (s *StockFinance) GetCashFlow(stockCode string) ([]types.CashFlow, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	// 报告类型：年报、中报、三季报、一季报
	reportTypes := []string{"年报", "中报", "三季报", "一季报"}

	var allData []types.CashFlow

	for _, reportType := range reportTypes {
		baseURL := "https://datacenter.eastmoney.com/securities/api/data/get"
		params := map[string]string{
			"type":         "RPT_F10_FINANCE_CASHFLOW",
			"sty":          "APP_F10_CASHFLOW",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")(REPORT_TYPE="%s")`, url.QueryEscape(stockCodeWithExchange), url.QueryEscape(reportType)),
			"p":            "1",
			"ps":           "100",
			"sr":           "-1",
			"st":           "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}

		var result struct {
			Code   int `json:"code"`
			Result struct {
				Data []struct {
					SecurityCode           string  `json:"SECURITY_CODE"`
					ReportDate             string  `json:"REPORT_DATE"`
					ReportType             string  `json:"REPORT_TYPE"`
					NoticeDate             string  `json:"NOTICE_DATE"`
					NetCashFlowsOperAct    float64 `json:"NET_CASH_FLOWS_OPER_ACT"`
					CashInflowsOperAct     float64 `json:"CASH_INFLOWS_OPER_ACT"`
					CashOutflowsOperAct    float64 `json:"CASH_OUTFLOWS_OPER_ACT"`
					SalesServicesRender    float64 `json:"SALES_SERVICES_RENDER"`
					TaxRefunds             float64 `json:"TAX_REFUNDS"`
					OtherCashInflowsOper   float64 `json:"OTHER_CASH_INFLOWS_OPER"`
					PurchaseGoodsServices  float64 `json:"PURCHASE_GOODS_SERVICES"`
					PaymentStaffBenefits   float64 `json:"PAYMENT_STAFF_BENEFITS"`
					PaymentsTaxes          float64 `json:"PAYMENTS_TAXES"`
					OtherCashOutflowsOper  float64 `json:"OTHER_CASH_OUTFLOWS_OPER"`
					NetCashFlowsInvAct     float64 `json:"NET_CASH_FLOWS_INV_ACT"`
					CashInflowsInvAct      float64 `json:"CASH_INFLOWS_INV_ACT"`
					CashOutflowsInvAct     float64 `json:"CASH_OUTFLOWS_INV_ACT"`
					RecoveryInvestments    float64 `json:"RECOVERY_INVESTMENTS"`
					InvestIncomeReceived   float64 `json:"INVEST_INCOME_RECEIVED"`
					DisposalAssetsReceived float64 `json:"DISPOSAL_ASSETS_RECEIVED"`
					PurchaseAssets         float64 `json:"PURCHASE_ASSETS"`
					InvestPayments         float64 `json:"INVEST_PAYMENTS"`
					NetCashFlowsFinAct     float64 `json:"NET_CASH_FLOWS_FIN_ACT"`
					CashInflowsFinAct      float64 `json:"CASH_INFLOWS_FIN_ACT"`
					CashOutflowsFinAct     float64 `json:"CASH_OUTFLOWS_FIN_ACT"`
					BorrowingsReceived     float64 `json:"BORROWINGS_RECEIVED"`
					IssueSharesBonds       float64 `json:"ISSUE_SHARES_BONDS"`
					RepaymentBorrowings    float64 `json:"REPAYMENT_BORROWINGS"`
					DividendsPaid          float64 `json:"DIVIDENDS_PAID"`
					NetIncreaseCash        float64 `json:"NET_INCREASE_CASH"`
					CashBeginPeriod        float64 `json:"CASH_BEGIN_PERIOD"`
					CashEndPeriod          float64 `json:"CASH_END_PERIOD"`
				} `json:"data"`
			} `json:"result"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			continue
		}

		if result.Code == 0 && result.Result.Data != nil {
			for _, item := range result.Result.Data {
				cashFlow := types.CashFlow{
					StockCode:              item.SecurityCode,
					ReportDate:             item.ReportDate,
					ReportType:             item.ReportType,
					NoticeDate:             item.NoticeDate,
					NetCashFlowsOperAct:    item.NetCashFlowsOperAct,
					CashInflowsOperAct:     item.CashInflowsOperAct,
					CashOutflowsOperAct:    item.CashOutflowsOperAct,
					SalesServicesRender:    item.SalesServicesRender,
					TaxRefunds:             item.TaxRefunds,
					OtherCashInflowsOper:   item.OtherCashInflowsOper,
					PurchaseGoodsServices:  item.PurchaseGoodsServices,
					PaymentStaffBenefits:   item.PaymentStaffBenefits,
					PaymentsTaxes:          item.PaymentsTaxes,
					OtherCashOutflowsOper:  item.OtherCashOutflowsOper,
					NetCashFlowsInvAct:     item.NetCashFlowsInvAct,
					CashInflowsInvAct:      item.CashInflowsInvAct,
					CashOutflowsInvAct:     item.CashOutflowsInvAct,
					RecoveryInvestments:    item.RecoveryInvestments,
					InvestIncomeReceived:   item.InvestIncomeReceived,
					DisposalAssetsReceived: item.DisposalAssetsReceived,
					PurchaseAssets:         item.PurchaseAssets,
					InvestPayments:         item.InvestPayments,
					NetCashFlowsFinAct:     item.NetCashFlowsFinAct,
					CashInflowsFinAct:      item.CashInflowsFinAct,
					CashOutflowsFinAct:     item.CashOutflowsFinAct,
					BorrowingsReceived:     item.BorrowingsReceived,
					IssueSharesBonds:       item.IssueSharesBonds,
					RepaymentBorrowings:    item.RepaymentBorrowings,
					DividendsPaid:          item.DividendsPaid,
					NetIncreaseCash:        item.NetIncreaseCash,
					CashBeginPeriod:        item.CashBeginPeriod,
					CashEndPeriod:          item.CashEndPeriod,
				}
				allData = append(allData, cashFlow)
			}
		}
	}

	return allData, nil
}

// GetProfit 获取利润表数据
func (s *StockFinance) GetProfit(stockCode string) ([]types.Profit, error) {
	if !utils.IsValidStockCode(stockCode) {
		return nil, errors.ErrInvalidStockCode
	}

	stockCodeWithExchange := utils.CompileExchangeByStockCode(stockCode)

	// 报告类型：年报、中报、三季报、一季报
	reportTypes := []string{"年报", "中报", "三季报", "一季报"}

	var allData []types.Profit

	for _, reportType := range reportTypes {
		baseURL := "https://datacenter.eastmoney.com/securities/api/data/get"
		params := map[string]string{
			"type":         "RPT_F10_FINANCE_PROFIT",
			"sty":          "APP_F10_PROFIT",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(SECUCODE="%s")(REPORT_TYPE="%s")`, url.QueryEscape(stockCodeWithExchange), url.QueryEscape(reportType)),
			"p":            "1",
			"ps":           "100",
			"sr":           "-1",
			"st":           "REPORT_DATE",
			"source":       "HSF10",
			"client":       "PC",
		}

		var result struct {
			Code   int `json:"code"`
			Result struct {
				Data []struct {
					SecurityCode           string  `json:"SECURITY_CODE"`
					ReportDate             string  `json:"REPORT_DATE"`
					ReportType             string  `json:"REPORT_TYPE"`
					NoticeDate             string  `json:"NOTICE_DATE"`
					TotalOperatingRevenue  float64 `json:"TOTAL_OPERATING_REVENUE"`
					OperatingRevenue       float64 `json:"OPERATING_REVENUE"`
					InterestIncome         float64 `json:"INTEREST_INCOME"`
					PremiumsEarned         float64 `json:"PREMIUMS_EARNED"`
					CommissionIncome       float64 `json:"COMMISSION_INCOME"`
					TotalOperatingCost     float64 `json:"TOTAL_OPERATING_COST"`
					OperatingCost          float64 `json:"OPERATING_COST"`
					InterestExpense        float64 `json:"INTEREST_EXPENSE"`
					CommissionExpense      float64 `json:"COMMISSION_EXPENSE"`
					SurrenderValue         float64 `json:"SURRENDER_VALUE"`
					NetCompensationExpense float64 `json:"NET_COMPENSATION_EXPENSE"`
					NetAmortizationExpense float64 `json:"NET_AMORTIZATION_EXPENSE"`
					PolicyBonusExpense     float64 `json:"POLICY_BONUS_EXPENSE"`
					TaxesSurcharges        float64 `json:"TAXES_SURCHARGES"`
					SalesExpense           float64 `json:"SALES_EXPENSE"`
					AdminExpense           float64 `json:"ADMIN_EXPENSE"`
					FinExpense             float64 `json:"FIN_EXPENSE"`
					AssetImpairmentLoss    float64 `json:"ASSET_IMPAIRMENT_LOSS"`
					CreditImpairmentLoss   float64 `json:"CREDIT_IMPAIRMENT_LOSS"`
					GrossProfit            float64 `json:"GROSS_PROFIT"`
					OperatingProfit        float64 `json:"OPERATING_PROFIT"`
					NonOperatingIncome     float64 `json:"NON_OPERATING_INCOME"`
					NonOperatingExpense    float64 `json:"NON_OPERATING_EXPENSE"`
					LossDisposalAssets     float64 `json:"LOSS_DISPOSAL_ASSETS"`
					TotalProfit            float64 `json:"TOTAL_PROFIT"`
					IncomeTaxExpense       float64 `json:"INCOME_TAX_EXPENSE"`
					NetProfit              float64 `json:"NET_PROFIT"`
					NetProfitAttrSH        float64 `json:"NET_PROFIT_ATTR_SH"`
					NetProfitMinority      float64 `json:"NET_PROFIT_MINORITY"`
					NetProfitContinuing    float64 `json:"NET_PROFIT_CONTINUING"`
					NetProfitDiscontinued  float64 `json:"NET_PROFIT_DISCONTINUED"`
					BasicEPS               float64 `json:"BASIC_EPS"`
					DilutedEPS             float64 `json:"DILUTED_EPS"`
				} `json:"data"`
			} `json:"result"`
		}

		err := s.client.GetJSON(baseURL, params, headers.EastMoneyHeaders, &result)
		if err != nil {
			continue
		}

		if result.Code == 0 && result.Result.Data != nil {
			for _, item := range result.Result.Data {
				profit := types.Profit{
					StockCode:              item.SecurityCode,
					ReportDate:             item.ReportDate,
					ReportType:             item.ReportType,
					NoticeDate:             item.NoticeDate,
					TotalOperatingRevenue:  item.TotalOperatingRevenue,
					OperatingRevenue:       item.OperatingRevenue,
					InterestIncome:         item.InterestIncome,
					PremiumsEarned:         item.PremiumsEarned,
					CommissionIncome:       item.CommissionIncome,
					TotalOperatingCost:     item.TotalOperatingCost,
					OperatingCost:          item.OperatingCost,
					InterestExpense:        item.InterestExpense,
					CommissionExpense:      item.CommissionExpense,
					SurrenderValue:         item.SurrenderValue,
					NetCompensationExpense: item.NetCompensationExpense,
					NetAmortizationExpense: item.NetAmortizationExpense,
					PolicyBonusExpense:     item.PolicyBonusExpense,
					TaxesSurcharges:        item.TaxesSurcharges,
					SalesExpense:           item.SalesExpense,
					AdminExpense:           item.AdminExpense,
					FinExpense:             item.FinExpense,
					AssetImpairmentLoss:    item.AssetImpairmentLoss,
					CreditImpairmentLoss:   item.CreditImpairmentLoss,
					GrossProfit:            item.GrossProfit,
					OperatingProfit:        item.OperatingProfit,
					NonOperatingIncome:     item.NonOperatingIncome,
					NonOperatingExpense:    item.NonOperatingExpense,
					LossDisposalAssets:     item.LossDisposalAssets,
					TotalProfit:            item.TotalProfit,
					IncomeTaxExpense:       item.IncomeTaxExpense,
					NetProfit:              item.NetProfit,
					NetProfitAttrSH:        item.NetProfitAttrSH,
					NetProfitMinority:      item.NetProfitMinority,
					NetProfitContinuing:    item.NetProfitContinuing,
					NetProfitDiscontinued:  item.NetProfitDiscontinued,
					BasicEPS:               item.BasicEPS,
					DilutedEPS:             item.DilutedEPS,
				}
				allData = append(allData, profit)
			}
		}
	}

	return allData, nil
}
