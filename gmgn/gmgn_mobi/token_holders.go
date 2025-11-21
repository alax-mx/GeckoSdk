package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type List struct {
	Address              string   `json:"address"`
	AccountAddress       string   `json:"account_address"`
	AddrType             int      `json:"addr_type"`
	AmountCur            float64  `json:"amount_cur"`
	UsdValue             float64  `json:"usd_value"`
	CostCur              float64  `json:"cost_cur"`
	SellAmountCur        float64  `json:"sell_amount_cur"`
	SellAmountPercentage float64  `json:"sell_amount_percentage"`
	SellVolumeCur        float64  `json:"sell_volume_cur"`
	BuyVolumeCur         float64  `json:"buy_volume_cur"`
	BuyAmountCur         float64  `json:"buy_amount_cur"`
	NetflowUsd           float64  `json:"netflow_usd"`
	NetflowAmount        float64  `json:"netflow_amount"`
	BuyTxCountCur        float64  `json:"buy_tx_count_cur"`
	SellTxCountCur       float64  `json:"sell_tx_count_cur"`
	WalletTagV2          string   `json:"wallet_tag_v2"`
	NativeBalance        string   `json:"native_balance"`
	Balance              float64  `json:"balance"`
	Profit               float64  `json:"profit"`
	RealizedProfit       float64  `json:"realized_profit"`
	ProfitChange         any      `json:"profit_change"`
	AmountPercentage     float64  `json:"amount_percentage"`
	UnrealizedProfit     float64  `json:"unrealized_profit"`
	UnrealizedPnl        any      `json:"unrealized_pnl"`
	AvgCost              any      `json:"avg_cost"`
	AvgSold              any      `json:"avg_sold"`
	AccuAmount           float64  `json:"accu_amount"`
	AccuCost             float64  `json:"accu_cost"`
	Cost                 float64  `json:"cost"`
	TotalCost            float64  `json:"total_cost"`
	TransferIn           bool     `json:"transfer_in"`
	IsNew                bool     `json:"is_new"`
	IsSuspicious         bool     `json:"is_suspicious"`
	StartHoldingAt       int      `json:"start_holding_at"`
	EndHoldingAt         any      `json:"end_holding_at"`
	LastActiveTimestamp  int      `json:"last_active_timestamp"`
	NativeTransfer       any      `json:"native_transfer"`
	Tags                 []any    `json:"tags"`
	MakerTokenTags       []string `json:"maker_token_tags"`
	Name                 any      `json:"name"`
	Avatar               any      `json:"avatar"`
	TwitterUsername      any      `json:"twitter_username"`
	TwitterName          any      `json:"twitter_name"`
	CreatedAt            int      `json:"created_at"`
}
type STTokenHoldersData struct {
	List []List `json:"list"`
	Next string `json:"next"`
}

type GetTokenHoldersResp struct {
	Code    int                `json:"code"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
	Data    STTokenHoldersData `json:"data"`
}

type TokenHoldersTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenHoldersTool(baseUrl string, baseParam string, authStr string) *TokenHoldersTool {
	return &TokenHoldersTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tht *TokenHoldersTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tht.proxyInfo = proxyInfo
}

func (tdt *TokenHoldersTool) SetAuthString(authStr string) {
	tdt.authStr = authStr
}

func (tht *TokenHoldersTool) Get(chainType string, tokenAddress string, limit int) (*GetTokenHoldersResp, error) {
	url := "vas/api/v1/token_holders/" + chainType + "/" + tokenAddress + "?" + tht.baseParam
	url += "&limit=" + strconv.Itoa(limit)
	data, err := HttpGet(tht.baseUrl+url, tht.authStr, tht.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenHoldersResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
