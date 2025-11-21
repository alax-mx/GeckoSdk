package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type GetKolCardsResp struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    STCardsData `json:"data"`
}

type KolCardsTool struct {
	baseUrl   string
	baseParam string
	postData  []byte
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewKolCardsTool(baseUrl string, baseParam string, postData []byte, authStr string) *KolCardsTool {
	return &KolCardsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		postData:  postData,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tdt *KolCardsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tdt *KolCardsTool) Get(chainType string, interval string) (*GetKolCardsResp, error) {
	url := "api/v1/kol_cards/cards/" + chainType + "/" + interval + "?" + tdt.baseParam
	data, err := HttpPost(tdt.baseUrl+url, tdt.postData, tdt.authStr, tdt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetKolCardsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
