package geck_sdk

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/gecknet"
)

type STOhlcv struct {
	Time   int
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type STAttributes_Ohlcv struct {
	OhlcvList []*STOhlcv `json:"ohlcv_list"`
}

type STNetworkOhlcvData struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes STAttributes_Ohlcv `json:"attributes"`
}

type STMetaInfo struct {
	Address         string `json:"address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	CoingeckoCoinID string `json:"coingecko_coin_id"`
}

type STMeta struct {
	Base  STMetaInfo `json:"base"`
	Quote STMetaInfo `json:"quote"`
}

type STNetworkOhlcvResp struct {
	Data   STNetworkPoolsData `json:"data"`
	Meta   STMeta             `json:"meta"`
	Errors []*STErrors        `json:"errors"`
}

type NetworkOhlcvTool struct {
}

func NewNetworkOhlcvTool() *NetworkOhlcvTool {
	return &NetworkOhlcvTool{}
}

func (not *NetworkOhlcvTool) GetNetworkOhlcv(network string, poolAddress string, timeFrame string) (*STNetworkOhlcvResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress + "/ohlcv/" + timeFrame

	data, err := gecknet.HttpGet(newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkOhlcvResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
