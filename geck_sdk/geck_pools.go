package geck_sdk

import (
	"encoding/json"

	"flyu.gecksdk/gecknet"
)

type STAttributes_NetworkPools struct {
	Name                          string                  `json:"name"`
	Address                       string                  `json:"address"`
	BaseTokenPriceUsd             string                  `json:"base_token_price_usd"`
	QuoteTokenPriceUsd            string                  `json:"quote_token_price_usd"`
	BaseTokenPriceNativeCurrency  string                  `json:"base_token_price_native_currency"`
	QuoteTokenPriceNativeCurrency string                  `json:"quote_token_price_native_currency"`
	BaseTokenPriceQuoteToken      string                  `json:"base_token_price_quote_token"`
	QuoteTokenPriceBaseToken      string                  `json:"quote_token_price_base_token"`
	PoolCreatedAt                 string                  `json:"pool_created_at"`
	ReserveInUsd                  string                  `json:"reserve_in_usd"`
	FdvUsd                        string                  `json:"fdv_usd"`
	MarketCapUsd                  string                  `json:"market_cap_usd"`
	PriceChangePercentage         STPriceChangePercentage `json:"price_change_percentage"`
	Transactions                  STTransactions          `json:"transactions"`
	VolumeUSD                     STVolumeUSD             `json:"volume_usd"`
	LockedLiquidityPercentage     string                  `json:"locked_liquidity_percentage"`
}

type STNetworkPoolsData struct {
	ID         string                    `json:"id"`
	Type       string                    `json:"type"`
	Attributes STAttributes_NetworkPools `json:"attributes"`
}

type STNetworkPoolsResp struct {
	Data   STNetworkPoolsData `json:"data"`
	Errors []*STErrors        `json:"errors"`
}

type NetworkPoolsTool struct {
}

func NewNetworkPoolsTool() *NetworkPoolsTool {
	return &NetworkPoolsTool{}
}

func (ndt *NetworkPoolsTool) GetNetworkPools(network string, poolAddress string, include string) (*STNetworkPoolsResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress
	if len(include) > 0 {
		newUrl += "?include=" + include
	}

	data, err := gecknet.HttpGet(newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkPoolsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
