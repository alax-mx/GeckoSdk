package token

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STTokenHolderItem struct {
	Address  string `json:"address"`
	Amount   int    `json:"amount"`
	Decimals int    `json:"decimals"`
	Owner    string `json:"owner"`
	Rank     int    `json:"rank"`
}

type STTokenHoldersData struct {
	Total int                  `json:"total"`
	Items []*STTokenHolderItem `json:"items"`
}

type STTokenHoldersResp struct {
	Success bool               `json:"success"`
	Data    STTokenHoldersData `json:"data"`
	Errors  basedef.STErrors   `json:"errors"`
}

type TokenHoldersTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenHoldersTool(solanaInfo *basedef.STSolanaDefine) *TokenHoldersTool {
	return &TokenHoldersTool{
		solanaInfo: solanaInfo,
	}
}

func (tht *TokenHoldersTool) GetTokenHolders(address string, page int, pageSize int, fromAmount string, toAmount string) (*STTokenHoldersResp, error) {
	newUrl := basedef.G_TOKEN_HOLDERS_URL + "address=" + address
	if page > 0 {
		newUrl += "&page=" + strconv.Itoa(page)
	}
	if pageSize > 0 {
		newUrl += "&page_size=" + strconv.Itoa(pageSize)
	}
	if len(fromAmount) > 0 {
		newUrl += "&from_amount=" + fromAmount
	}
	if len(toAmount) > 0 {
		newUrl += "&to_amount=" + toAmount
	}

	data, err := httptool.HttpGet(newUrl, tht.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenHoldersResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
