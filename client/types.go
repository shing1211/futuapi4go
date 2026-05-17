// Package client provides a public Client type for the Futu OpenD SDK.
// This allows external projects to use the SDK.
package client

import "github.com/shing1211/futuapi4go/pkg/pb/qotcommon"

// Quote represents a real-time quote.
type Quote struct {
	Symbol         string  `json:"symbol"`
	Market         int32   `json:"market"`
	Price          float64 `json:"price"`
	Open           float64 `json:"open"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Volume         int64   `json:"volume"`
	Timestamp      string  `json:"timestamp"`
	Name           string  `json:"name"`
	LastClose      float64 `json:"lastClose"`
	Turnover       float64 `json:"turnover"`
	TurnoverRate   float64 `json:"turnoverRate"`
	Amplitude      float64 `json:"amplitude"`
	IsSuspended    bool    `json:"isSuspended"`
	SecStatus      int32   `json:"secStatus"`
	ListTime       string  `json:"listTime"`
	PriceSpread    float64 `json:"priceSpread"`
	DarkStatus     int32   `json:"darkStatus"`
	ListTimestamp  float64 `json:"listTimestamp"`
	UpdateTimestamp float64 `json:"updateTimestamp"`
}

// KLine represents a K-line (candlestick) data point.
type KLine struct {
	Time         string  `json:"time"`
	IsBlank      bool    `json:"isBlank"`
	Open         float64 `json:"open"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Close        float64 `json:"close"`
	Volume       int64   `json:"volume"`
	LastClose    float64 `json:"lastClose"`
	Turnover     float64 `json:"turnover"`
	TurnoverRate float64 `json:"turnoverRate"`
	ChangeRate   float64 `json:"changeRate"`
	Timestamp    float64 `json:"timestamp"`
}

// Account represents a trading account.
type Account struct {
	AccID             uint64 `json:"accID"`
	AccType           int32 `json:"accType"`
	TrdEnv            int32 `json:"trdEnv"`
	CardNum           string `json:"cardNum"`
	AccStatus         int32 `json:"accStatus"`
	TrdMarketAuthList []int32 `json:"trdMarketAuthList"`
	SecurityFirm      int32 `json:"securityFirm"`
	SimAccType        int32 `json:"simAccType"`
	UniCardNum        string `json:"uniCardNum"`
	AccRole           int32 `json:"accRole"`
	JpAccType         []int32 `json:"jpAccType"`
}

// PlaceOrderResult represents a place order result.
type PlaceOrderResult struct {
	OrderID   uint64 `json:"orderID"`
	OrderIDEx string `json:"orderIDEx"`
}

// Position represents a position.
type Position struct {
	PositionID       uint64  `json:"positionID"`
	PositionSide      int32   `json:"positionSide"`
	Code             string  `json:"code"`
	Name             string  `json:"name"`
	Market           int32   `json:"market"`
	Quantity         float64 `json:"quantity"`
	CanSellQty       float64 `json:"canSellQty"`
	CostPrice        float64 `json:"costPrice"`
	CurPrice         float64 `json:"curPrice"`
	MarketVal        float64 `json:"marketVal"`
	PnL              float64 `json:"pnL"`
	PnLRate          float64 `json:"pnLRate"`
	TodayBuyQty      float64 `json:"todayBuyQty"`
	TodayBuyVal      float64 `json:"todayBuyVal"`
	TodaySellQty     float64 `json:"todaySellQty"`
	TodaySellVal     float64 `json:"todaySellVal"`
	TodayPnL         float64 `json:"todayPnL"`
	UnrealizedPL     float64 `json:"unrealizedPL"`
	RealizedPL       float64 `json:"realizedPL"`
	Currency         int32   `json:"currency"`
	TrdMarket        int32   `json:"trdMarket"`
	DilutedCostPrice float64 `json:"dilutedCostPrice"`
	AverageCostPrice float64 `json:"averageCostPrice"`
	AveragePnLRate   float64 `json:"averagePnLRate"`
	SecMarket        int32   `json:"secMarket"`
	TdTrdVal         float64 `json:"tdTrdVal"`
}

// AccCashInfo represents per-currency cash (futures accounts).
type AccCashInfo struct {
	Currency         int32 `json:"currency"`
	Cash             float64 `json:"cash"`
	AvailableBalance float64 `json:"availableBalance"`
	NetCashPower     float64 `json:"netCashPower"`
}

// AccMarketInfo represents per-market assets.
type AccMarketInfo struct {
	TrdMarket int32 `json:"trdMarket"`
	Assets    float64 `json:"assets"`
}

// Funds represents account funds.
type Funds struct {
	Power             float64 `json:"power"`
	TotalAssets       float64 `json:"totalAssets"`
	Cash              float64 `json:"cash"`
	MarketVal         float64 `json:"marketVal"`
	FrozenCash        float64 `json:"frozenCash"`
	DebtCash          float64 `json:"debtCash"`
	AvlWithdrawalCash float64 `json:"avlWithdrawalCash"`
	Currency          int32 `json:"currency"`
	AvailableFunds    float64 `json:"availableFunds"`
	UnrealizedPL      float64 `json:"unrealizedPL"`
	RealizedPL        float64 `json:"realizedPL"`
	RiskLevel         int32 `json:"riskLevel"`
	InitialMargin     float64 `json:"initialMargin"`
	MaintenanceMargin float64 `json:"maintenanceMargin"`
	MaxPowerShort     float64 `json:"maxPowerShort"`
	NetCashPower      float64 `json:"netCashPower"`
	LongMv            float64 `json:"longMv"`
	ShortMv           float64 `json:"shortMv"`
	PendingAsset      float64 `json:"pendingAsset"`
	MaxWithdrawal     float64 `json:"maxWithdrawal"`
	RiskStatus        int32 `json:"riskStatus"`
	MarginCallMargin  float64 `json:"marginCallMargin"`
	// IsPDT indicates whether the account is a Pattern Day Trader (US margin accounts).
	IsPDT  bool `json:"isPDT"`
	// PDTSeq is the PDT sequence number.
	PDTSeq string `json:"pDTSeq"`
	BeginningDTBP     float64 `json:"beginningDTBP"`
	RemainingDTBP     float64 `json:"remainingDTBP"`
	DtCallAmount      float64 `json:"dtCallAmount"`
	DtStatus          int32 `json:"dtStatus"`
	CashInfoList      []AccCashInfo `json:"cashInfoList"`
	MarketInfoList    []AccMarketInfo `json:"marketInfoList"`
	SecuritiesAssets  float64 `json:"securitiesAssets"`
	FundAssets       float64 `json:"fundAssets"`
	BondAssets       float64 `json:"bondAssets"`
}

// Order represents an order.
type Order struct {
	OrderID         uint64 `json:"orderID"`
	OrderIDEx       string `json:"orderIDEx"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	TrdSide         int32 `json:"trdSide"`
	OrderType       int32 `json:"orderType"`
	OrderStatus     int32 `json:"orderStatus"`
	Price           float64 `json:"price"`
	Qty             float64 `json:"qty"`
	FillQty         float64 `json:"fillQty"`
	FillAvgPrice    float64 `json:"fillAvgPrice"`
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updateTime"`
	LastErrMsg      string `json:"lastErrMsg"`
	SecMarket       int32 `json:"secMarket"`
	CreateTimestamp float64 `json:"createTimestamp"`
	UpdateTimestamp float64 `json:"updateTimestamp"`
	Remark          string `json:"remark"`
	TimeInForce     int32 `json:"timeInForce"`
	FillOutsideRTH  bool `json:"fillOutsideRTH"`
	AuxPrice        float64 `json:"auxPrice"`
	TrailType       int32 `json:"trailType"`
	TrailValue      float64 `json:"trailValue"`
	TrailSpread     float64 `json:"trailSpread"`
	Currency        int32 `json:"currency"`
	TrdMarket       int32 `json:"trdMarket"`
	Session         int32 `json:"session"`
	JpAccType       int32 `json:"jpAccType"`
}

// OrderFill represents an order fill.
type OrderFill struct {
	FillID            uint64 `json:"fillID"`
	FillIDEx          string `json:"fillIDEx"`
	OrderID           uint64 `json:"orderID"`
	OrderIDEx         string `json:"orderIDEx"`
	Code              string `json:"code"`
	Name              string `json:"name"`
	TrdSide           int32 `json:"trdSide"`
	Price             float64 `json:"price"`
	Qty               float64 `json:"qty"`
	CreateTime        string `json:"createTime"`
	CounterBrokerID   int32 `json:"counterBrokerID"`
	CounterBrokerName string `json:"counterBrokerName"`
	SecMarket         int32 `json:"secMarket"`
	CreateTimestamp   float64 `json:"createTimestamp"`
	UpdateTimestamp   float64 `json:"updateTimestamp"`
	Status            int32 `json:"status"`
	TrdMarket         int32 `json:"trdMarket"`
	JpAccType         int32 `json:"jpAccType"`
}

// OrderBook represents order book data.
type OrderBook struct {
	Bids                    []OrderBookItem `json:"bids"`
	Asks                    []OrderBookItem `json:"asks"`
	SvrRecvTimeBid          string `json:"svrRecvTimeBid"`
	SvrRecvTimeBidTimestamp float64 `json:"svrRecvTimeBidTimestamp"`
	SvrRecvTimeAsk          string `json:"svrRecvTimeAsk"`
	SvrRecvTimeAskTimestamp float64 `json:"svrRecvTimeAskTimestamp"`
}

// OrderBookItem represents a single order book entry.
type OrderBookItem struct {
	Price      float64 `json:"price"`
	Volume     int64 `json:"volume"`
	OrderCount int32 `json:"orderCount"`
	DetailList []OrderBookDetail `json:"detailList"`
}

// OrderBookDetail represents a detail entry in an order book item.
type OrderBookDetail struct {
	OrderID int64 `json:"orderID"`
	Volume  int64 `json:"volume"`
}

// Ticker represents ticker data.
type Ticker struct {
	Time         string  `json:"time"`
	Sequence     int64   `json:"sequence"`
	Price        float64 `json:"price"`
	Volume       int64   `json:"volume"`
	Direction    string  `json:"direction"`
	Turnover     float64 `json:"turnover"`
	RecvTime     float64 `json:"recvTime"`
	Type         int32   `json:"type"`
	TypeSign     int32   `json:"typeSign"`
	Timestamp    float64 `json:"timestamp"`
	PushDataType int32   `json:"pushDataType"`
}

// RT represents real-time data.
type RT struct {
	Time      string  `json:"time"`
	Minute    int32   `json:"minute"`
	IsBlank   bool    `json:"isBlank"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	LastClose float64 `json:"lastClose"`
	AvgPrice  float64 `json:"avgPrice"`
	Turnover  float64 `json:"turnover"`
	Timestamp float64 `json:"timestamp"`
}

// Broker represents broker data.
type Broker struct {
	ID      int64  `json:"iD"`
	Name    string `json:"name"`
	Pos     int32  `json:"pos"`
	Volume  int64  `json:"volume"`
	OrderID int64  `json:"orderID"`
}

// StaticInfo represents static security info.
type StaticInfo struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Type     int32 `json:"type"`
	ListTime string `json:"listTime"`
	LotSize  int32 `json:"lotSize"`
}

// CapitalFlow represents capital flow data.
type CapitalFlow struct {
	Time        string `json:"time"`
	InFlow      float64 `json:"inFlow"`
	MainInFlow  float64 `json:"mainInFlow"`
	SuperInFlow float64 `json:"superInFlow"`
	BigInFlow   float64 `json:"bigInFlow"`
	MidInFlow   float64 `json:"midInFlow"`
	SmlInFlow   float64 `json:"smlInFlow"`
	Timestamp   float64 `json:"timestamp"`
}

// CapitalDistribution represents capital distribution.
type CapitalDistribution struct {
	MainInflow      float64 `json:"mainInflow"`
	MainOutflow     float64 `json:"mainOutflow"`
	MidInflow       float64 `json:"midInflow"`
	MidOutflow      float64 `json:"midOutflow"`
	SmallInflow     float64 `json:"smallInflow"`
	SmallOutflow    float64 `json:"smallOutflow"`
	BigInflow       float64 `json:"bigInflow"`
	BigOutflow      float64 `json:"bigOutflow"`
	UpdateTime      string `json:"updateTime"`
	UpdateTimestamp float64 `json:"updateTimestamp"`
}

// OptionExpiration represents option expiration date.
type OptionExpiration struct {
	Date string `json:"date"`
	Days int32 `json:"days"`
	Desc string `json:"desc"`
}

// FutureInfo represents futures info.
type FutureInfo struct {
	Code               string `json:"code"`
	Name               string `json:"name"`
	Expire             string `json:"expire"`
	LastTradeTimestamp float64 `json:"lastTradeTimestamp"`
	Owner              string `json:"owner"`
	OwnerOther         string `json:"ownerOther"`
	Exchange           string `json:"exchange"`
	ContractType       string `json:"contractType"`
	ContractSize       float64 `json:"contractSize"`
	ContractSizeUnit   string `json:"contractSizeUnit"`
	QuoteCurrency      string `json:"quoteCurrency"`
	MinVar             float64 `json:"minVar"`
	MinVarUnit         string `json:"minVarUnit"`
	QuoteUnit          string `json:"quoteUnit"`
	TimeZone           string `json:"timeZone"`
	ExchangeFormatUrl  string `json:"exchangeFormatUrl"`
}

// Plate represents a market plate (æ¿å—).
type Plate struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// IpoData represents IPO data.
type IpoData struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	ListDate      string `json:"listDate"`
	ListTimestamp float64 `json:"listTimestamp"`
}

// UserSecurityGroup represents user security group.
type UserSecurityGroup struct {
	Name      string `json:"name"`
	GroupType int32 `json:"groupType"`
}

// SubInfo represents subscription info.
type SubInfo struct {
	IsSub    bool `json:"isSub"`
	SubTypes []int32 `json:"subTypes"`
	Security string `json:"security"`
}

// StockFilterResult represents a single stock filter result.
type StockFilterResult struct {
	Security   *qotcommon.Security `json:"security"`
	Name       string `json:"name"`
	CurPrice   float64 `json:"curPrice"`
	ChangeRate float64 `json:"changeRate"`
	Volume     int64 `json:"volume"`
	Turnover   float64 `json:"turnover"`
	HighPrice  float64 `json:"highPrice"`
	LowPrice   float64 `json:"lowPrice"`
}

// OptChainItem represents a pair of call and put options at the same strike price.
type OptChainItem struct {
	Call *qotcommon.SecurityStaticInfo `json:"call"`
	Put  *qotcommon.SecurityStaticInfo `json:"put"`
}

// OptChain represents the option chain for a single expiration date.
type OptChain struct {
	StrikeTime      string `json:"strikeTime"`
	StrikeTimestamp float64 `json:"strikeTimestamp"`
	Option          []*OptChainItem `json:"option"`
}

// WarrantData represents warrant data.
type WarrantData struct {
	Stock              *qotcommon.Security `json:"stock"`
	Owner              *qotcommon.Security `json:"owner"`
	Type               int32 `json:"type"`
	Issuer             int32 `json:"issuer"`
	MaturityTime       string `json:"maturityTime"`
	MaturityTimestamp  float64 `json:"maturityTimestamp"`
	ListTime           string `json:"listTime"`
	ListTimestamp      float64 `json:"listTimestamp"`
	LastTradeTime      string `json:"lastTradeTime"`
	LastTradeTimestamp float64 `json:"lastTradeTimestamp"`
	RecoveryPrice      float64 `json:"recoveryPrice"`
	ConversionRatio    float64 `json:"conversionRatio"`
	LotSize            int32 `json:"lotSize"`
	StrikePrice        float64 `json:"strikePrice"`
	LastClosePrice     float64 `json:"lastClosePrice"`
	Name               string `json:"name"`
	CurPrice           float64 `json:"curPrice"`
	PriceChangeVal     float64 `json:"priceChangeVal"`
	ChangeRate         float64 `json:"changeRate"`
	Status             int32 `json:"status"`
	BidPrice           float64 `json:"bidPrice"`
	AskPrice           float64 `json:"askPrice"`
	BidVol             int64 `json:"bidVol"`
	AskVol             int64 `json:"askVol"`
	Volume             int64 `json:"volume"`
	Turnover           float64 `json:"turnover"`
	Score              float64 `json:"score"`
	Premium            float64 `json:"premium"`
	BreakEvenPoint     float64 `json:"breakEvenPoint"`
	Leverage           float64 `json:"leverage"`
	Ipop               float64 `json:"ipop"`
	PriceRecoveryRatio float64 `json:"priceRecoveryRatio"`
	ConversionPrice    float64 `json:"conversionPrice"`
	StreetRate         float64 `json:"streetRate"`
	StreetVol          int64 `json:"streetVol"`
	Amplitude          float64 `json:"amplitude"`
	IssueSize          int64 `json:"issueSize"`
	HighPrice          float64 `json:"highPrice"`
	LowPrice           float64 `json:"lowPrice"`
	ImpliedVolatility  float64 `json:"impliedVolatility"`
	Delta              float64 `json:"delta"`
	EffectiveLeverage  float64 `json:"effectiveLeverage"`
	UpperStrikePrice   float64 `json:"upperStrikePrice"`
	LowerStrikePrice   float64 `json:"lowerStrikePrice"`
	InLinePriceStatus  int32 `json:"inLinePriceStatus"`
}

// Snapshot represents security snapshot data.
type Snapshot struct {
	Security                *qotcommon.Security `json:"security"`
	Name                    string `json:"name"`
	Type                    int32 `json:"type"`
	IsSuspend               bool `json:"isSuspend"`
	LotSize                 int32 `json:"lotSize"`
	CurPrice                float64 `json:"curPrice"`
	ChangeVal               float64 `json:"changeVal"`
	HighPrice               float64 `json:"highPrice"`
	LowPrice                float64 `json:"lowPrice"`
	OpenPrice               float64 `json:"openPrice"`
	LastClose               float64 `json:"lastClose"`
	Volume                  int64 `json:"volume"`
	Turnover                float64 `json:"turnover"`
	ListTime                string `json:"listTime"`
	PriceSpread             float64 `json:"priceSpread"`
	UpdateTime              string `json:"updateTime"`
	TurnoverRate            float64 `json:"turnoverRate"`
	ListTimestamp           float64 `json:"listTimestamp"`
	UpdateTimestamp         float64 `json:"updateTimestamp"`
	AskPrice                float64 `json:"askPrice"`
	BidPrice                float64 `json:"bidPrice"`
	AskVol                  int64 `json:"askVol"`
	BidVol                  int64 `json:"bidVol"`
	EnableMargin            bool `json:"enableMargin"`
	MortgageRatio           float64 `json:"mortgageRatio"`
	LongMarginInitialRatio  float64 `json:"longMarginInitialRatio"`
	EnableShortSell         bool `json:"enableShortSell"`
	ShortSellRate           float64 `json:"shortSellRate"`
	ShortAvailableVolume    int64 `json:"shortAvailableVolume"`
	ShortMarginInitialRatio float64 `json:"shortMarginInitialRatio"`
	Amplitude               float64 `json:"amplitude"`
	AvgPrice                float64 `json:"avgPrice"`
	BidAskRatio             float64 `json:"bidAskRatio"`
	VolumeRatio             float64 `json:"volumeRatio"`
	Highest52WeeksPrice     float64 `json:"highest52WeeksPrice"`
	Lowest52WeeksPrice      float64 `json:"lowest52WeeksPrice"`
	HighestHistoryPrice     float64 `json:"highestHistoryPrice"`
	LowestHistoryPrice      float64 `json:"lowestHistoryPrice"`
	SecStatus               int32 `json:"secStatus"`
	ClosePrice5Minute       float64 `json:"closePrice5Minute"`
}

// CodeChangeInfo represents information about a code change.
type CodeChangeInfo struct {
	Type            int32 `json:"type"`
	Security        *qotcommon.Security `json:"security"`
	RelatedSecurity *qotcommon.Security `json:"relatedSecurity"`
	PublicTime      string `json:"publicTime"`
	EffectiveTime   string `json:"effectiveTime"`
}

// GlobalState represents global connection state.
type GlobalState struct {
	ServerVer         int32  `json:"serverVer"`
	ServerBuildNo     int32  `json:"serverBuildNo"`
	Time              int64  `json:"time"`
	LocalTime         float64 `json:"localTime"`
	QotLogined        bool   `json:"qotLogined"`
	TrdLogined        bool   `json:"trdLogined"`
	MarketHK          int32  `json:"marketHK"`
	MarketUS          int32  `json:"marketUS"`
	MarketSH          int32  `json:"marketSH"`
	MarketSZ          int32  `json:"marketSZ"`
	MarketHKFuture    int32  `json:"marketHKFuture"`
	MarketUSFuture    int32  `json:"marketUSFuture"`
	MarketSGFuture    int32  `json:"marketSGFuture"`
	MarketJPFuture    int32  `json:"marketJPFuture"`
	ProgramStatus     int32  `json:"programStatus"`
	ProgramStatusDesc string `json:"programStatusDesc"`
	ConnID            uint64 `json:"connID"`
	QotSvrIpAddr      string `json:"qotSvrIpAddr"`
	TrdSvrIpAddr      string `json:"trdSvrIpAddr"`
}

// UserInfo represents user information.
type UserInfo struct {
	UserID               int64  `json:"userID"`
	NickName             string `json:"nickName"`
	AvatarUrl            string `json:"avatarUrl"`
	ApiLevel             string `json:"apiLevel"`
	IsNeedAgreeDisclaimer bool   `json:"isNeedAgreeDisclaimer"`
	ShQotRight           int32  `json:"shQotRight"`
	SzQotRight           int32  `json:"szQotRight"`
	Extra                int32  `json:"extra"`
	HkQotRight           int32  `json:"hkQotRight"`
	UsQotRight           int32  `json:"usQotRight"`
	CnQotRight           int32  `json:"cnQotRight"`
	SubQuota             int32  `json:"subQuota"`
	HistoryKLQuota       int32  `json:"historyKLQuota"`
}

// DelayStatistics represents delay statistics for Qot push.
type DelayStatistics struct {
	QotPushType    int32 `json:"qotPushType"`
	DelayAvg       float64 `json:"delayAvg"`
	Count          int32 `json:"count"`
	ItemList       []DelayStatisticsItem `json:"itemList"`
	ReqReplyList   []ReqReplyStatisticsItem `json:"reqReplyList"`
	PlaceOrderList []PlaceOrderStatisticsItem `json:"placeOrderList"`
}

// DelayStatisticsItem represents a single delay statistics item.
type DelayStatisticsItem struct {
	Begin           int32 `json:"begin"`
	End             int32 `json:"end"`
	Count           int32 `json:"count"`
	Proportion      float64 `json:"proportion"`
	CumulativeRatio float64 `json:"cumulativeRatio"`
}

// ReqReplyStatisticsItem represents request-reply statistics.
type ReqReplyStatisticsItem struct {
	ProtoID      int32 `json:"protoID"`
	Count        int32 `json:"count"`
	TotalCostAvg float64 `json:"totalCostAvg"`
	OpenDCostAvg float64 `json:"openDCostAvg"`
	NetDelayAvg  float64 `json:"netDelayAvg"`
	IsLocalReply bool `json:"isLocalReply"`
}

// PlaceOrderStatisticsItem represents order placement statistics.
type PlaceOrderStatisticsItem struct {
	OrderID    string `json:"orderID"`
	TotalCost  float64 `json:"totalCost"`
	OpenDCost  float64 `json:"openDCost"`
	NetDelay   float64 `json:"netDelay"`
	UpdateCost float64 `json:"updateCost"`
}

// TestCmdResult represents the result of a test command sent to OpenD.
type TestCmdResult struct {
	Cmd    string `json:"cmd"`
	Result string `json:"result"`
}

// SuspendInfo represents suspension time for a security.
type SuspendInfo struct {
	Time      string `json:"time"`
	Timestamp float64 `json:"timestamp"`
}

// PriceReminderInfo represents a price reminder.
type PriceReminderInfo struct {
	Security *qotcommon.Security `json:"security"`
	Name     string `json:"name"`
	ItemList []PriceReminderItemInfo `json:"itemList"`
}

// PriceReminderItemInfo represents a single price reminder item.
type PriceReminderItemInfo struct {
	Key                 int64 `json:"key"`
	Type                int32 `json:"type"`
	Freq                int32 `json:"freq"`
	Value               float64 `json:"value"`
	Note                string `json:"note"`
	IsEnable            bool `json:"isEnable"`
	ReminderSessionList []int32 `json:"reminderSessionList"`
}

// ReconfirmOrderResult represents an order reconfirmation result.
type ReconfirmOrderResult struct {
	AccID     uint64 `json:"accID"`
	TrdEnv    int32 `json:"trdEnv"`
	TrdMarket int32 `json:"trdMarket"`
	OrderID   uint64 `json:"orderID"`
}

// HoldingChangeInfo represents a holding change entry.
type HoldingChangeInfo struct {
	HolderName   string `json:"holderName"`
	HoldingQty   float64 `json:"holdingQty"`
	HoldingRatio float64 `json:"holdingRatio"`
	ChangeQty    float64 `json:"changeQty"`
	ChangeRatio  float64 `json:"changeRatio"`
	Time         string `json:"time"`
	Timestamp    float64 `json:"timestamp"`
}

// RehabInfo represents rehabilitation (å¤æƒ) data.
type RehabInfo struct {
	Time       string `json:"time"`
	FwdFactorA float64 `json:"fwdFactorA"`
	FwdFactorB float64 `json:"fwdFactorB"`
	BwdFactorA float64 `json:"bwdFactorA"`
	BwdFactorB float64 `json:"bwdFactorB"`
	SplitBase  int32 `json:"splitBase"`
	SplitErt   int32 `json:"splitErt"`
	JoinBase   int32 `json:"joinBase"`
	JoinErt    int32 `json:"joinErt"`
	BonusBase  int32 `json:"bonusBase"`
	BonusErt   int32 `json:"bonusErt"`
	AllotBase  int32 `json:"allotBase"`
	AllotErt   int32 `json:"allotErt"`
	AllotPrice float64 `json:"allotPrice"`
}

// HistoryKLQuotaInfo represents historical K-line quota info.
type HistoryKLQuotaInfo struct {
	UsedQuota   int32 `json:"usedQuota"`
	RemainQuota int32 `json:"remainQuota"`
	DetailList  []HistoryKLQuotaDetail `json:"detailList"`
}

// HistoryKLQuotaDetail represents a single quota detail entry.
type HistoryKLQuotaDetail struct {
	Security         *qotcommon.Security `json:"security"`
	Name             string `json:"name"`
	RequestTime      string `json:"requestTime"`
	RequestTimestamp int64 `json:"requestTimestamp"`
}

// FlowSummaryInfo represents a single cash flow entry.
type FlowSummaryInfo struct {
	CashFlowID        uint64 `json:"cashFlowID"`
	ClearingDate      string `json:"clearingDate"`
	SettlementDate    string `json:"settlementDate"`
	Currency          int32 `json:"currency"`
	CashFlowType      string `json:"cashFlowType"`
	CashFlowDirection int32 `json:"cashFlowDirection"`
	CashFlowAmount    float64 `json:"cashFlowAmount"`
	CashFlowRemark    string `json:"cashFlowRemark"`
}

// OrderFeeInfo represents fee information for an order.
type OrderFeeInfo struct {
	OrderIDEx string `json:"orderIDEx"`
	FeeAmount float64 `json:"feeAmount"`
	FeeList   []OrderFeeItemInfo `json:"feeList"`
}

// OrderFeeItemInfo represents a single fee item.
type OrderFeeItemInfo struct {
	Title string `json:"title"`
	Value float64 `json:"value"`
}

// MarginRatioInfo represents margin ratio for a security.
type MarginRatioInfo struct {
	Security       *qotcommon.Security `json:"security"`
	IsLongPermit   bool    `json:"isLongPermit"`
	IsShortPermit  bool    `json:"isShortPermit"`
	ShortFeeRate   float64 `json:"shortFeeRate"`
	ImLongRatio    float64 `json:"imLongRatio"`
	ImShortRatio   float64 `json:"imShortRatio"`
	ShortPoolRemain float64 `json:"shortPoolRemain"`
	AlertLongRatio  float64 `json:"alertLongRatio"`
	AlertShortRatio float64 `json:"alertShortRatio"`
	McmLongRatio   float64 `json:"mcmLongRatio"`
	McmShortRatio  float64 `json:"mcmShortRatio"`
	MmLongRatio    float64 `json:"mmLongRatio"`
	MmShortRatio   float64 `json:"mmShortRatio"`
}

// AccTradingInfo represents trading capability for a security.
type AccTradingInfo struct {
	MaxCashBuy          float64 `json:"maxCashBuy"`
	MaxCashAndMarginBuy float64 `json:"maxCashAndMarginBuy"`
	MaxPositionSell     float64 `json:"maxPositionSell"`
	MaxSellShort        float64 `json:"maxSellShort"`
	MaxBuyBack          float64 `json:"maxBuyBack"`
	LongRequiredIM      float64 `json:"longRequiredIM"`
	ShortRequiredIM     float64 `json:"shortRequiredIM"`
}

// MaxTrdQtysInfo represents maximum tradable quantities.
type MaxTrdQtysInfo struct {
	MaxCashBuy          float64 `json:"maxCashBuy"`
	MaxCashAndMarginBuy float64 `json:"maxCashAndMarginBuy"`
	MaxPositionSell     float64 `json:"maxPositionSell"`
	MaxSellShort        float64 `json:"maxSellShort"`
	MaxBuyBack          float64 `json:"maxBuyBack"`
	LongRequiredIM      float64 `json:"longRequiredIM"`
	ShortRequiredIM     float64 `json:"shortRequiredIM"`
}

// PushQuote represents a parsed real-time quote push notification.
type PushQuote struct {
	Market    int32 `json:"market"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CurPrice  float64 `json:"curPrice"`
	OpenPrice float64 `json:"openPrice"`
	HighPrice float64 `json:"highPrice"`
	LowPrice  float64 `json:"lowPrice"`
	Volume    int64 `json:"volume"`
	Turnover  float64 `json:"turnover"`
}

// PushKLine represents a parsed K-line push notification.
type PushKLine struct {
	Market int32 `json:"market"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	KLType int32 `json:"kLType"`
	KLine
}

// PushOrderBook represents a parsed order book push notification.
type PushOrderBook struct {
	Market int32 `json:"market"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Bids   []OBItem `json:"bids"`
	Asks   []OBItem `json:"asks"`
}

// OBItem represents a single price level in the order book push data.
type OBItem struct {
	Price      float64 `json:"price"`
	Volume     int64 `json:"volume"`
	OrderCount int64 `json:"orderCount"`
}

// PushTicker represents a parsed tick-by-tick push notification.
type PushTicker struct {
	Market    int32   `json:"market"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	Turnover  float64 `json:"turnover"`
	Side      int32   `json:"side"`
	Sequence  int64   `json:"sequence"`
	Dir       int32   `json:"dir"`
	RecvTime  float64 `json:"recvTime"`
	Type      int32   `json:"type"`
	TypeSign  int32   `json:"typeSign"`
}

// PushRT represents a parsed real-time minute data push notification.
type PushRT struct {
	Market    int32   `json:"market"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Time      string  `json:"time"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	AvgPrice  float64 `json:"avgPrice"`
	Turnover  float64 `json:"turnover"`
	Minute    int32   `json:"minute"`
	IsBlank   bool    `json:"isBlank"`
	Timestamp float64 `json:"timestamp"`
}

// PushBroker represents a parsed broker queue push notification.
type PushBroker struct {
	Market int32 `json:"market"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Asks   []BrokerItem `json:"asks"`
	Bids   []BrokerItem `json:"bids"`
}

// BrokerItem represents a single broker queue entry.
type BrokerItem struct {
	Price    float64 `json:"price"`
	Volume   int64   `json:"volume"`
	BrokerID int32   `json:"brokerID"`
	Name     string  `json:"name"`
}

// PushOrderUpdate represents an order status update push.
type PushOrderUpdate struct {
	OrderID     uint64 `json:"orderID"`
	OrderIDEx   string `json:"orderIDEx"`
	Code        string `json:"code"`
	SecMarket   int32 `json:"secMarket"`
	TrdSide     int32 `json:"trdSide"`
	Qty         float64 `json:"qty"`
	Price       float64 `json:"price"`
	OrderStatus int32 `json:"orderStatus"`
}

// PushOrderFill represents an order fill push.
type PushOrderFill struct {
	OrderID        uint64 `json:"orderID"`
	OrderIDEx      string `json:"orderIDEx"`
	Code           string `json:"code"`
	SecMarket      int32 `json:"secMarket"`
	TrdSide        int32 `json:"trdSide"`
	Qty            float64 `json:"qty"`
	Price          float64 `json:"price"`
	FillID         uint64 `json:"fillID"`
	FillIDEx       string `json:"fillIDEx"`
	FillCreateTime string `json:"fillCreateTime"`
}

// PushPriceReminder represents a parsed price reminder push notification.
type PushPriceReminder struct {
	Market       int32 `json:"market"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Price        float64 `json:"price"`
	ChangeRate   float64 `json:"changeRate"`
	MarketStatus int32 `json:"marketStatus"`
	Content      string `json:"content"`
	Note         string `json:"note"`
	Key          int64 `json:"key"`
	Type         int32 `json:"type"`
	SetValue     float64 `json:"setValue"`
	CurValue     float64 `json:"curValue"`
}

// PushTrdNotify represents a parsed trading notification push.
type PushTrdNotify struct {
	AccID     uint64 `json:"accID"`
	TrdEnv    int32 `json:"trdEnv"`
	TrdMarket int32 `json:"trdMarket"`
	Type      int32 `json:"type"`
}

// TradeDate represents a single trade date with its type (full day, half day, etc.).
type TradeDate struct {
	Time          string  `json:"time"`
	Timestamp     float64 `json:"timestamp,omitempty"`
	TradeDateType int32   `json:"tradeDateType,omitempty"`
}

// UsedQuotaInfo represents the quota usage for subscriptions and historical K-line requests.
type UsedQuotaInfo struct {
	UsedSubQuota   int32 `json:"usedSubQuota"`
	UsedKLineQuota int32 `json:"usedKLineQuota"`
}
