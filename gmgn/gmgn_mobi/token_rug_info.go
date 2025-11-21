package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STRuggedTokens struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
}

type STTokenRugData struct {
	Address         string           `json:"address"`
	RugRatio        string           `json:"rug_ratio"`
	HolderRuggedNum int              `json:"holder_rugged_num"`
	HolderTokenNum  int              `json:"holder_token_num"`
	RuggedTokens    []STRuggedTokens `json:"rugged_tokens"`
}

type GetTokenRugInfoResp struct {
	Code    int            `json:"code"`
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Data    STTokenRugData `json:"data"`
}

type TokenRugInfoTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenRugInfoTool(baseUrl string, baseParam string, authStr string) *TokenRugInfoTool {
	return &TokenRugInfoTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tit *TokenRugInfoTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tit.proxyInfo = proxyInfo
}

func (tit *TokenRugInfoTool) Get(chainType string, tokenAddress string) (*GetTokenRugInfoResp, error) {
	url := "api/v1/token_rug_info/" + chainType + "/" + tokenAddress + "?" + tit.baseParam
	data, err := HttpGet(tit.baseUrl+url, tit.authStr, tit.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenRugInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
