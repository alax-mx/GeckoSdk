package token

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STTokenPriceData struct {
	Date  int     `json:"date"`
	Price float64 `json:"price"`
}

type STTokenPriceResp struct {
	Success  bool                `json:"success"`
	DataList []*STTokenPriceData `json:"data"`
	Errors   basedef.STErrors    `json:"errors"`
}

type TokenPriceTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenPriceTool(solanaInfo *basedef.STSolanaDefine) *TokenPriceTool {
	return &TokenPriceTool{
		solanaInfo: solanaInfo,
	}
}

func (tpt *TokenPriceTool) GetTokenPrice(address string, fromTime int, toTime int) (*STTokenPriceResp, error) {
	newUrl := basedef.G_TOKEN_PRICE_URL + "address=" + address
	if fromTime > 0 && toTime > 0 {
		newUrl += "&from_time=" + strconv.Itoa(fromTime) + "&to_time=" + strconv.Itoa(toTime)
	}

	data, err := httptool.HttpGet(newUrl, tpt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenPriceResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
