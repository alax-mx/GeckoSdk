package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenWalletTagStat struct {
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

type GetTokenWalletTagStatResp struct {
	Code    int                  `json:"code"`
	Reason  string               `json:"reason"`
	Message string               `json:"message"`
	Data    STTokenWalletTagStat `json:"data"`
}

type TokenWalletTagStatTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenWalletTagStatTool(baseUrl string, baseParam string, authStr string) *TokenWalletTagStatTool {
	return &TokenWalletTagStatTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (twtst *TokenWalletTagStatTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	twtst.proxyInfo = proxyInfo
}

func (twtst *TokenWalletTagStatTool) Get(chainType string, tokenAddress string) (*GetTokenWalletTagStatResp, error) {
	url := "api/v1/token_wallet_tags_stat/" + chainType + "/" + tokenAddress + "?" + twtst.baseParam
	data, err := HttpGet(twtst.baseUrl+url, twtst.authStr, twtst.proxyInfo)
	if err != nil {
		return nil, err

	}
	ret := &GetTokenWalletTagStatResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
