package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STActivityToken struct {
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	Logo        string `json:"logo"`
	TotalSupply string `json:"total_supply"`
}

type STQuoteToken struct {
	TokenAddress string `json:"token_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Decimals     int    `json:"decimals"`
	Logo         string `json:"logo"`
}

type STActivities struct {
	Wallet        string          `json:"wallet"`
	Chain         string          `json:"chain"`
	TxHash        string          `json:"tx_hash"`
	Timestamp     int             `json:"timestamp"`
	EventType     string          `json:"event_type"`
	Token         STActivityToken `json:"token"`
	TokenAmount   string          `json:"token_amount"`
	QuoteAmount   string          `json:"quote_amount"`
	CostUsd       any             `json:"cost_usd"`
	BuyCostUsd    any             `json:"buy_cost_usd"`
	PriceUsd      any             `json:"price_usd"`
	IsOpenOrClose int             `json:"is_open_or_close"`
	QuoteToken    STQuoteToken    `json:"quote_token"`
	FromAddress   string          `json:"from_address"`
	ToAddress     string          `json:"to_address"`
	Gas           any             `json:"gas"`
}

type STWalletActivityData struct {
	Activities []STActivities `json:"activities"`
	Next       string         `json:"next"`
}

type GetWalletActivityResp struct {
	Code    int                  `json:"code"`
	Reason  string               `json:"reason"`
	Message string               `json:"message"`
	Data    STWalletActivityData `json:"data"`
}

type WalletActivity struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewWalletActivity(baseUrl string, baseParam string, authStr string) *WalletActivity {
	return &WalletActivity{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (wa *WalletActivity) SetProxy(proxyInfo *proxy.STProxyInfo) {
	wa.proxyInfo = proxyInfo
}

func (wa *WalletActivity) Get(chainType string, walletAddress string, limit int) (*GetWalletActivityResp, error) {
	url := "api/v1/wallet_activity/" + chainType + "?" + wa.baseParam
	url += "&wallet=" + walletAddress
	url += "&limit=" + strconv.Itoa(limit)
	url += "&type=buy"
	url += "&type=sell"
	data, err := HttpGet(wa.baseUrl+url, wa.authStr, wa.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetWalletActivityResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
