package gecko_sdk

import (
	"encoding/json"
)

type STOhlcv struct {
	Time   int64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type STAttributes_Ohlcv struct {
	OhlcvList [][]float64 `json:"ohlcv_list"`
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
	Data   STNetworkOhlcvData `json:"data"`
	Meta   STMeta             `json:"meta"`
	Errors []*STErrors        `json:"errors"`
}

type NetworkOhlcvTool struct {
	apiKey string
}

func NewNetworkOhlcvTool(apiKey string) *NetworkOhlcvTool {
	return &NetworkOhlcvTool{
		apiKey: apiKey,
	}
}

func (not *NetworkOhlcvTool) GetNetworkOhlcv(network string, poolAddress string,
	timeFrame string, aggregate string, token string) (*STNetworkOhlcvResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress + "/ohlcv/" + timeFrame

	count := 0
	if len(aggregate) > 0 {
		newUrl += "?aggregate=" + aggregate
		count++
	}

	if len(token) > 0 {
		if count <= 0 {
			newUrl += "?token=" + token
		} else {
			newUrl += "&token=" + token
		}
	}

	data, err := HttpGet(not.apiKey, newUrl)
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

func (not *NetworkOhlcvTool) ParseOHLCVData(data [][]float64) []*STOhlcv {
	retList := make([]*STOhlcv, 0)
	count := len(data)
	for i := 0; i < count; i++ {
		infoDataList := data[i]
		if len(infoDataList) != 6 {
			return nil
		}

		ohlcvInfo := &STOhlcv{}
		ohlcvInfo.Time = int64(infoDataList[0])
		ohlcvInfo.Open = infoDataList[1]
		ohlcvInfo.High = infoDataList[2]
		ohlcvInfo.Low = infoDataList[3]
		ohlcvInfo.Close = infoDataList[4]
		ohlcvInfo.Volume = infoDataList[5]
		retList = append(retList, ohlcvInfo)
	}
	return retList
}
