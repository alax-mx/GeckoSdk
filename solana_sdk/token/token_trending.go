package token

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STTokenTrendingData struct {
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type STTokenTrendingResp struct {
	Success  bool                   `json:"success"`
	DataList []*STTokenTrendingData `json:"data"`
	Errors   basedef.STErrors       `json:"errors"`
}

type TokenTrendingTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenTrendingTool(solanaInfo *basedef.STSolanaDefine) *TokenTrendingTool {
	return &TokenTrendingTool{
		solanaInfo: solanaInfo,
	}
}

func (ttt *TokenTrendingTool) GetTokenTrending(limit int) (*STTokenTrendingResp, error) {
	newUrl := basedef.G_TOKEN_TOP_URL + "limit=" + strconv.Itoa(limit)
	data, err := httptool.HttpGet(newUrl, ttt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenTrendingResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
