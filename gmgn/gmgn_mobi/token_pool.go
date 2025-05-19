package gmgn_mobi

import (
	"encoding/json"
)

type STTokenPoolInfo struct {
	Address             string `json:"address"`
	Exchange            string `json:"exchange"`
	PoolAddress         string `json:"pool_address"`
	QuoteAddress        string `json:"quote_address"`
	QuoteSymbol         string `json:"quote_symbol"`
	Liquidity           string `json:"liquidity"`
	BaseReserve         string `json:"base_reserve"`
	QuoteReserve        string `json:"quote_reserve"`
	InitialLiquidity    string `json:"initial_liquidity"`
	InitialBaseReserve  string `json:"initial_base_reserve"`
	InitialQuoteReserve string `json:"initial_quote_reserve"`
	CreationTimestamp   int    `json:"creation_timestamp"`
	BaseReserveValue    string `json:"base_reserve_value"`
	QuoteReserveValue   string `json:"quote_reserve_value"`
	QuoteVaultAddress   string `json:"quote_vault_address"`
	BaseVaultAddress    string `json:"base_vault_address"`
	Creator             string `json:"creator"`
	FeeRatio            string `json:"fee_ratio"`
}

type GetTokenPoolInfoResp struct {
	Code    int             `json:"code"`
	Reason  string          `json:"reason"`
	Message string          `json:"message"`
	Data    STTokenPoolInfo `json:"data"`
}

type TokenPoolTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenPoolTool(baseUrl string, baseParam string) *TokenPoolTool {
	return &TokenPoolTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tpt *TokenPoolTool) Get(tokenAddress string) (*GetTokenPoolInfoResp, error) {
	url := "api/v1/token_pool_info_sol/sol/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPoolInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
