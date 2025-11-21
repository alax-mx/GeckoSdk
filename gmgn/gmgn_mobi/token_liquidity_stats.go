package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenLiquidityStatsData struct {
	Chain            string `json:"chain"`
	TokenAddress     string `json:"token_address"`
	SmartWallets     int    `json:"smart_wallets"`
	FreshWallets     int    `json:"fresh_wallets"`
	RenownedWallets  int    `json:"renowned_wallets"`
	CreatorWallets   int    `json:"creator_wallets"`
	SniperWallets    int    `json:"sniper_wallets"`
	RatTraderWallets int    `json:"rat_trader_wallets"`
	WhaleWallets     int    `json:"whale_wallets"`
	TopWallets       int    `json:"top_wallets"`
	FollowingWallets int    `json:"following_wallets"`
	BundlerWallets   int    `json:"bundler_wallets"`
}

type GetTokenLiquidityStatsResp struct {
	Code    int                       `json:"code"`
	Reason  string                    `json:"reason"`
	Message string                    `json:"message"`
	Data    STTokenLiquidityStatsData `json:"data"`
}

type TokenLiquidityStatsTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenLiquidityStatsTool(baseUrl string, baseParam string, authStr string) *TokenLiquidityStatsTool {
	return &TokenLiquidityStatsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tpt *TokenLiquidityStatsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tdt *TokenLiquidityStatsTool) SetAuthString(authStr string) {
	tdt.authStr = authStr
}

func (tpt *TokenLiquidityStatsTool) Get(chainType string, tokenAddress string) (*GetTokenLiquidityStatsResp, error) {
	url := "api/v1/token_liquidity_stats/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.authStr, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLiquidityStatsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
