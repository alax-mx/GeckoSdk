package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenLinkInfo struct {
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

type GetTokenLinkInfoResp struct {
	Code    int             `json:"code"`
	Reason  string          `json:"reason"`
	Message string          `json:"message"`
	Data    STTokenLinkInfo `json:"data"`
}

type TokenLinkTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenLinkTool(baseUrl string, baseParam string) *TokenLinkTool {
	return &TokenLinkTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenLinkTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenLinkTool) Get(chainType string, tokenAddress string) (*GetTokenLinkInfoResp, error) {
	url := "api/v1/token_link/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLinkInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
