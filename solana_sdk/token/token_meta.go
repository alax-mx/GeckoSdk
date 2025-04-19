package token

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STMetaData struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Twitter     string `json:"twitter"`
	Website     string `json:"website"`
}

type STTokenMetaData struct {
	Address         string     `json:"address"`
	Name            string     `json:"name"`
	Symbol          string     `json:"symbol"`
	Icon            string     `json:"icon"`
	Decimals        int        `json:"decimals"`
	Holder          int        `json:"holder"`
	Creator         string     `json:"creator"`
	CreateTX        string     `json:"create_tx"`
	CreateTime      int        `json:"created_time"`
	MetaData        STMetaData `json:"metadata"`
	MintAuthority   string     `json:"mint_authority"`
	FreezeAuthority string     `json:"freeze_authority"`
	Supply          string     `json:"supply"`
	Price           float64    `json:"price"`
	Volume24H       int        `json:"volume_24h"`
	MarketCap       float64    `json:"market_cap"`
	MarketCapRank   int        `json:"market_cap_rank"`
	PriceChange24H  float64    `json:"price_change_24h"`
}

type STTokenMetaResp struct {
	Success bool             `json:"success"`
	Data    STTokenMetaData  `json:"data"`
	Errors  basedef.STErrors `json:"errors"`
}

type TokenMetaTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenMetaTool(solanaInfo *basedef.STSolanaDefine) *TokenMetaTool {
	return &TokenMetaTool{
		solanaInfo: solanaInfo,
	}
}

func (tmt *TokenMetaTool) GetTokenMeta(address string) (*STTokenMetaResp, error) {
	newUrl := basedef.G_TOKEN_META_URL + "address=" + address
	data, err := httptool.HttpGet(newUrl, tmt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenMetaResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
