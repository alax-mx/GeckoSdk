package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STPumps struct {
	Address                  string  `json:"address"`
	Name                     string  `json:"name"`
	Symbol                   string  `json:"symbol"`
	Twitter                  string  `json:"twitter"`
	Website                  any     `json:"website"`
	Telegram                 string  `json:"telegram"`
	CreatedTimestamp         int     `json:"created_timestamp"`
	Complete                 int     `json:"complete"`
	LastTradeTimestamp       int     `json:"last_trade_timestamp"`
	KingOfTheHillTimestamp   int     `json:"king_of_the_hill_timestamp"`
	ReplyCount               any     `json:"reply_count"`
	LastReply                int     `json:"last_reply"`
	UsdMarketCap             string  `json:"usd_market_cap"`
	Price                    any     `json:"price"`
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
	MarketCap5M              string  `json:"market_cap_5m"`
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
	BotDegenCount            string  `json:"bot_degen_count"`
	Top10HolderRate          any     `json:"top_10_holder_rate"`
	RatTraderAmountRate      any     `json:"rat_trader_amount_rate"`
	BundlerTraderAmountRate  any     `json:"bundler_trader_amount_rate"`
	CreatorTokenStatus       string  `json:"creator_token_status"`
	CreatorBalance           string  `json:"creator_balance"`
	CreatorBalanceRate       any     `json:"creator_balance_rate"`
	BluechipOwnerPercentage  float64 `json:"bluechip_owner_percentage"`
	SniperCount              int     `json:"sniper_count"`
	SmartDegenCount          int     `json:"smart_degen_count"`
	RenownedCount            int     `json:"renowned_count"`
	RugRatio                 any     `json:"rug_ratio"`
	IsWashTrading            bool    `json:"is_wash_trading"`
	ImageDup                 string  `json:"image_dup"`
}
type STNewCreations struct {
	Address                  string `json:"address"`
	Name                     string `json:"name"`
	Symbol                   string `json:"symbol"`
	Logo                     any    `json:"logo"`
	LastReply                any    `json:"last_reply"`
	CreatedTimestamp         int    `json:"created_timestamp"`
	LastTradeTimestamp       int    `json:"last_trade_timestamp"`
	KingOfTheHillTimestamp   int    `json:"king_of_the_hill_timestamp"`
	Complete                 int    `json:"complete"`
	TotalSupply              int    `json:"total_supply"`
	TimeSinceKoth            int    `json:"time_since_koth"`
	KothDuration             int    `json:"koth_duration"`
	OpenTimestamp            any    `json:"open_timestamp"`
	HolderCount              int    `json:"holder_count"`
	RugRatio                 any    `json:"rug_ratio"`
	DevTokenBurnAmount       any    `json:"dev_token_burn_amount"`
	DevTokenBurnRatio        any    `json:"dev_token_burn_ratio"`
	DexscrAd                 int    `json:"dexscr_ad"`
	DexscrUpdateLink         int    `json:"dexscr_update_link"`
	CtoFlag                  int    `json:"cto_flag"`
	TwitterChangeFlag        int    `json:"twitter_change_flag"`
	TwitterRenameCount       int    `json:"twitter_rename_count"`
	CreatorCreatedInnerCount int    `json:"creator_created_inner_count"`
	CreatorCreatedOpenCount  int    `json:"creator_created_open_count"`
	CreatorCreatedOpenRatio  any    `json:"creator_created_open_ratio"`
	Status                   int    `json:"status"`
	CreatorBalanceRate       any    `json:"creator_balance_rate"`
	CreatorTokenStatus       string `json:"creator_token_status"`
	Creator                  string `json:"creator"`
	SniperCount              int    `json:"sniper_count"`
	SmartDegenCount          int    `json:"smart_degen_count"`
	RenownedCount            int    `json:"renowned_count"`
	RatTraderAmountRate      any    `json:"rat_trader_amount_rate"`
	BluechipOwnerPercentage  any    `json:"bluechip_owner_percentage"`
	UsdMarketCap             any    `json:"usd_market_cap"`
	IsWashTrading            bool   `json:"is_wash_trading"`
	Top10HolderRate          any    `json:"top_10_holder_rate"`
	Progress                 any    `json:"progress"`
	Price                    any    `json:"price"`
	Volume1M                 any    `json:"volume_1m"`
	Volume5M                 any    `json:"volume_5m"`
	Volume1H                 any    `json:"volume_1h"`
	Volume6H                 any    `json:"volume_6h"`
	Volume24H                any    `json:"volume_24h"`
	Swaps1M                  int    `json:"swaps_1m"`
	Swaps5M                  int    `json:"swaps_5m"`
	Swaps1H                  int    `json:"swaps_1h"`
	Swaps6H                  int    `json:"swaps_6h"`
	Swaps24H                 int    `json:"swaps_24h"`
	PriceChangePercent1M     any    `json:"price_change_percent1m"`
	PriceChangePercent5M     any    `json:"price_change_percent5m"`
	MarketCap1M              any    `json:"market_cap_1m"`
	MarketCap5M              any    `json:"market_cap_5m"`
	ZsetTs                   int64  `json:"zset_ts"`
	ZsetAction               string `json:"zset_action"`
	ZsetSource               string `json:"zset_source"`
	Twitter                  string `json:"twitter,omitempty"`
	Website                  string `json:"website,omitempty"`
	Telegram                 string `json:"telegram,omitempty"`
	ReplyCount               any    `json:"reply_count,omitempty"`
	UpdatedAt                any    `json:"updated_at,omitempty"`
	CreatorClose             any    `json:"creator_close,omitempty"`
	CreatorTokenBalance      string `json:"creator_token_balance,omitempty"`
	Buys1M                   int    `json:"buys_1m,omitempty"`
	Buys5M                   int    `json:"buys_5m,omitempty"`
	Buys1H                   int    `json:"buys_1h,omitempty"`
	Buys6H                   int    `json:"buys_6h,omitempty"`
	Buys24H                  int    `json:"buys_24h,omitempty"`
	Sells1M                  int    `json:"sells_1m,omitempty"`
	Sells5M                  int    `json:"sells_5m,omitempty"`
	Sells1H                  int    `json:"sells_1h,omitempty"`
	Sells6H                  int    `json:"sells_6h,omitempty"`
	Sells24H                 int    `json:"sells_24h,omitempty"`
	TwitterDelPostTokenCount int    `json:"twitter_del_post_token_count,omitempty"`
	EntrapmentRatio          any    `json:"entrapment_ratio,omitempty"`
	BotDegenCount            any    `json:"bot_degen_count,omitempty"`
	BundlerTraderAmountRate  any    `json:"bundler_trader_amount_rate,omitempty"`
	CreatorBalance           string `json:"creator_balance,omitempty"`
	ImageDup                 string `json:"image_dup,omitempty"`
}
type STCompleteds struct {
	Symbol                   string `json:"symbol"`
	Name                     string `json:"name"`
	Logo                     string `json:"logo"`
	TotalSupply              int    `json:"total_supply"`
	Price                    any    `json:"price"`
	HolderCount              int    `json:"holder_count"`
	PriceChangePercent1M     any    `json:"price_change_percent1m"`
	PriceChangePercent5M     any    `json:"price_change_percent5m"`
	PriceChangePercent1H     any    `json:"price_change_percent1h"`
	BurnRatio                any    `json:"burn_ratio"`
	BurnStatus               string `json:"burn_status"`
	IsShowAlert              bool   `json:"is_show_alert"`
	HotLevel                 int    `json:"hot_level"`
	Liquidity                string `json:"liquidity"`
	Top10HolderRate          any    `json:"top_10_holder_rate"`
	RenouncedMint            int    `json:"renounced_mint"`
	RenouncedFreezeAccount   int    `json:"renounced_freeze_account"`
	TwitterRenameCount       int    `json:"twitter_rename_count"`
	RugRatio                 any    `json:"rug_ratio"`
	SniperCount              int    `json:"sniper_count"`
	SmartDegenCount          int    `json:"smart_degen_count"`
	RenownedCount            int    `json:"renowned_count"`
	MarketCap                string `json:"market_cap"`
	IsWashTrading            bool   `json:"is_wash_trading"`
	Creator                  string `json:"creator"`
	CreatorCreatedInnerCount int    `json:"creator_created_inner_count"`
	CreatorCreatedOpenCount  int    `json:"creator_created_open_count"`
	CreatorCreatedOpenRatio  any    `json:"creator_created_open_ratio"`
	CreatorBalanceRate       any    `json:"creator_balance_rate"`
	CreatorTokenStatus       string `json:"creator_token_status"`
	RatTraderAmountRate      any    `json:"rat_trader_amount_rate"`
	BundlerTraderAmountRate  any    `json:"bundler_trader_amount_rate"`
	BluechipOwnerPercentage  any    `json:"bluechip_owner_percentage"`
	Volume                   any    `json:"volume"`
	Swaps                    int    `json:"swaps"`
	Buys                     int    `json:"buys"`
	Sells                    int    `json:"sells"`
	BuyTax                   any    `json:"buy_tax"`
	SellTax                  any    `json:"sell_tax"`
	IsHoneypot               any    `json:"is_honeypot"`
	Renounced                any    `json:"renounced"`
	DevTokenBurnAmount       any    `json:"dev_token_burn_amount"`
	DevTokenBurnRatio        any    `json:"dev_token_burn_ratio"`
	DexscrAd                 int    `json:"dexscr_ad"`
	DexscrUpdateLink         int    `json:"dexscr_update_link"`
	CtoFlag                  int    `json:"cto_flag"`
	TwitterChangeFlag        int    `json:"twitter_change_flag"`
	Address                  string `json:"address"`
	Twitter                  string `json:"twitter"`
	Website                  string `json:"website"`
	Telegram                 string `json:"telegram"`
	OpenTimestamp            int    `json:"open_timestamp"`
	CreatedTimestamp         int    `json:"created_timestamp"`
	UsdMarketCap             string `json:"usd_market_cap"`
	Swaps1H                  int    `json:"swaps_1h"`
	Volume1H                 string `json:"volume_1h"`
	Buys1H                   int    `json:"buys_1h"`
	Sells1H                  int    `json:"sells_1h"`
	BotDegenCount            string `json:"bot_degen_count"`
	QuoteAddress             string `json:"quote_address"`
}

type STTokenPumpRankData struct {
	Pumps        []STPumps        `json:"pumps"`
	NewCreations []STNewCreations `json:"new_creations"`
	Completeds   []STCompleteds   `json:"completeds"`
}

type GetTokenPumpRankResp struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data STTokenPumpRankData `json:"data"`
}

type TokenPumpRankTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenPumpRankTool(baseUrl string, baseParam string) *TokenPumpRankTool {
	return &TokenPumpRankTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tpt *TokenPumpRankTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tpt.proxyInfo = proxyInfo
}

func (tpt *TokenPumpRankTool) Get(chainType string, interval string, limit int) (*GetTokenPumpRankResp, error) {
	url := "defi/quotation/v1/rank/" + chainType + "/pump_ranks/" + interval + "?" + tpt.baseParam
	tmpParam := "{\"filters\":[\"not_wash_trading\"],\"limit\":" + strconv.Itoa(limit) + "}"
	url += "&new_creation=" + tmpParam
	url += "&completed=" + tmpParam
	url += "&pump=" + tmpParam
	data, err := HttpGet(tpt.baseUrl+url, tpt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenPumpRankResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
