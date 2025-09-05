// Package types 定义项目中使用的所有数据类型
package types

import "time"

// StockCode 股票代码信息
type StockCode struct {
	StockCode string `json:"stock_code"` // 股票代码
	ShortName string `json:"short_name"` // 股票简称
	Exchange  string `json:"exchange"`   // 交易所
	ListDate  string `json:"list_date"`  // 上市日期
}

// IndexCode 指数代码信息
type IndexCode struct {
	IndexCode string `json:"index_code"` // 指数代码
	IndexName string `json:"index_name"` // 指数名称
	Exchange  string `json:"exchange"`   // 交易所
}

// ConceptCode 概念代码信息
type ConceptCode struct {
	ConceptCode string `json:"concept_code"` // 概念代码
	ConceptName string `json:"concept_name"` // 概念名称
}

// ConceptInfo 概念详细信息
type ConceptInfo struct {
	ConceptCode string `json:"concept_code"` // 概念代码
	ConceptName string `json:"concept_name"` // 概念名称
	Source      string `json:"source"`       // 数据来源
}

// MarketParams 行情查询参数
type MarketParams struct {
	StockCode  string    `json:"stock_code"`  // 股票代码
	StartDate  time.Time `json:"start_date"`  // 开始日期
	EndDate    time.Time `json:"end_date"`    // 结束日期
	KType      int       `json:"k_type"`      // K线类型: 1-日线, 2-周线, 3-月线
	AdjustType int       `json:"adjust_type"` // 复权类型: 0-不复权, 1-前复权, 2-后复权
}

// MarketData 行情数据
type MarketData struct {
	StockCode string  `json:"stock_code"` // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Open      float64 `json:"open"`       // 开盘价
	High      float64 `json:"high"`       // 最高价
	Low       float64 `json:"low"`        // 最低价
	Close     float64 `json:"close"`      // 收盘价
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额
	Change    float64 `json:"change"`     // 涨跌额
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Turnover  float64 `json:"turnover"`   // 换手率
	PreClose  float64 `json:"pre_close"`  // 前收盘价
}

// CurrentMarket 当前行情数据
type CurrentMarket struct {
	StockCode  string  `json:"stock_code"`  // 股票代码
	ShortName  string  `json:"short_name"`  // 股票简称
	Price      float64 `json:"price"`       // 当前价格
	Change     float64 `json:"change"`      // 涨跌额
	ChangePct  float64 `json:"change_pct"`  // 涨跌幅
	Volume     int64   `json:"volume"`      // 成交量
	Amount     float64 `json:"amount"`      // 成交额
	Open       float64 `json:"open"`        // 开盘价
	High       float64 `json:"high"`        // 最高价
	Low        float64 `json:"low"`         // 最低价
	PreClose   float64 `json:"pre_close"`   // 昨收价
	Turnover   float64 `json:"turnover"`    // 换手率
	MarketCap  float64 `json:"market_cap"`  // 总市值
	CircMarket float64 `json:"circ_market"` // 流通市值
	PB         float64 `json:"pb"`          // 市净率
	PE         float64 `json:"pe"`          // 市盈率
}

// FinanceData 财务数据
type FinanceData struct {
	StockCode    string  `json:"stock_code"`     // 股票代码
	ReportDate   string  `json:"report_date"`    // 报告期
	EPS          float64 `json:"eps"`            // 每股收益
	Revenue      float64 `json:"revenue"`        // 营业收入
	NetProfit    float64 `json:"net_profit"`     // 净利润
	TotalAssets  float64 `json:"total_assets"`   // 总资产
	TotalEquity  float64 `json:"total_equity"`   // 股东权益
	NetAssetPS   float64 `json:"net_asset_ps"`   // 每股净资产
	EPSYOY       float64 `json:"eps_yoy"`        // 每股收益同比
	RevenueYOY   float64 `json:"revenue_yoy"`    // 营收同比
	NetProfitYOY float64 `json:"net_profit_yoy"` // 净利润同比
	ROE          float64 `json:"roe"`            // 净资产收益率
	DebtRatio    float64 `json:"debt_ratio"`     // 资产负债率
	GrossProfit  float64 `json:"gross_profit"`   // 毛利润
	GrossProfitR float64 `json:"gross_profit_r"` // 毛利率
}

// ETFInfo ETF基础信息
type ETFInfo struct {
	ETFCode     string  `json:"etf_code"`     // ETF代码
	ETFName     string  `json:"etf_name"`     // ETF名称
	Exchange    string  `json:"exchange"`     // 交易所
	ListDate    string  `json:"list_date"`    // 上市日期
	TotalShares float64 `json:"total_shares"` // 总份额
	NetValue    float64 `json:"net_value"`    // 净值
}

// BondInfo 债券信息
type BondInfo struct {
	BondCode string `json:"bond_code"` // 债券代码
	BondName string `json:"bond_name"` // 债券名称
	Exchange string `json:"exchange"`  // 交易所
	ListDate string `json:"list_date"` // 上市日期
}

// StockShares 股本信息
type StockShares struct {
	StockCode    string  `json:"stock_code"`    // 股票代码
	ChangeDate   string  `json:"change_date"`   // 变更日期
	TotalShares  float64 `json:"total_shares"`  // 总股本
	LimitShares  float64 `json:"limit_shares"`  // 限售股本
	ListAShares  float64 `json:"list_a_shares"` // 流通A股
	ChangeReason string  `json:"change_reason"` // 变更原因
}

// IndustrySW 申万行业信息
type IndustrySW struct {
	StockCode    string `json:"stock_code"`    // 股票代码
	SWCode       string `json:"sw_code"`       // 申万行业代码
	IndustryName string `json:"industry_name"` // 行业名称
	IndustryType string `json:"industry_type"` // 行业类型
	Source       string `json:"source"`        // 数据来源
}

// TradeCalendar 交易日历
type TradeCalendar struct {
	TradeDate   string `json:"trade_date"`   // 交易日期
	TradeStatus int    `json:"trade_status"` // 交易状态:1，交易日；0，非交易日
	DayWeek     int    `json:"day_week"`     // 一周的第几天，从星期日开始
}

// MarketMin 分时行情数据
type MarketMin struct {
	StockCode string    `json:"stock_code"` // 股票代码
	TradeTime time.Time `json:"trade_time"` // 交易时间
	Price     float64   `json:"price"`      // 价格
	Change    float64   `json:"change"`     // 涨跌额
	ChangePct float64   `json:"change_pct"` // 涨跌幅
	AvgPrice  float64   `json:"avg_price"`  // 均价
	Volume    int64     `json:"volume"`     // 成交量
	Amount    float64   `json:"amount"`     // 成交额
}

// MarketFive 五档行情
type MarketFive struct {
	StockCode string  `json:"stock_code"` // 股票代码
	Price     float64 `json:"price"`      // 当前价格
	Change    float64 `json:"change"`     // 涨跌额
	ChangePct float64 `json:"change_pct"` // 涨跌幅
	Volume    int64   `json:"volume"`     // 成交量
	Amount    float64 `json:"amount"`     // 成交额

	// 五档买盘
	BuyPrices  [5]float64 `json:"buy_prices"`  // 买一到买五价格
	BuyVolumes [5]int64   `json:"buy_volumes"` // 买一到买五成交量

	// 五档卖盘
	SellPrices  [5]float64 `json:"sell_prices"`  // 卖一到卖五价格
	SellVolumes [5]int64   `json:"sell_volumes"` // 卖一到卖五成交量
}

// CapitalFlow 资金流向
type CapitalFlow struct {
	StockCode      string    `json:"stock_code"`       // 股票代码
	TradeDate      time.Time `json:"trade_date"`       // 交易日期
	MainInflow     float64   `json:"main_inflow"`      // 主力净流入
	MainInflowRate float64   `json:"main_inflow_rate"` // 主力净流入占比
	SuperInflow    float64   `json:"super_inflow"`     // 超大单净流入
	LargeInflow    float64   `json:"large_inflow"`     // 大单净流入
	MediumInflow   float64   `json:"medium_inflow"`    // 中单净流入
	SmallInflow    float64   `json:"small_inflow"`     // 小单净流入
}

// FinanceCore 核心财务数据
type FinanceCore struct {
	StockCode  string `json:"stock_code"`  // 股票代码
	ShortName  string `json:"short_name"`  // 股票简称
	ReportDate string `json:"report_date"` // 报告期
	ReportType string `json:"report_type"` // 报告类型
	NoticeDate string `json:"notice_date"` // 公告日期

	// 每股指标
	BasicEPS       float64 `json:"basic_eps"`        // 基本每股收益
	DilutedEPS     float64 `json:"diluted_eps"`      // 稀释每股收益
	NonGaapEPS     float64 `json:"non_gaap_eps"`     // 非GAAP每股收益
	NetAssetPS     float64 `json:"net_asset_ps"`     // 每股净资产
	CapReservePS   float64 `json:"cap_reserve_ps"`   // 每股资本公积
	UndistProfitPS float64 `json:"undist_profit_ps"` // 每股未分配利润
	OperCFPS       float64 `json:"oper_cf_ps"`       // 每股经营现金流

	// 盈利能力
	TotalRev              float64 `json:"total_rev"`                  // 营业总收入
	GrossProfit           float64 `json:"gross_profit"`               // 毛利润
	NetProfitAttrSH       float64 `json:"net_profit_attr_sh"`         // 归属母公司净利润
	NonGaapNetProfit      float64 `json:"non_gaap_net_profit"`        // 非GAAP净利润
	TotalRevYoYGR         float64 `json:"total_rev_yoy_gr"`           // 营业总收入同比增长
	NetProfitYoYGR        float64 `json:"net_profit_yoy_gr"`          // 净利润同比增长
	NonGaapNetProfitYoYGR float64 `json:"non_gaap_net_profit_yoy_gr"` // 非GAAP净利润同比增长
	TotalRevQoQGR         float64 `json:"total_rev_qoq_gr"`           // 营业总收入环比增长
	NetProfitQoQGR        float64 `json:"net_profit_qoq_gr"`          // 净利润环比增长
	NonGaapNetProfitQoQGR float64 `json:"non_gaap_net_profit_qoq_gr"` // 非GAAP净利润环比增长

	// 盈利质量
	ROEWtd        float64 `json:"roe_wtd"`          // 加权平均净资产收益率
	ROENonGaapWtd float64 `json:"roe_non_gaap_wtd"` // 非GAAP加权平均净资产收益率
	ROAWtd        float64 `json:"roa_wtd"`          // 总资产收益率
	GrossMargin   float64 `json:"gross_margin"`     // 毛利率
	NetMargin     float64 `json:"net_margin"`       // 净利率

	// 现金流相关
	AdvReceiptsToRev float64 `json:"adv_receipts_to_rev"` // 预收账款占营收比例
	NetCFSalesToRev  float64 `json:"net_cf_sales_to_rev"` // 销售现金流占营收比例
	OperCFToRev      float64 `json:"oper_cf_to_rev"`      // 经营现金流占营收比例
	EffTaxRate       float64 `json:"eff_tax_rate"`        // 有效税率

	// 偿债能力
	CurrRatio        float64 `json:"curr_ratio"`        // 流动比率
	QuickRatio       float64 `json:"quick_ratio"`       // 速动比率
	CashFlowRatio    float64 `json:"cash_flow_ratio"`   // 现金流量比率
	AssetLiabRatio   float64 `json:"asset_liab_ratio"`  // 资产负债率
	EquityMultiplier float64 `json:"equity_multiplier"` // 权益乘数
	EquityRatio      float64 `json:"equity_ratio"`      // 产权比率

	// 运营能力
	TotalAssetTurnDays float64 `json:"total_asset_turn_days"` // 总资产周转天数
	InvTurnDays        float64 `json:"inv_turn_days"`         // 存货周转天数
	AcctRecvTurnDays   float64 `json:"acct_recv_turn_days"`   // 应收账款周转天数
	TotalAssetTurnRate float64 `json:"total_asset_turn_rate"` // 总资产周转率
	InvTurnRate        float64 `json:"inv_turn_rate"`         // 存货周转率
	AcctRecvTurnRate   float64 `json:"acct_recv_turn_rate"`   // 应收账款周转率
}

// BalanceSheet 资产负债表
type BalanceSheet struct {
	StockCode  string `json:"stock_code"`  // 股票代码
	ReportDate string `json:"report_date"` // 报告期
	ReportType string `json:"report_type"` // 报告类型
	NoticeDate string `json:"notice_date"` // 公告日期

	// 资产
	TotalAssets            float64 `json:"total_assets"`              // 资产总计
	CurrentAssets          float64 `json:"current_assets"`            // 流动资产
	NonCurrentAssets       float64 `json:"non_current_assets"`        // 非流动资产
	CashAndCashEquivalents float64 `json:"cash_and_cash_equivalents"` // 货币资金
	AccountsReceivable     float64 `json:"accounts_receivable"`       // 应收账款
	Inventory              float64 `json:"inventory"`                 // 存货
	FixedAssets            float64 `json:"fixed_assets"`              // 固定资产
	IntangibleAssets       float64 `json:"intangible_assets"`         // 无形资产

	// 负债
	TotalLiabilities      float64 `json:"total_liabilities"`       // 负债总计
	CurrentLiabilities    float64 `json:"current_liabilities"`     // 流动负债
	NonCurrentLiabilities float64 `json:"non_current_liabilities"` // 非流动负债
	ShortTermBorrowing    float64 `json:"short_term_borrowing"`    // 短期借款
	AccountsPayable       float64 `json:"accounts_payable"`        // 应付账款
	LongTermBorrowing     float64 `json:"long_term_borrowing"`     // 长期借款

	// 所有者权益
	TotalEquity      float64 `json:"total_equity"`      // 所有者权益合计
	ShareCapital     float64 `json:"share_capital"`     // 实收资本(股本)
	CapitalReserve   float64 `json:"capital_reserve"`   // 资本公积
	RetainedEarnings float64 `json:"retained_earnings"` // 未分配利润
}

// CashFlow 现金流量表
type CashFlow struct {
	StockCode  string `json:"stock_code"`  // 股票代码
	ReportDate string `json:"report_date"` // 报告期
	ReportType string `json:"report_type"` // 报告类型
	NoticeDate string `json:"notice_date"` // 公告日期

	// 经营活动现金流
	NetCashFlowsOperAct   float64 `json:"net_cash_flows_oper_act"`  // 经营活动产生的现金流量净额
	CashInflowsOperAct    float64 `json:"cash_inflows_oper_act"`    // 经营活动现金流入小计
	CashOutflowsOperAct   float64 `json:"cash_outflows_oper_act"`   // 经营活动现金流出小计
	SalesServicesRender   float64 `json:"sales_services_render"`    // 销售商品、提供劳务收到的现金
	TaxRefunds            float64 `json:"tax_refunds"`              // 收到的税费返还
	OtherCashInflowsOper  float64 `json:"other_cash_inflows_oper"`  // 收到其他与经营活动有关的现金
	PurchaseGoodsServices float64 `json:"purchase_goods_services"`  // 购买商品、接受劳务支付的现金
	PaymentStaffBenefits  float64 `json:"payment_staff_benefits"`   // 支付给职工以及为职工支付的现金
	PaymentsTaxes         float64 `json:"payments_taxes"`           // 支付的各项税费
	OtherCashOutflowsOper float64 `json:"other_cash_outflows_oper"` // 支付其他与经营活动有关的现金

	// 投资活动现金流
	NetCashFlowsInvAct     float64 `json:"net_cash_flows_inv_act"`   // 投资活动产生的现金流量净额
	CashInflowsInvAct      float64 `json:"cash_inflows_inv_act"`     // 投资活动现金流入小计
	CashOutflowsInvAct     float64 `json:"cash_outflows_inv_act"`    // 投资活动现金流出小计
	RecoveryInvestments    float64 `json:"recovery_investments"`     // 收回投资收到的现金
	InvestIncomeReceived   float64 `json:"invest_income_received"`   // 取得投资收益收到的现金
	DisposalAssetsReceived float64 `json:"disposal_assets_received"` // 处置固定资产、无形资产和其他长期资产收回的现金净额
	PurchaseAssets         float64 `json:"purchase_assets"`          // 购建固定资产、无形资产和其他长期资产支付的现金
	InvestPayments         float64 `json:"invest_payments"`          // 投资支付的现金

	// 筹资活动现金流
	NetCashFlowsFinAct  float64 `json:"net_cash_flows_fin_act"` // 筹资活动产生的现金流量净额
	CashInflowsFinAct   float64 `json:"cash_inflows_fin_act"`   // 筹资活动现金流入小计
	CashOutflowsFinAct  float64 `json:"cash_outflows_fin_act"`  // 筹资活动现金流出小计
	BorrowingsReceived  float64 `json:"borrowings_received"`    // 取得借款收到的现金
	IssueSharesBonds    float64 `json:"issue_shares_bonds"`     // 吸收投资收到的现金
	RepaymentBorrowings float64 `json:"repayment_borrowings"`   // 偿还债务支付的现金
	DividendsPaid       float64 `json:"dividends_paid"`         // 分配股利、利润或偿付利息支付的现金

	// 现金净增加额
	NetIncreaseCash float64 `json:"net_increase_cash"` // 现金及现金等价物净增加额
	CashBeginPeriod float64 `json:"cash_begin_period"` // 期初现金及现金等价物余额
	CashEndPeriod   float64 `json:"cash_end_period"`   // 期末现金及现金等价物余额
}

// Profit 利润表
type Profit struct {
	StockCode  string `json:"stock_code"`  // 股票代码
	ReportDate string `json:"report_date"` // 报告期
	ReportType string `json:"report_type"` // 报告类型
	NoticeDate string `json:"notice_date"` // 公告日期

	// 营业收入
	TotalOperatingRevenue float64 `json:"total_operating_revenue"` // 营业总收入
	OperatingRevenue      float64 `json:"operating_revenue"`       // 营业收入
	InterestIncome        float64 `json:"interest_income"`         // 利息收入
	PremiumsEarned        float64 `json:"premiums_earned"`         // 已赚保费
	CommissionIncome      float64 `json:"commission_income"`       // 手续费及佣金收入

	// 营业总成本
	TotalOperatingCost     float64 `json:"total_operating_cost"`     // 营业总成本
	OperatingCost          float64 `json:"operating_cost"`           // 营业成本
	InterestExpense        float64 `json:"interest_expense"`         // 利息支出
	CommissionExpense      float64 `json:"commission_expense"`       // 手续费及佣金支出
	SurrenderValue         float64 `json:"surrender_value"`          // 退保金
	NetCompensationExpense float64 `json:"net_compensation_expense"` // 赔付支出净额
	NetAmortizationExpense float64 `json:"net_amortization_expense"` // 提取保险责任合同准备金净额
	PolicyBonusExpense     float64 `json:"policy_bonus_expense"`     // 保单红利支出
	TaxesSurcharges        float64 `json:"taxes_surcharges"`         // 税金及附加
	SalesExpense           float64 `json:"sales_expense"`            // 销售费用
	AdminExpense           float64 `json:"admin_expense"`            // 管理费用
	FinExpense             float64 `json:"fin_expense"`              // 财务费用
	AssetImpairmentLoss    float64 `json:"asset_impairment_loss"`    // 资产减值损失
	CreditImpairmentLoss   float64 `json:"credit_impairment_loss"`   // 信用减值损失

	// 营业利润
	GrossProfit         float64 `json:"gross_profit"`          // 毛利润
	OperatingProfit     float64 `json:"operating_profit"`      // 营业利润
	NonOperatingIncome  float64 `json:"non_operating_income"`  // 营业外收入
	NonOperatingExpense float64 `json:"non_operating_expense"` // 营业外支出
	LossDisposalAssets  float64 `json:"loss_disposal_assets"`  // 资产处置收益

	// 利润总额
	TotalProfit      float64 `json:"total_profit"`       // 利润总额
	IncomeTaxExpense float64 `json:"income_tax_expense"` // 所得税费用

	// 净利润
	NetProfit             float64 `json:"net_profit"`              // 净利润
	NetProfitAttrSH       float64 `json:"net_profit_attr_sh"`      // 归属于母公司股东的净利润
	NetProfitMinority     float64 `json:"net_profit_minority"`     // 少数股东损益
	NetProfitContinuing   float64 `json:"net_profit_continuing"`   // 持续经营净利润
	NetProfitDiscontinued float64 `json:"net_profit_discontinued"` // 终止经营净利润

	// 每股收益
	BasicEPS   float64 `json:"basic_eps"`   // 基本每股收益
	DilutedEPS float64 `json:"diluted_eps"` // 稀释每股收益
}
