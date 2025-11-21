package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenStat struct {
	HolderCount                   int    `json:"holder_count"`
	BluechipOwnerCount            int    `json:"bluechip_owner_count"`
	BluechipOwnerPercentage       string `json:"bluechip_owner_percentage"`
	SignalCount                   int    `json:"signal_count"`
	DegenCallCount                int    `json:"degen_call_count"`
	TopRatTraderPercentage        string `json:"top_rat_trader_percentage"`
	TopBundlerTraderPercentage    string `json:"top_bundler_trader_percentage"`
	TopEntrapmentTraderPercentage string `json:"top_entrapment_trader_percentage"`
}

type GetTokenStatResp struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    STTokenStat `json:"data"`
}

type TokenStatTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenStatTool(baseUrl string, baseParam string, authStr string) *TokenStatTool {
	return &TokenStatTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tst *TokenStatTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tst.proxyInfo = proxyInfo
}

func (tst *TokenStatTool) Get(chainType string, tokenAddress string) (*GetTokenStatResp, error) {
	url := "api/v1/token_stat/" + chainType + "/" + tokenAddress + "?" + tst.baseParam
	data, err := HttpGet(tst.baseUrl+url, tst.authStr, tst.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenStatResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
