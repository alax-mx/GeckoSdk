package token

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STTokenTopItem struct {
	Address        string  `json:"address"`
	Decimals       int     `json:"decimals"`
	Name           string  `json:"name"`
	Symbol         string  `json:"symbol"`
	MarketCap      int     `json:"market_cap"`
	Price          float64 `json:"price"`
	Price24HChange float64 `json:"price_24h_change"`
	Holder         int     `json:"holder"`
	CreatedTime    int     `json:"created_time"`
}

type STTokenTopData struct {
	Total int               `json:"total"`
	Items []*STTokenTopItem `json:"items"`
}

type STTokenTopResp struct {
	Success bool             `json:"success"`
	Data    STTokenTopData   `json:"data"`
	Errors  basedef.STErrors `json:"errors"`
}

type TokenTopTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenTopTool(solanaInfo *basedef.STSolanaDefine) *TokenTopTool {
	return &TokenTopTool{
		solanaInfo: solanaInfo,
	}
}

func (ttt *TokenTopTool) GetTokenTop() (*STTokenTopResp, error) {
	data, err := httptool.HttpGet(basedef.G_TOKEN_TOP_URL, ttt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenTopResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
