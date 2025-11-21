package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenPriceInfo struct {
	Address       string `json:"address"`
	Price         string `json:"price"`
	Price1M       string `json:"price_1m"`
	Buys1M        int    `json:"buys_1m"`
	Sells1M       int    `json:"sells_1m"`
	Swaps1M       int    `json:"swaps_1m"`
	BuyVolume1M   string `json:"buy_volume_1m"`
	SellVolume1M  string `json:"sell_volume_1m"`
	Volume1M      string `json:"volume_1m"`
	Price5M       string `json:"price_5m"`
	Buys5M        int    `json:"buys_5m"`
	Sells5M       int    `json:"sells_5m"`
	Swaps5M       int    `json:"swaps_5m"`
	BuyVolume5M   string `json:"buy_volume_5m"`
	SellVolume5M  string `json:"sell_volume_5m"`
	Volume5M      string `json:"volume_5m"`
	Price15M      string `json:"price_15m"`
	Buys15M       int    `json:"buys_15m"`
	Sells15M      int    `json:"sells_15m"`
	Swaps15M      int    `json:"swaps_15m"`
	BuyVolume15M  string `json:"buy_volume_15m"`
	SellVolume15M string `json:"sell_volume_15m"`
	Volume15M     string `json:"volume_15m"`
	Price1H       string `json:"price_1h"`
	Buys1H        int    `json:"buys_1h"`
	Sells1H       int    `json:"sells_1h"`
	Swaps1H       int    `json:"swaps_1h"`
	BuyVolume1H   string `json:"buy_volume_1h"`
	SellVolume1H  string `json:"sell_volume_1h"`
	Volume1H      string `json:"volume_1h"`
	Price6H       string `json:"price_6h"`
	Buys6H        int    `json:"buys_6h"`
	Sells6H       int    `json:"sells_6h"`
	Swaps6H       int    `json:"swaps_6h"`
	BuyVolume6H   string `json:"buy_volume_6h"`
	SellVolume6H  string `json:"sell_volume_6h"`
	Volume6H      string `json:"volume_6h"`
	Price24H      string `json:"price_24h"`
	Buys24H       int    `json:"buys_24h"`
	Sells24H      int    `json:"sells_24h"`
	Swaps24H      int    `json:"swaps_24h"`
	BuyVolume24H  string `json:"buy_volume_24h"`
	SellVolume24H string `json:"sell_volume_24h"`
	Volume24H     string `json:"volume_24h"`
}

type GetTokenPriceInfoResp struct {
	Code    int              `json:"code"`
	Reason  string           `json:"reason"`
	Message string           `json:"message"`
	Data    STTokenPriceInfo `json:"data"`
}

type TokenPriceTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenPriceTool(baseUrl string, baseParam string, authStr string) *TokenPriceTool {
	return &TokenPriceTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tpt *TokenPriceTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenPriceTool) SetAuthString(authStr string) {
	tpt.authStr = authStr
}

func (tpt *TokenPriceTool) Get(chainType string, tokenAddress string) (*GetTokenPriceInfoResp, error) {
	url := "api/v1/token_price_info/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.authStr, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPriceInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
