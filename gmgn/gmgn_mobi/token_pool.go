package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenPoolInfoSol struct {
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

type STTokenPoolInfoEvm struct {
	Address             string `json:"address"`
	PoolAddress         string `json:"pool_address"`
	Exchange            string `json:"exchange"`
	Token0Address       string `json:"token0_address"`
	Token1Address       string `json:"token1_address"`
	BaseAddress         string `json:"base_address"`
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
}

type GetTokenPoolInfoSolResp struct {
	Code    int                `json:"code"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
	Data    STTokenPoolInfoSol `json:"data"`
}

type GetTokenPoolInfoEvmResp struct {
	Code    int                `json:"code"`
	Reason  string             `json:"reason"`
	Message string             `json:"message"`
	Data    STTokenPoolInfoEvm `json:"data"`
}

type TokenPoolTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenPoolTool(baseUrl string, baseParam string) *TokenPoolTool {
	return &TokenPoolTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenPoolTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenPoolTool) GetPoolInfoSol(tokenAddress string) (*GetTokenPoolInfoSolResp, error) {
	url := "api/v1/token_pool_info_sol/sol/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPoolInfoSolResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func (tpt *TokenPoolTool) GetPoolInfoEvm(chainType string, tokenAddress string) (*GetTokenPoolInfoEvmResp, error) {
	url := "api/v1/token_pool_info_evm/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPoolInfoEvmResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
