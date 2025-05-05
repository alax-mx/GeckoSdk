package geck_sdk

import (
	"encoding/json"
	"strconv"
)

type STAttributes_Trades struct {
	BlockNumber              int    `json:"block_number"`
	BlockTimestamp           string `json:"block_timestamp"`
	TxHash                   string `json:"tx_hash"`
	TxRromAddress            string `json:"tx_from_address"`
	FromTokenAmount          string `json:"from_token_amount"`
	ToTokenAmount            string `json:"to_token_amount"`
	PriceFromInCurrencyToken string `json:"price_from_in_currency_token"`
	PriceToInCurrencyToken   string `json:"price_to_in_currency_token"`
	PriceFromInUsd           string `json:"price_from_in_usd"`
	PriceToInUsd             string `json:"price_to_in_usd"`
	Kind                     string `json:"kind"`
	VolumeInUsd              string `json:"volume_in_usd"`
	FromTokenAddress         string `json:"from_token_address"`
	ToTokenAddress           string `json:"to_token_address"`
}

type STNetworkTradesData struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes STAttributes_Trades `json:"attributes"`
}

type STNetworkTradesResp struct {
	Data   []*STNetworkTradesData `json:"data"`
	Errors []*STErrors            `json:"errors"`
}

type NetworkTradesTool struct {
}

func NewNetworkTradesTool() *NetworkTradesTool {
	return &NetworkTradesTool{}
}

func (ntt *NetworkTradesTool) GetNetworkTrades(network string, poolAddress string, volumeThan int, token string) (*STNetworkTradesResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress + "/trades"
	count := 0
	if volumeThan > 0 {
		newUrl += "?trade_volume_in_usd_greater_than=" + strconv.Itoa(volumeThan)
		count++
	}

	if len(token) > 0 {
		if count <= 0 {
			newUrl += "?token=" + token
		} else {
			newUrl += "&token=" + token
		}
	}

	data, err := HttpGet(newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkTradesResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
