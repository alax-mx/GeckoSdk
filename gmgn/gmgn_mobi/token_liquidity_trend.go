package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenLiquidityData struct {
	PoolSize string `json:"pool_size"`
	Pools    string `json:"pools"`
}

type GetTokenLiquidityTrendResp struct {
	Code    int                  `json:"code"`
	Reason  string               `json:"reason"`
	Message string               `json:"message"`
	Data    STTokenLiquidityData `json:"data"`
}

type TokenLiquidityTrendTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenLiquidityTrendTool(baseUrl string, baseParam string) *TokenLiquidityTrendTool {
	return &TokenLiquidityTrendTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenLiquidityTrendTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenLiquidityTrendTool) Get(tokenAddress string) (*GetTokenLiquidityTrendResp, error) {
	url := "vas/api/v1/token_liquidity_trend/sol/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLiquidityTrendResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
