package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenDevInfo struct {
	Address                  string `json:"address"`
	CreatorAddress           string `json:"creator_address"`
	CreatorTokenBalance      string `json:"creator_token_balance"`
	CreatorTokenStatus       string `json:"creator_token_status"`
	TwitterNameChangeHistory []any  `json:"twitter_name_change_history"`
	Top10HolderRate          string `json:"top_10_holder_rate"`
	DexscrAd                 int    `json:"dexscr_ad"`
	DexscrUpdateLink         int    `json:"dexscr_update_link"`
	CtoFlag                  int    `json:"cto_flag"`
}

type GetTokenDevInfoResp struct {
	Code    int            `json:"code"`
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Data    STTokenDevInfo `json:"data"`
}

type TokenDevTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenDevTool(baseUrl string, baseParam string) *TokenDevTool {
	return &TokenDevTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tdt *TokenDevTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tdt *TokenDevTool) Get(chainType string, tokenAddress string) (*GetTokenDevInfoResp, error) {
	url := "api/v1/token_dev_info/" + chainType + "/" + tokenAddress + "?" + tdt.baseParam
	data, err := HttpGet(tdt.baseUrl+url, tdt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenDevInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
