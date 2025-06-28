package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenHolderStatData struct {
	SmartDegenCount    int `json:"smart_degen_count"`
	RenownedCount      int `json:"renowned_count"`
	FreshWalletCount   int `json:"fresh_wallet_count"`
	DexBotCount        int `json:"dex_bot_count"`
	InsiderCount       int `json:"insider_count"`
	FollowingCount     int `json:"following_count"`
	DevCount           int `json:"dev_count"`
	BluechipOwnerCount int `json:"bluechip_owner_count"`
	BundlerCount       int `json:"bundler_count"`
}

type GetTokenHolderStatResp struct {
	Code    int                   `json:"code"`
	Reason  string                `json:"reason"`
	Message string                `json:"message"`
	Data    STTokenHolderStatData `json:"data"`
}

type TokenHolderStatTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenHolderStatTool(baseUrl string, baseParam string) *TokenHolderStatTool {
	return &TokenHolderStatTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (thst *TokenHolderStatTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	thst.proxyInfo = proxyInfo
}

func (thst *TokenHolderStatTool) Get(tokenAddress string) (*GetTokenHolderStatResp, error) {
	url := "vas/api/v1/token_holder_stat/sol/" + tokenAddress + "?" + thst.baseParam
	data, err := HttpGet(thst.baseUrl+url, thst.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenHolderStatResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
