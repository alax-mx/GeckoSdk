package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STHistory struct {
	Maker                string   `json:"maker"`
	BaseAmount           string   `json:"base_amount"`
	QuoteAmount          string   `json:"quote_amount"`
	AmountUsd            string   `json:"amount_usd"`
	Timestamp            int      `json:"timestamp"`
	Event                string   `json:"event"`
	TxHash               string   `json:"tx_hash"`
	PriceUsd             string   `json:"price_usd"`
	ID                   string   `json:"id"`
	IsOpenOrClose        int      `json:"is_open_or_close"`
	TokenAddress         string   `json:"token_address"`
	MakerTags            []any    `json:"maker_tags"`
	MakerTokenTags       []string `json:"maker_token_tags"`
	QuoteAddress         string   `json:"quote_address"`
	QuoteSymbol          string   `json:"quote_symbol"`
	TotalTrade           int      `json:"total_trade"`
	Balance              string   `json:"balance"`
	HistoryBoughtAmount  string   `json:"history_bought_amount"`
	HistorySoldIncome    string   `json:"history_sold_income"`
	HistorySoldAmount    string   `json:"history_sold_amount"`
	RealizedProfit       string   `json:"realized_profit"`
	UnrealizedProfit     string   `json:"unrealized_profit"`
	MakerName            string   `json:"maker_name"`
	MakerTwitterUsername string   `json:"maker_twitter_username"`
	MakerTwitterName     string   `json:"maker_twitter_name"`
	MakerAvatar          string   `json:"maker_avatar"`
	MakerEns             string   `json:"maker_ens"`
}

type STTokenTradesData struct {
	History []STHistory `json:"history"`
	Next    string      `json:"next"`
}

type GetTokenTradesResp struct {
	Code    int               `json:"code"`
	Reason  string            `json:"reason"`
	Message string            `json:"message"`
	Data    STTokenTradesData `json:"data"`
}

type TokenTradesTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenTradesTool(baseUrl string, baseParam string) *TokenTradesTool {
	return &TokenTradesTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tst *TokenTradesTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tst.proxyInfo = proxyInfo
}

func (tst *TokenTradesTool) Get(tokenAddress string, limit int) (*GetTokenTradesResp, error) {
	url := "vas/api/v1/token_trades/sol/" + tokenAddress + "?" + tst.baseParam
	url += "&event=buy"
	url += "&event=sell"
	url += "&limit=" + strconv.Itoa(limit)
	url += "&dev=dev"
	data, err := HttpGet(tst.baseUrl+url, tst.proxyInfo)
	if err != nil {
		return nil, err
	}
	ret := &GetTokenTradesResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
