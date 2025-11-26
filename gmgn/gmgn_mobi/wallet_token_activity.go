package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STWalletTokenActivityData struct {
	Activities []STActivities `json:"activities"`
	Next       any            `json:"next"`
}

type GetWalletTokenActivityResp struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data STWalletTokenActivityData `json:"data"`
}

type WalletTokenActivity struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewWalletTokenActivity(baseUrl string, baseParam string) *WalletTokenActivity {
	return &WalletTokenActivity{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (wta *WalletTokenActivity) SetProxy(proxyInfo *proxy.STProxyInfo) {
	wta.proxyInfo = proxyInfo
}

func (wa *WalletTokenActivity) Get(chainType string, walletAddress string, tokenAddress string, limit int) (*GetWalletTokenActivityResp, error) {
	url := "defi/quotation/v1/wallet_token_activity/" + chainType + "?" + wa.baseParam
	url += "&wallet=" + walletAddress
	url += "&token=" + tokenAddress
	url += "&limit=" + strconv.Itoa(limit)
	data, err := HttpGet(wa.baseUrl+url, "", wa.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetWalletTokenActivityResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
