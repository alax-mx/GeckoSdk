package gecko_sdk

import (
	"encoding/json"
	"errors"
)

type STAttributes_Tokens struct {
	Name              string      `json:"name"`
	Address           string      `json:"address"`
	Symbol            string      `json:"symbol"`
	ImageURL          string      `json:"image_url"`
	Decimals          int         `json:"decimals"`
	TotalSupply       string      `json:"total_supply"`
	CoinGeckoCoinID   string      `json:"coingecko_coin_id"`
	PriceUsd          string      `json:"price_usd"`
	FdvUsd            string      `json:"fdv_usd"`
	TotalReserveInUsd string      `json:"total_reserve_in_usd"`
	VolumeUSD         STVolumeUSD `json:"volume_usd"`
	MarketCapUsd      string      `json:"market_cap_usd"`
}

type STAttributes_TokenInfo struct {
	Name            string            `json:"name"`
	Address         string            `json:"address"`
	Symbol          string            `json:"symbol"`
	Decimals        int               `json:"decimals"`
	CoinGeckoCoinID string            `json:"coingecko_coin_id"`
	ImageURL        string            `json:"image_url"`
	WebSites        []string          `json:"websites"`
	Description     string            `json:"description"`
	DiscordURL      string            `json:"discord_url"`
	TelegramHandle  string            `json:"telegram_handle"`
	TwitterHandle   string            `json:"twitter_handle"`
	Categories      []string          `json:"categories"`
	GTCategoryIds   []string          `json:"gt_category_ids"`
	GTScore         float64           `json:"gt_score"`
	GTScoreDetails  *STGTScoreDetails `json:"gt_score_details"`
	Holders         STHolders         `json:"holders"`
	MintAuthority   string            `json:"mint_authority"`
	FreezeAuthority string            `json:"freeze_authority"`
}

type STAttributes_TokenInfoRecentlyUpdated struct {
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	Symbol            string   `json:"symbol"`
	Decimals          int      `json:"decimals"`
	CoinGeckoCoinID   string   `json:"coingecko_coin_id"`
	ImageURL          string   `json:"image_url"`
	WebSites          []string `json:"websites"`
	Description       string   `json:"description"`
	DiscordURL        string   `json:"discord_url"`
	TelegramHandle    string   `json:"telegram_handle"`
	TwitterHandle     string   `json:"twitter_handle"`
	GTScore           float64  `json:"gt_score"`
	MetadataUpdatedAt string   `json:"metadata_updated_at"`
}

type STNetworkTokenData struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes STAttributes_Tokens `json:"attributes"`
}

type STNetworkTokenInfoData struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes STAttributes_TokenInfo `json:"attributes"`
}

type STNetworkTokenInfoRecentlyUpdatedData struct {
	ID         string                                `json:"id"`
	Type       string                                `json:"type"`
	Attributes STAttributes_TokenInfoRecentlyUpdated `json:"attributes"`
}

type STNetworkTokensResp struct {
	Data     STNetworkTokenData    `json:"data"`
	Included []*STNetworkPoolsData `json:"included"`
	Errors   []*STErrors           `json:"errors"`
}

type STNetworkMultiTokensResp struct {
	Data     []*STNetworkTokenData `json:"data"`
	Included []*STNetworkPoolsData `json:"included"`
	Errors   []*STErrors           `json:"errors"`
}

type STNetworkTokenInfoResp struct {
	Data   STNetworkTokenInfoData `json:"data"`
	Errors []*STErrors            `json:"errors"`
}

type STNetworkTokenTopPoolsResp struct {
	Data   []*STNetworkPoolsData `json:"data"`
	Errors []*STErrors           `json:"errors"`
}

type STNetworkPoolTokensInfoResp struct {
	Data   []*STNetworkTokenInfoData `json:"data"`
	Errors []*STErrors               `json:"errors"`
}

type STRecentlyUpdatedTokensInfoResp struct {
	Data   []*STNetworkTokenInfoRecentlyUpdatedData `json:"data"`
	Errors []*STErrors                              `json:"errors"`
}

type NetworkTokensTool struct {
	apiKey string
}

func NewNetworkTokensTool(apiKey string) *NetworkTokensTool {
	return &NetworkTokensTool{
		apiKey: apiKey,
	}
}

func (ntt *NetworkTokensTool) GetNetworkTokens(network string, tokenAddress string, include string) (*STNetworkTokensResp, error) {
	newUrl := "/networks/" + network + "/tokens/" + tokenAddress
	if len(include) > 0 {
		newUrl += "?include=" + include
	}

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkTokensResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ntt *NetworkTokensTool) GetNetworkMultiTokens(network string, tokenAddress []string, include string) (*STNetworkMultiTokensResp, error) {
	newUrl := "/networks/" + network + "/tokens/multi/"
	if len(tokenAddress) <= 0 {
		return nil, errors.New("err: GetNetworkMultiTokens len(tokenAddress) <= 0")
	}

	for i := 0; i < len(tokenAddress); i++ {
		if i > 0 {
			newUrl += ","
		}
		newUrl += tokenAddress[i]
	}

	if len(include) > 0 {
		newUrl += "?include=" + include
	}

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkMultiTokensResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ntt *NetworkTokensTool) GetNetworkTokenInfo(network string, tokenAddress string) (*STNetworkTokenInfoResp, error) {
	newUrl := "/networks/" + network + "/tokens/" + tokenAddress + "/info"

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkTokenInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ntt *NetworkTokensTool) GetNetworkTokenTopPools(network string, tokenAddress string) (*STNetworkTokenTopPoolsResp, error) {
	newUrl := "/networks/" + network + "/tokens/" + tokenAddress + "/pools"

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkTokenTopPoolsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ntt *NetworkTokensTool) GetNetworkPoolTokensInfo(network string, poolAddress string) (*STNetworkPoolTokensInfoResp, error) {
	newUrl := "/networks/" + network + "/pools/" + poolAddress + "/info"

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STNetworkPoolTokensInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (ntt *NetworkTokensTool) GetRecentlyUpdateTokens(network string, include string) (*STRecentlyUpdatedTokensInfoResp, error) {
	newUrl := "/tokens/info_recently_updated"
	count := 0
	if len(network) > 0 {
		newUrl += "?network=" + network
		count++
	}

	if len(include) > 0 {
		if count <= 0 {
			newUrl += "?include=" + include
		} else {
			newUrl += "&include=" + include
		}
	}

	data, err := HttpGet(ntt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STRecentlyUpdatedTokensInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
