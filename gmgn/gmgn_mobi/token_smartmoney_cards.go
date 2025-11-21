package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STWallets struct {
	Ens             string   `json:"ens"`
	TwitterUsername string   `json:"twitter_username"`
	TwitterName     string   `json:"twitter_name"`
	Tags            []string `json:"tags"`
	Avatar          string   `json:"avatar"`
	Name            string   `json:"name"`
	Balance         string   `json:"balance"`
	BalanceTs       int      `json:"balance_ts"`
	WalletAddress   string   `json:"wallet_address"`
	NetInflow       string   `json:"net_inflow"`
	NetAmount       string   `json:"net_amount"`
	AmountTotal     string   `json:"amount_total"`
	Buys            int      `json:"buys"`
	Sells           int      `json:"sells"`
	Side            string   `json:"side"`
	IsOpenOrClose   int      `json:"is_open_or_close"`
	Timestamp       int      `json:"timestamp"`
}
type STCards struct {
	Address            string       `json:"address"`
	OpenTimestamp      int          `json:"open_timestamp"`
	TotalSupply        string       `json:"total_supply"`
	Liquidity          string       `json:"liquidity"`
	InitialLiquidity   string       `json:"initial_liquidity"`
	QuoteSymbol        string       `json:"quote_symbol"`
	MarketCap          string       `json:"market_cap"`
	Symbol             string       `json:"symbol"`
	Name               string       `json:"name"`
	DexscrAd           int          `json:"dexscr_ad"`
	DexscrUpdateLink   int          `json:"dexscr_update_link"`
	CtoFlag            int          `json:"cto_flag"`
	HotLevel           int          `json:"hot_level"`
	Launchpad          string       `json:"launchpad"`
	LaunchpadStatus    int          `json:"launchpad_status"`
	Logo               string       `json:"logo"`
	TokenLinks         STTokenLinks `json:"token_links"`
	Price              string       `json:"price"`
	Swaps1H            int          `json:"swaps_1h"`
	Swaps              int          `json:"swaps"`
	Buys               int          `json:"buys"`
	Sells              int          `json:"sells"`
	Volume             string       `json:"volume"`
	HistoryPrice       string       `json:"history_price"`
	TwitterChangeFlag  bool         `json:"twitter_change_flag"`
	TwitterRenameCount int          `json:"twitter_rename_count"`
	HolderCount        int          `json:"holder_count"`
	Security           STSecurity   `json:"security"`
	FirstBuyPriceUsd   float64      `json:"first_buy_price_usd"`
	Wallets            []STWallets  `json:"wallets"`
}

type STCardsData struct {
	Cards []STCards `json:"cards"`
}

type GetSmartMoneyCardsResp struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    STCardsData `json:"data"`
}

type SmartMoneyCardsTool struct {
	baseUrl   string
	baseParam string
	postData  []byte
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewSmartMoneyCardsTool(baseUrl string, baseParam string, postData []byte, authStr string) *SmartMoneyCardsTool {
	return &SmartMoneyCardsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		postData:  postData,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tdt *SmartMoneyCardsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tpt *SmartMoneyCardsTool) SetAuthString(authStr string) {
	tpt.authStr = authStr
}

func (tdt *SmartMoneyCardsTool) Get(chainType string, interval string) (*GetSmartMoneyCardsResp, error) {
	url := "api/v1/smartmoney_cards/cards/" + chainType + "/" + interval + "?" + tdt.baseParam
	data, err := HttpPost(tdt.baseUrl+url, tdt.postData, tdt.authStr, tdt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetSmartMoneyCardsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
