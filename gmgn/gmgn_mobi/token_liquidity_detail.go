package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STPoolDetails struct {
	PairAddress          string `json:"pair_address"`
	ContractAddress      string `json:"contract_address"`
	Liquidity            string `json:"liquidity"`
	BaseToken            string `json:"base_token"`
	QuoteToken           string `json:"quote_token"`
	BaseAmount           string `json:"base_amount"`
	QuoteAmount          string `json:"quote_amount"`
	Maker                string `json:"maker"`
	MakerName            string `json:"maker_name"`
	MakerTwitterName     string `json:"maker_twitter_name"`
	MakerTwitterUsername string `json:"maker_twitter_username"`
	MakerAvatar          string `json:"maker_avatar"`
	MakerEns             string `json:"maker_ens"`
	MakerTags            any    `json:"maker_tags"`
	MakerTokenTags       any    `json:"maker_token_tags"`
	CreatorHash          string `json:"creator_hash"`
	QuoteSymbol          string `json:"quote_symbol"`
	TotalTrade           int    `json:"total_trade"`
	Balance              string `json:"balance"`
	HistoryBoughtAmount  string `json:"history_bought_amount"`
	HistorySoldIncome    string `json:"history_sold_income"`
	HistorySoldAmount    string `json:"history_sold_amount"`
	RealizedProfit       string `json:"realized_profit"`
	UnrealizedProfit     string `json:"unrealized_profit"`
	CreatedAt            int    `json:"created_at"`
}

type STTokenLiquidityDetailData struct {
	PoolDetails []STPoolDetails `json:"pool_details"`
}

type GetTokenLiquidityDetailResp struct {
	Code    int                        `json:"code"`
	Reason  string                     `json:"reason"`
	Message string                     `json:"message"`
	Data    STTokenLiquidityDetailData `json:"data"`
}

type TokenLiquidityDetailTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenLiquidityDetailTool(baseUrl string, baseParam string, authStr string) *TokenLiquidityDetailTool {
	return &TokenLiquidityDetailTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tpt *TokenLiquidityDetailTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tdt *TokenLiquidityDetailTool) SetAuthString(authStr string) {
	tdt.authStr = authStr
}

func (tpt *TokenLiquidityDetailTool) Get(chainType string, tokenAddress string) (*GetTokenLiquidityDetailResp, error) {
	url := "vas/api/v1/token_liquidity_detail/" + chainType + "/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.authStr, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLiquidityDetailResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
