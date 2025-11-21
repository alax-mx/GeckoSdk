package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STToken struct {
	Address       string `json:"address"`
	TokenAddress  string `json:"token_address"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	Decimals      int    `json:"decimals"`
	Logo          string `json:"logo"`
	PriceChange6H string `json:"price_change_6h"`
	IsShowAlert   bool   `json:"is_show_alert"`
	IsHoneypot    any    `json:"is_honeypot"`
}
type STHoldingsData struct {
	Token               STToken  `json:"token"`
	Balance             string   `json:"balance"`
	UsdValue            string   `json:"usd_value"`
	RealizedProfit30D   string   `json:"realized_profit_30d"`
	RealizedProfit      string   `json:"realized_profit"`
	RealizedPnl         string   `json:"realized_pnl"`
	RealizedPnl30D      string   `json:"realized_pnl_30d"`
	UnrealizedProfit    string   `json:"unrealized_profit"`
	UnrealizedPnl       string   `json:"unrealized_pnl"`
	TotalProfit         string   `json:"total_profit"`
	TotalProfitPnl      string   `json:"total_profit_pnl"`
	AvgCost             string   `json:"avg_cost"`
	AvgSold             string   `json:"avg_sold"`
	Buy30D              int      `json:"buy_30d"`
	Sell30D             int      `json:"sell_30d"`
	Sells               int      `json:"sells"`
	Price               string   `json:"price"`
	Cost                string   `json:"cost"`
	PositionPercent     string   `json:"position_percent"`
	LastActiveTimestamp int      `json:"last_active_timestamp"`
	HistorySoldIncome   string   `json:"history_sold_income"`
	HistoryBoughtCost   string   `json:"history_bought_cost"`
	StartHoldingAt      any      `json:"start_holding_at"`
	EndHoldingAt        any      `json:"end_holding_at"`
	Liquidity           string   `json:"liquidity"`
	TotalSupply         string   `json:"total_supply"`
	WalletTokenTags     []string `json:"wallet_token_tags"`
	LastBlock           int      `json:"last_block"`
}

type STData struct {
	Holdings []STHoldingsData `json:"holdings"`
	Next     string           `json:"next"`
}

type GetWalletHoldingsResp struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Data    STData `json:"data"`
}

type WalletHoldingsTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewWalletHoldingsTool(baseUrl string, baseParam string, authStr string) *WalletHoldingsTool {
	return &WalletHoldingsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (wht *WalletHoldingsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	wht.proxyInfo = proxyInfo
}

func (tpt *WalletHoldingsTool) SetAuthString(authStr string) {
	tpt.authStr = authStr
}

func (wht *WalletHoldingsTool) Get(chainType string, walletAddress string, limit int, orderBy string,
	direction string, showSmall bool, sellOout bool, hideABNormal bool) (*GetWalletHoldingsResp, error) {
	url := "api/v1/wallet_holdings/" + chainType + "/" + walletAddress + "?" + wht.baseParam
	url += "&limit=" + strconv.Itoa(limit)
	url += "&orderby=" + orderBy
	url += "&direction=" + direction
	if showSmall {
		url += "&showsmall=true"
	} else {
		url += "&showsmall=false"
	}
	if sellOout {
		url += "&sellout=true"
	} else {
		url += "&sellout=false"
	}
	if hideABNormal {
		url += "&hide_abnormal=true"
	} else {
		url += "&hide_abnormal=false"
	}
	data, err := HttpGet(wht.baseUrl+url, wht.authStr, wht.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetWalletHoldingsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
