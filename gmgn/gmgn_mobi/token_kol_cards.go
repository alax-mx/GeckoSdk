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
	proxyInfo *proxy.STProxyInfo
}

func NewKolCardsTool(baseUrl string, baseParam string, postData []byte) *KolCardsTool {
	return &KolCardsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		postData:  postData,
		proxyInfo: nil,
	}
}

func (tdt *KolCardsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tdt *KolCardsTool) Get(chainType string, interval string) (*GetKolCardsResp, error) {
	url := "api/v1/kol_cards/cards/" + chainType + "/" + interval + "?" + tdt.baseParam
	data, err := HttpPost(tdt.baseUrl+url, tdt.postData, tdt.proxyInfo)
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
