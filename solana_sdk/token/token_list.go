package token

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

const (
	TOKEN_LIST_SORT_BY_HOLDER       string = "holder"
	TOKEN_LIST_SORT_BY_MARKET_CAP   string = "market_cap"
	TOKEN_LIST_SORT_BY_CREATED_TIME string = "created_time"
)

const (
	TOKEN_LIST_SORT_ORDER_ASC  string = "asc"
	TOKEN_LIST_SORT_ORDER_DESC string = "desc"
)

const (
	TOKEN_PAGE_SIZE_10  int = 10
	TOKEN_PAGE_SIZE_20  int = 20
	TOKEN_PAGE_SIZE_30  int = 30
	TOKEN_PAGE_SIZE_40  int = 40
	TOKEN_PAGE_SIZE_60  int = 60
	TOKEN_PAGE_SIZE_100 int = 100
)

type STTokenListData struct {
	Address        string  `json:"address"`
	Decimals       int     `json:"decimals"`
	Name           string  `json:"name"`
	Symbol         string  `json:"symbol"`
	MarketCap      float64 `json:"market_cap"`
	Price          float64 `json:"price"`
	Price24HChange float64 `json:"price_24h_change"`
	Holder         int     `json:"holder"`
	CreateTime     int     `json:"created_time"`
}

type STTokenListResp struct {
	Success  bool               `json:"success"`
	DataList []*STTokenListData `json:"data"`
	Errors   basedef.STErrors   `json:"errors"`
}

type TokenListTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenListTool(solanaInfo *basedef.STSolanaDefine) *TokenListTool {
	return &TokenListTool{
		solanaInfo: solanaInfo,
	}
}

func (tlt *TokenListTool) GetTokenList(sortBy string, sortOrder string, page int, pageSize int) (*STTokenListResp, error) {
	newUrl := basedef.G_TOKEN_LIST_URL
	count := 0
	if len(sortBy) > 0 {
		newUrl += "sort_by=" + sortBy
		count++
	}

	if len(sortOrder) > 0 {
		if count > 0 {
			newUrl += "&"
		}
		newUrl += "sort_order=" + sortOrder
		count++
	}

	if page >= 0 {
		if count > 0 {
			newUrl += "&"
		}
		newUrl += "page=" + strconv.Itoa(page)
		count++
	}

	if count > 0 {
		newUrl += "&"
	}
	newUrl += "page_size=" + strconv.Itoa(pageSize)

	data, err := httptool.HttpGet(newUrl, tlt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STTokenListResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
