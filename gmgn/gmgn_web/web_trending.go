package gmgn_web

import (
	"encoding/json"
	"strconv"
)

type STTrendingTokenInfo struct {
	ID                        int     `json:"id"`
	Chain                     string  `json:"chain"`
	Address                   string  `json:"address"`
	Symbol                    string  `json:"symbol"`
	Logo                      string  `json:"logo"`
	Price                     float64 `json:"price"`
	PriceChangePercent        float64 `json:"price_change_percent"`
	Swaps                     int     `json:"swaps"`
	Volume                    float64 `json:"volume"`
	Liquidity                 float64 `json:"liquidity"`
	MarketCap                 float64 `json:"market_cap"`
	HotLevel                  int     `json:"hot_level"`
	PoolCreationTimestamp     int     `json:"pool_creation_timestamp"`
	HolderCount               int     `json:"holder_count"`
	PoolType                  int     `json:"pool_type"`
	PoolTypeStr               string  `json:"pool_type_str"`
	TwitterUsername           string  `json:"twitter_username"`
	Website                   string  `json:"website"`
	Telegram                  string  `json:"telegram"`
	TotalSupply               int     `json:"total_supply"`
	OpenTimestamp             int     `json:"open_timestamp"`
	PriceChangePercent1m      float64 `json:"price_change_percent1m"`
	PriceChangePercent5m      float64 `json:"price_change_percent5m"`
	PriceChangePercent1h      float64 `json:"price_change_percent1h"`
	Buys                      int     `json:"buys"`
	Sells                     int     `json:"sells"`
	InitialLiquidity          float64 `json:"initial_liquidity"`
	IsShowAlert               bool    `json:"is_show_alert"`
	Top10HolderRate           float64 `json:"top_10_holder_rate"`
	RenouncedMint             int     `json:"renounced_mint"`
	RenouncedFreezeAccount    int     `json:"renounced_freeze_account"`
	BurnStatus                string  `json:"burn_status"`
	TwitterChangeFlag         int     `json:"twitter_change_flag"`
	TwitterRenameCount        int     `json:"twitter_rename_count"`
	CreatorTokenStatus        string  `json:"creator_token_status"`
	CreatorClose              bool    `json:"creator_close"`
	Creator                   string  `json:"creator"`
	LaunchpadStatus           int     `json:"launchpad_status"`
	RatTraderAmountRate       float64 `json:"rat_trader_amount_rate"`
	CreatorCreatedInner_count int     `json:"creator_created_inner_count"`
	CreatorCreatedOpenCount   int     `json:"creator_created_open_count"`
	BluechipOwnerPercentage   float64 `json:"bluechip_owner_percentage"`
	SniperCount               int     `json:"sniper_count"`
	SmartDegenCount           int     `json:"smart_degen_count"`
	RenownedCount             int     `json:"renowned_count"`
	IsWashTrading             bool    `json:"is_wash_trading"`
}

type STData struct {
	RankList []*STTrendingTokenInfo `json:"rank"`
}

type GetTrendingTokensResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data STData `json:"data"`
}

type WebTrendingTool struct {
	baseUrl string
}

func NewWebTrendingTool(baseUrl string) *WebTrendingTool {
	return &WebTrendingTool{
		baseUrl: baseUrl,
	}
}

func (gtt *WebTrendingTool) GetTrendingTokens(timeFrame string, direction string, limit int) (*GetTrendingTokensResp, error) {
	urlAddr := "v1/rank/sol/swaps/" + timeFrame + "?"
	urlAddr += "&orderby=swaps&direction=" + direction
	urlAddr += "&limit" + strconv.Itoa(limit)
	data, err := HttpGet(gtt.baseUrl + urlAddr)
	if err != nil {
		return nil, err
	}

	resp := &GetTrendingTokensResp{}
	err2 := json.Unmarshal(data, resp)
	if err2 != nil {
		return nil, err2
	}
	return resp, nil
}
