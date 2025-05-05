package gecko_sdk

import (
	"encoding/json"
	"strconv"
)

const (
	SORT_TENDING_5M  string = "5m"
	SORT_TENDING_1H  string = "1h"
	SORT_TENDING_6H  string = "6h"
	SORT_TENDING_24H string = "24h"
)

type STAttributes_TrendingPools struct {
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
}

type STTrendingPoolsData struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    STAttributes_TrendingPools `json:"attributes"`
	Relationships STRelationsShips           `json:"relationships"`
}

type STTrendingPoolsResp struct {
	DataList []*STTrendingPoolsData `json:"data"`
	Errors   []*STErrors            `json:"errors"`
}

type TrendingPoolsTool struct {
	apiKey string
}

func NewTrendingPoolsTool(apiKey string) *TrendingPoolsTool {
	return &TrendingPoolsTool{
		apiKey: apiKey,
	}
}

func (ndt *TrendingPoolsTool) GetTrendingPools(network string, include string, page int, duration string) (*STTrendingPoolsResp, error) {
	newUrl := "/networks/trending_pools?"
	if len(network) > 0 {
		newUrl = "/networks/" + network + "/trending_pools?"
	}
	count := 0
	if len(include) > 0 {
		newUrl += "include=" + include
		count++
	}

	if page > 0 {
		if count > 0 {
			newUrl += "&"
		}
		newUrl += "page=" + strconv.Itoa(page)
		count++
	}

	if len(duration) > 0 {
		if count > 0 {
			newUrl += "&"
		}
		newUrl += "duration=" + duration
	}

	data, err := HttpGet(ndt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STTrendingPoolsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
