package gmgn_mobi

import (
	"encoding/json"
	"fmt"
)

type STStatusNow struct {
	Hold                 int     `json:"hold"`
	BoughtMore           int     `json:"bought_more"`
	SoldPart             int     `json:"sold_part"`
	Sold                 int     `json:"sold"`
	Transfered           int     `json:"transfered"`
	BoughtRate           string  `json:"bought_rate"`
	HoldingRate          string  `json:"holding_rate"`
	SmartPos             []any   `json:"smart_pos"`
	SmartCountHold       any     `json:"smart_count_hold"`
	SmartCountBoughtMore any     `json:"smart_count_bought_more"`
	SmartCountSoldPart   any     `json:"smart_count_sold_part"`
	SmartCountSold       any     `json:"smart_count_sold"`
	SmartCountTransfered any     `json:"smart_count_transfered"`
	Top10HolderRate      float64 `json:"top_10_holder_rate"`
}
type STHolderInfo struct {
	Status         string   `json:"status"`
	WalletAddress  string   `json:"wallet_address"`
	Tags           []string `json:"tags"`
	MakerTokenTags []string `json:"maker_token_tags"`
}
type STHolderData struct {
	Chain        string         `json:"chain"`
	HolderCount  int            `json:"holder_count"`
	StatusNow    STStatusNow    `json:"statusNow"`
	SoldDiff     int            `json:"sold_diff"`
	SoldPartDiff int            `json:"sold_part_diff"`
	HoldDiff     int            `json:"hold_diff"`
	BoughtMore   int            `json:"bought_more"`
	HolderInfo   []STHolderInfo `json:"holderInfo"`
}

type STHoldersData struct {
	Holders STHolderData `json:"holders"`
}

type GetTokenTopBuyersResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data STHoldersData `json:"data"`
}

type TokenTopBuyersTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenTopBuyersTool(baseUrl string, baseParam string) *TokenTopBuyersTool {
	return &TokenTopBuyersTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (ttbt *TokenTopBuyersTool) Get(tokenAddress string) (*GetTokenTopBuyersResp, error) {
	url := "defi/quotation/v1/tokens/top_buyers/sol/" + tokenAddress + "?" + ttbt.baseParam
	data, err := HttpGet(ttbt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	ret := &GetTokenTopBuyersResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
