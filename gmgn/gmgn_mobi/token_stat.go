package gmgn_mobi

import (
	"encoding/json"
)

type STTokenStat struct {
	HolderCount                   int    `json:"holder_count"`
	BluechipOwnerCount            int    `json:"bluechip_owner_count"`
	BluechipOwnerPercentage       string `json:"bluechip_owner_percentage"`
	SignalCount                   int    `json:"signal_count"`
	DegenCallCount                int    `json:"degen_call_count"`
	TopRatTraderPercentage        string `json:"top_rat_trader_percentage"`
	TopBundlerTraderPercentage    string `json:"top_bundler_trader_percentage"`
	TopEntrapmentTraderPercentage string `json:"top_entrapment_trader_percentage"`
}

type GetTokenStatResp struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    STTokenStat `json:"data"`
}

type TokenStatTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenStatTool(baseUrl string, baseParam string) *TokenStatTool {
	return &TokenStatTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tst *TokenStatTool) Get(tokenAddress string) (*GetTokenStatResp, error) {
	url := "api/v1/token_stat/sol/" + tokenAddress + "?" + tst.baseParam
	data, err := HttpGet(tst.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenStatResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
