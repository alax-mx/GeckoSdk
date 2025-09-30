package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenLaunchpadInfo struct {
	Address           string `json:"address"`
	Launchpad         string `json:"launchpad"`
	LaunchpadStatus   int    `json:"launchpad_status"`
	LaunchpadProgress string `json:"launchpad_progress"`
	Description       string `json:"description"`
	LaunchpadPlatform string `json:"launchpad_platform"`
}

type GetTokenLaunchpadInfoResp struct {
	Code    int                  `json:"code"`
	Reason  string               `json:"reason"`
	Message string               `json:"message"`
	Data    STTokenLaunchpadInfo `json:"data"`
}

type TokenLaunchpadInfoTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenLaunchpadInfoTool(baseUrl string, baseParam string) *TokenLaunchpadInfoTool {
	return &TokenLaunchpadInfoTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenLaunchpadInfoTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenLaunchpadInfoTool) Get(chainType string, tokenAddress string) (*GetTokenLaunchpadInfoResp, error) {
	url := "api/v1/token_launchpad_info/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLaunchpadInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
