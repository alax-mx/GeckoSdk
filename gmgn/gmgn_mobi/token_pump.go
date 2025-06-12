package gmgn_mobi

import (
	"encoding/json"
	"strconv"
)

type STRank struct {
	Address                  string  `json:"address"`
	Name                     string  `json:"name"`
	Symbol                   string  `json:"symbol"`
	Twitter                  any     `json:"twitter"`
	Website                  any     `json:"website"`
	Telegram                 any     `json:"telegram"`
	CreatedTimestamp         int     `json:"created_timestamp"`
	Complete                 int     `json:"complete"`
	LastTradeTimestamp       int     `json:"last_trade_timestamp"`
	KingOfTheHillTimestamp   int     `json:"king_of_the_hill_timestamp"`
	ReplyCount               any     `json:"reply_count"`
	LastReply                int     `json:"last_reply"`
	UsdMarketCap             string  `json:"usd_market_cap"`
	Price                    float64 `json:"price"`
	UpdatedAt                int     `json:"updated_at"`
	Creator                  string  `json:"creator"`
	CreatorClose             any     `json:"creator_close"`
	CreatorTokenBalance      string  `json:"creator_token_balance"`
	Progress                 any     `json:"progress"`
	TotalSupply              int     `json:"total_supply"`
	Status                   int     `json:"status"`
	Logo                     string  `json:"logo"`
	KothDuration             int     `json:"koth_duration"`
	TimeSinceKoth            int     `json:"time_since_koth"`
	OpenTimestamp            any     `json:"open_timestamp"`
	Volume1M                 any     `json:"volume_1m"`
	Volume5M                 any     `json:"volume_5m"`
	Volume1H                 any     `json:"volume_1h"`
	Volume6H                 any     `json:"volume_6h"`
	Volume24H                any     `json:"volume_24h"`
	Swaps1M                  int     `json:"swaps_1m"`
	Swaps5M                  int     `json:"swaps_5m"`
	Swaps1H                  int     `json:"swaps_1h"`
	Swaps6H                  int     `json:"swaps_6h"`
	Swaps24H                 int     `json:"swaps_24h"`
	Buys1M                   int     `json:"buys_1m"`
	Buys5M                   int     `json:"buys_5m"`
	Buys1H                   int     `json:"buys_1h"`
	Buys6H                   int     `json:"buys_6h"`
	Buys24H                  int     `json:"buys_24h"`
	Sells1M                  int     `json:"sells_1m"`
	Sells5M                  int     `json:"sells_5m"`
	Sells1H                  int     `json:"sells_1h"`
	Sells6H                  int     `json:"sells_6h"`
	Sells24H                 int     `json:"sells_24h"`
	PriceChangePercent1M     any     `json:"price_change_percent1m"`
	PriceChangePercent5M     any     `json:"price_change_percent5m"`
	MarketCap1M              any     `json:"market_cap_1m"`
	MarketCap5M              any     `json:"market_cap_5m"`
	HolderCount              int     `json:"holder_count"`
	DevTokenBurnAmount       any     `json:"dev_token_burn_amount"`
	DevTokenBurnRatio        any     `json:"dev_token_burn_ratio"`
	DexscrAd                 int     `json:"dexscr_ad"`
	DexscrUpdateLink         int     `json:"dexscr_update_link"`
	CtoFlag                  int     `json:"cto_flag"`
	TwitterChangeFlag        int     `json:"twitter_change_flag"`
	TwitterRenameCount       int     `json:"twitter_rename_count"`
	TwitterDelPostTokenCount int     `json:"twitter_del_post_token_count"`
	CreatorCreatedInnerCount int     `json:"creator_created_inner_count"`
	CreatorCreatedOpenCount  int     `json:"creator_created_open_count"`
	CreatorCreatedOpenRatio  any     `json:"creator_created_open_ratio"`
	EntrapmentRatio          any     `json:"entrapment_ratio"`
	BotDegenCount            any     `json:"bot_degen_count"`
	Top10HolderRate          float64 `json:"top_10_holder_rate"`
	RatTraderAmountRate      any     `json:"rat_trader_amount_rate"`
	BundlerTraderAmountRate  any     `json:"bundler_trader_amount_rate"`
	CreatorTokenStatus       any     `json:"creator_token_status"`
	CreatorBalance           any     `json:"creator_balance"`
	CreatorBalanceRate       any     `json:"creator_balance_rate"`
	BluechipOwnerPercentage  any     `json:"bluechip_owner_percentage"`
	SniperCount              int     `json:"sniper_count"`
	SmartDegenCount          int     `json:"smart_degen_count"`
	RenownedCount            int     `json:"renowned_count"`
	RugRatio                 any     `json:"rug_ratio"`
	IsWashTrading            bool    `json:"is_wash_trading"`
	ImageDup                 string  `json:"image_dup"`
}

type STTokenPumpData struct {
	Rank []STRank `json:"rank"`
}

type GetTokenPumpResp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data STTokenPumpData `json:"data"`
}

type TokenPumpTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenPumpTool(baseUrl string, baseParam string) *TokenPumpTool {
	return &TokenPumpTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tpt *TokenPumpTool) Get(interval string, limit int) (*GetTokenPumpResp, error) {
	url := "defi/quotation/v1/rank/sol/pump/" + interval + "?" + tpt.baseParam
	url += "&limit=" + strconv.Itoa(limit)
	url += "&orderby=progress"
	url += "&direction=desc"
	url += "&filters[]=not_wash_trading"
	url += "&soaring=true"

	data, err := HttpGet(tpt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPumpResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
