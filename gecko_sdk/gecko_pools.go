package gecko_sdk

import (
	"encoding/json"
	"errors"
	"strconv"
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
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	Attributes    STAttributes_NetworkPools `json:"attributes"`
	RelationShips STRelationsShips          `json:"relationships"`
}

type STNetworkPoolsResp struct {
	Data     STNetworkPoolsData `json:"data"`
	Included []*STIncluded      `json:"included"`
	Errors   []*STErrors        `json:"errors"`
}

type STNetworkNewPoolsResp struct {
	Data     []*STNetworkPoolsData `json:"data"`
	Included []*STIncluded         `json:"included"`
	Errors   []*STErrors           `json:"errors"`
}

type STNetworkTopPoolsResp struct {
	Data     []*STNetworkPoolsData `json:"data"`
	Included []*STIncluded         `json:"included"`
	Errors   []*STErrors           `json:"errors"`
}

type NetworkPoolsTool struct {
	apiKey string
}

func NewNetworkPoolsTool(apiKey string) *NetworkPoolsTool {
	return &NetworkPoolsTool{
		apiKey: apiKey,
	}
}

func (ndt *NetworkPoolsTool) GetNetworkPools(network string, poolAddress string, include string) (*STNetworkPoolsResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress
	if len(include) > 0 {
		newUrl += "?include=" + include
	}

	data, err := HttpGet(ndt.apiKey, newUrl)
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

func (ndt *NetworkPoolsTool) GetNetworkMultisPools(network string, poolAddress []string, include string) (*STNetworkPoolsResp, error) {
	newUrl := "/networks/" + network + "/pools/multi/"
	if len(poolAddress) <= 0 {
		return nil, errors.New("err: GetNetworkMultisPools len(poolAddress) <= 0")
	}

	for i := 0; i < len(poolAddress); i++ {
		if i > 0 {
			newUrl += ","
		}
		newUrl += poolAddress[i]
	}

	if len(include) > 0 {
		newUrl += "?include=" + include
	}

	data, err := HttpGet(ndt.apiKey, newUrl)
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

func (ndt *NetworkPoolsTool) GetNetworkTopPools(network string, include string, page int, sortBy string) (*STNetworkTopPoolsResp, error) {
	newUrl := "/networks/" + network + "/pools"
	count := 0
	if len(include) > 0 {
		newUrl += "?include=" + include
		count++
	}

	if page > 0 {
		if count == 0 {
			newUrl += "?page=" + strconv.Itoa(page)
		} else {
			newUrl += "&page=" + strconv.Itoa(page)
		}
		count++
	}

	if len(sortBy) > 0 {
		if count == 0 {
			newUrl += "?sort=" + sortBy
		} else {
			newUrl += "&sort=" + sortBy
		}
	}

	data, err := HttpGet(ndt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkTopPoolsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ndt *NetworkPoolsTool) GetNetworkNewPools(network string, include string, page int) (*STNetworkNewPoolsResp, error) {
	newUrl := "/networks/" + network + "/new_pools"
	count := 0
	if len(include) > 0 {
		newUrl += "?include=" + include
		count++
	}

	if page > 0 {
		if count == 0 {
			newUrl += "?page=" + strconv.Itoa(page)
		} else {
			newUrl += "&page=" + strconv.Itoa(page)
		}
	}

	data, err := HttpGet(ndt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkNewPoolsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
