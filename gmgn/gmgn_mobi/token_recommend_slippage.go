package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenRecommendSlippageData struct {
	RecommendSlippage string `json:"recommend_slippage"`
	HasTax            bool   `json:"has_tax"`
	DisplaySlippage   string `json:"display_slippage"`
	Volatility        int    `json:"volatility"`
}

type GetTokenRecommendSlippageResp struct {
	Code    int                          `json:"code"`
	Reason  string                       `json:"reason"`
	Message string                       `json:"message"`
	Data    STTokenRecommendSlippageData `json:"data"`
}

type TokenRecommendSlippageTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenRecommendSlippageTool(baseUrl string, baseParam string) *TokenRecommendSlippageTool {
	return &TokenRecommendSlippageTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenRecommendSlippageTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenRecommendSlippageTool) Get(chainType string, tokenAddress string) (*GetTokenRecommendSlippageResp, error) {
	url := "api/v1/recommend_slippage/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenRecommendSlippageResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
