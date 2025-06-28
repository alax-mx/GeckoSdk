package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenLinks struct {
	Address         string `json:"address"`
	Gmgn            string `json:"gmgn"`
	Geckoterminal   string `json:"geckoterminal"`
	TwitterUsername string `json:"twitter_username"`
	Website         string `json:"website"`
	Telegram        string `json:"telegram"`
	Bitbucket       string `json:"bitbucket"`
	Discord         string `json:"discord"`
	Description     string `json:"description"`
	Facebook        string `json:"facebook"`
	Github          string `json:"github"`
	Instagram       string `json:"instagram"`
	Linkedin        string `json:"linkedin"`
	Medium          string `json:"medium"`
	Reddit          string `json:"reddit"`
	Tiktok          string `json:"tiktok"`
	Youtube         string `json:"youtube"`
	VerifyStatus    int    `json:"verify_status"`
}
type STSecurityInfo struct {
	BurnRatio              string `json:"burn_ratio"`
	BurnStatus             string `json:"burn_status"`
	DevTokenBurnAmount     string `json:"dev_token_burn_amount"`
	DevTokenBurnRatio      string `json:"dev_token_burn_ratio"`
	IsShowAlert            bool   `json:"is_show_alert"`
	RenouncedFreezeAccount int    `json:"renounced_freeze_account"`
	RenouncedMint          int    `json:"renounced_mint"`
	Top10HolderRate        string `json:"top_10_holder_rate"`
}
type STBluchipRankData struct {
	Chain                string         `json:"chain"`
	Address              string         `json:"address"`
	Symbol               string         `json:"symbol"`
	Creator              string         `json:"creator"`
	CreatorTokenStatus   string         `json:"creator_token_status"`
	Price                string         `json:"price"`
	PriceChangePercent1M string         `json:"price_change_percent_1m"`
	PriceChangePercent5M string         `json:"price_change_percent_5m"`
	PriceChangePercent1H string         `json:"price_change_percent_1h"`
	CtoFlag              int            `json:"cto_flag"`
	DexscrAd             int            `json:"dexscr_ad"`
	DexscrUpdateLink     int            `json:"dexscr_update_link"`
	HolderCount          int            `json:"holder_count"`
	HotLevel             int            `json:"hot_level"`
	IsShowAlert          bool           `json:"is_show_alert"`
	Launchpad            string         `json:"launchpad"`
	LaunchpadStatus      int            `json:"launchpad_status"`
	Liquidity            string         `json:"liquidity"`
	Logo                 string         `json:"logo"`
	AtLeastOneLink       bool           `json:"at_least_one_link"`
	TokenLinks           STTokenLinks   `json:"token_links"`
	MarketCap            string         `json:"market_cap"`
	OpenTimestamp        int            `json:"open_timestamp"`
	BluechipRate         string         `json:"bluechip_rate"`
	Buys                 int            `json:"buys"`
	SmartBuys24H         int            `json:"smart_buys_24h"`
	SmartSells24H        int            `json:"smart_sells_24h"`
	Sells                int            `json:"sells"`
	Swaps                int            `json:"swaps"`
	Volume               string         `json:"volume"`
	RatTraderAmountRate  string         `json:"rat_trader_amount_rate"`
	SecurityInfo         STSecurityInfo `json:"security_info"`
	TwitterChangeFlag    int            `json:"twitter_change_flag"`
	TwitterRenameCount   int            `json:"twitter_rename_count"`
	IsWashTrading        bool           `json:"is_wash_trading"`
}

type GetTokenBluchipRankResp struct {
	Code    int                 `json:"code"`
	Reason  string              `json:"reason"`
	Message string              `json:"message"`
	Data    []STBluchipRankData `json:"data"`
}

type TokenBluchipRankTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenBluchipRankTool(baseUrl string, baseParam string) *TokenBluchipRankTool {
	return &TokenBluchipRankTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tbrt *TokenBluchipRankTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tbrt.proxyInfo = proxyInfo
}

func (tbrt *TokenBluchipRankTool) Get(interval string, limit int) (*GetTokenBluchipRankResp, error) {
	url := "api/v1/bluechip_rank/sol?"
	url += "interval=" + interval
	url += "&" + tbrt.baseParam
	url += "&limit=" + strconv.Itoa(limit)
	url += "&order_by=marketcap"
	url += "&orderby=marketcap"
	url += "&direction=desc"
	url += "&filters=not_wash_trading"
	url += "&filters=renounced"
	url += "&filters=frozen"
	data, err := HttpGet(tbrt.baseUrl+url, tbrt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenBluchipRankResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
