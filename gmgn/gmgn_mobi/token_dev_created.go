package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokens struct {
	TokenAddress             string `json:"token_address"`
	Symbol                   string `json:"symbol"`
	Chain                    string `json:"chain"`
	CreateTimestamp          int    `json:"create_timestamp"`
	IsOpen                   bool   `json:"is_open"`
	MarcketCap               string `json:"marcket_cap"`
	BiggestPool              string `json:"biggest_pool"`
	PoolLiquidity            string `json:"pool_liquidity"`
	Holders                  int    `json:"holders"`
	Swap1H                   int    `json:"swap_1h"`
	Volume1H                 string `json:"volume_1h"`
	Logo                     string `json:"logo"`
	LiquidityLess4K          bool   `json:"liquidity_less_4k"`
	CtoFlag                  bool   `json:"cto_flag"`
	TwitterNameChangeHistory []any  `json:"twitter_name_change_history"`
	DexscrAd                 bool   `json:"dexscr_ad"`
	DexscrUpdateLink         bool   `json:"dexscr_update_link"`
	IsPump                   bool   `json:"is_pump"`
	LaunchpadPlatform        string `json:"launchpad_platform"`
}

type STTokenDevCreatedData struct {
	LastCreateTimestamp int        `json:"last_create_timestamp"`
	InnerCount          int        `json:"inner_count"`
	OpenCount           int        `json:"open_count"`
	OpenRatio           string     `json:"open_ratio"`
	Tokens              []STTokens `json:"tokens"`
}

type GetTokenDevCreatedResp struct {
	Code    int                   `json:"code"`
	Reason  string                `json:"reason"`
	Message string                `json:"message"`
	Data    STTokenDevCreatedData `json:"data"`
}

type TokenDevCreatedTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenDevCreatedTool(baseUrl string, baseParam string, authStr string) *TokenDevCreatedTool {
	return &TokenDevCreatedTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tdt *TokenDevCreatedTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tdt *TokenDevCreatedTool) SetAuthString(authStr string) {
	tdt.authStr = authStr
}

func (tdt *TokenDevCreatedTool) Get(chainType string, tokenAddress string) (*GetTokenDevCreatedResp, error) {
	url := "api/v1/dev_created_tokens/" + chainType + "/" + tokenAddress + "?" + tdt.baseParam
	data, err := HttpGet(tdt.baseUrl+url, tdt.authStr, tdt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenDevCreatedResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
