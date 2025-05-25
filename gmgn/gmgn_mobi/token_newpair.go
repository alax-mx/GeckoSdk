package gmgn_mobi

import (
	"encoding/json"
	"strconv"
)

type STSocialLinks struct {
	TwitterUsername string `json:"twitter_username"`
	Website         string `json:"website"`
	Telegram        string `json:"telegram"`
}
type STBaseTokenInfo struct {
	Name                    string        `json:"name"`
	Address                 string        `json:"address"`
	Symbol                  string        `json:"symbol"`
	Logo                    any           `json:"logo"`
	TotalSupply             any           `json:"total_supply"`
	Creator                 string        `json:"creator"`
	CreatorClose            bool          `json:"creator_close"`
	CreatorTokenStatus      string        `json:"creator_token_status"`
	BurnRatio               string        `json:"burn_ratio"`
	BurnStatus              string        `json:"burn_status"`
	PoolID                  string        `json:"pool_id"`
	BiggestPoolAddress      string        `json:"biggest_pool_address"`
	LaunchpadStatus         int           `json:"launchpad_status"`
	IsShowAlert             bool          `json:"is_show_alert"`
	HotLevel                int           `json:"hot_level"`
	BuyTax                  any           `json:"buy_tax"`
	SellTax                 any           `json:"sell_tax"`
	IsHoneypot              any           `json:"is_honeypot"`
	Renounced               any           `json:"renounced"`
	RenouncedMint           int           `json:"renounced_mint"`
	RenouncedFreezeAccount  int           `json:"renounced_freeze_account"`
	SocialLinks             STSocialLinks `json:"social_links"`
	DevTokenBurnAmount      any           `json:"dev_token_burn_amount"`
	DevTokenBurnRatio       any           `json:"dev_token_burn_ratio"`
	DexscrAd                int           `json:"dexscr_ad"`
	DexscrUpdateLink        int           `json:"dexscr_update_link"`
	CtoFlag                 int           `json:"cto_flag"`
	TwitterChangeFlag       int           `json:"twitter_change_flag"`
	TwitterRenameCount      int           `json:"twitter_rename_count"`
	Liquidity               any           `json:"liquidity"`
	RugRatio                any           `json:"rug_ratio"`
	Top10HolderRate         any           `json:"top_10_holder_rate"`
	CreatorBalanceRate      any           `json:"creator_balance_rate"`
	Price                   any           `json:"price"`
	PriceChangePercent1M    any           `json:"price_change_percent1m"`
	PriceChangePercent5M    any           `json:"price_change_percent5m"`
	PriceChangePercent1H    any           `json:"price_change_percent1h"`
	HolderCount             any           `json:"holder_count"`
	SniperCount             any           `json:"sniper_count"`
	SmartDegenCount         any           `json:"smart_degen_count"`
	RenownedCount           any           `json:"renowned_count"`
	RatTraderAmountRate     any           `json:"rat_trader_amount_rate"`
	BluechipOwnerPercentage any           `json:"bluechip_owner_percentage"`
	MarketCap               any           `json:"market_cap"`
	IsWashTrading           bool          `json:"is_wash_trading"`
	Volume                  any           `json:"volume"`
	Swaps                   any           `json:"swaps"`
	Buys                    any           `json:"buys"`
	Sells                   any           `json:"sells"`
}
type STPairs struct {
	ID                  any             `json:"id"`
	Address             string          `json:"address"`
	BaseAddress         string          `json:"base_address"`
	QuoteAddress        string          `json:"quote_address"`
	CreationTimestamp   int             `json:"creation_timestamp"`
	PoolType            int             `json:"pool_type"`
	QuoteSymbol         string          `json:"quote_symbol"`
	BurnRatio           string          `json:"burn_ratio"`
	BurnStatus          string          `json:"burn_status"`
	OpenTimestamp       int             `json:"open_timestamp"`
	Launchpad           string          `json:"launchpad"`
	Liquidity           any             `json:"liquidity"`
	QuoteReserve        any             `json:"quote_reserve"`
	InitialLiquidity    any             `json:"initial_liquidity"`
	InitialQuoteReserve any             `json:"initial_quote_reserve"`
	QuoteReserveUsd     any             `json:"quote_reserve_usd"`
	BaseTokenInfo       STBaseTokenInfo `json:"base_token_info"`
}

type TokenNewPairData struct {
	Pairs []STPairs `json:"pairs"`
}

type GetTokenNewPairResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data TokenNewPairData `json:"data"`
}

type TokenNewPairTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenNewPairTool(baseUrl string, baseParam string) *TokenNewPairTool {
	return &TokenNewPairTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tnpt *TokenNewPairTool) Get(period string, limit int, orderBy string, direction string) (*GetTokenNewPairResp, error) {
	url := "defi/quotation/v1/pairs/sol/new_pairs/" + period + "?" + tnpt.baseParam
	url += "&period=" + period
	url += "&limit=" + strconv.Itoa(limit)
	url += "&orderby=" + orderBy
	url += "&direction=" + direction
	data, err := HttpGet(tnpt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenNewPairResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
