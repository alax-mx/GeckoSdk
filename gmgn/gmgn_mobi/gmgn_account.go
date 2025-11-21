package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STFreshTokenData struct {
	RefreshToken string `json:"refresh_token"`
}

type STData1 struct {
	ExpireAt int    `json:"expire_at"`
	Token    string `json:"token"`
}

type STData0 struct {
	Data STData1 `json:"data"`
	Done bool    `json:"done"`
	Step int     `json:"step"`
}

type FreshTokenResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    STData0 `json:"data"`
}

type AccountTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewAccountTool(baseUrl string, baseParam string) *AccountTool {
	return &AccountTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (gat *AccountTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	gat.proxyInfo = proxyInfo
}

func (gat *AccountTool) FrashAccessToken(authStr string) (*FreshTokenResp, error) {
	url := "account/account/refresh_access_token?" + gat.baseParam
	freshData := &STFreshTokenData{
		RefreshToken: authStr,
	}

	postData, err := json.Marshal(freshData)
	if err != nil {
		return nil, err
	}

	data, err := HttpPost(gat.baseUrl+url, postData, "", gat.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &FreshTokenResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
