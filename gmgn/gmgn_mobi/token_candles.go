package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STCandlesData struct {
	Time   int64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type STTokenCandlesInfo struct {
	Time   int64  `json:"time"`
	Open   string `json:"open"`
	Close  string `json:"close"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"volume"`
}

type STTokenCandlesData struct {
	List []STTokenCandlesInfo `json:"list"`
}

type GetTokenCandlesResp struct {
	Code    int                `json:"code"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
	Data    STTokenCandlesData `json:"data"`
}

type TokenCandlesTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenCandlesTool(baseUrl string, baseParam string) *TokenCandlesTool {
	return &TokenCandlesTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tct *TokenCandlesTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tct.proxyInfo = proxyInfo
}

func (tct *TokenCandlesTool) Get(chainType string, tokenAddress string, resolution string, limit int) (*GetTokenCandlesResp, error) {
	url := "api/v1/token_candles/" + chainType + "/" + tokenAddress + "?" + tct.baseParam
	url += "&resolution=" + resolution
	url += "&limit=" + strconv.Itoa(limit)
	data, err := HttpGet(tct.baseUrl+url, "", tct.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenCandlesResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func (not *TokenCandlesTool) ParseOHLCVData(info STTokenCandlesInfo) *STCandlesData {
	candlesData := &STCandlesData{}
	candlesData.Time = info.Time
	candlesData.Open, _ = strconv.ParseFloat(info.Open, 64)
	candlesData.High, _ = strconv.ParseFloat(info.High, 64)
	candlesData.Low, _ = strconv.ParseFloat(info.Low, 64)
	candlesData.Close, _ = strconv.ParseFloat(info.Close, 64)
	candlesData.Volume, _ = strconv.ParseFloat(info.Volume, 64)
	return candlesData
}
