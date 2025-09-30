package gmgn_mobi

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/proxy"
)

type STSwapsRank struct {
	ID                       int     `json:"id"`
	Chain                    string  `json:"chain"`
	Address                  string  `json:"address"`
	Symbol                   string  `json:"symbol"`
	Logo                     string  `json:"logo"`
	Price                    float64 `json:"price"`
	PriceChangePercent       float64 `json:"price_change_percent"`
	Swaps                    int     `json:"swaps"`
	Volume                   float64 `json:"volume"`
	Liquidity                any     `json:"liquidity"`
	MarketCap                any     `json:"market_cap"`
	HotLevel                 int     `json:"hot_level"`
	PoolCreationTimestamp    int     `json:"pool_creation_timestamp"`
	HolderCount              int     `json:"holder_count"`
	PoolType                 int     `json:"pool_type"`
	PoolTypeStr              string  `json:"pool_type_str"`
	TwitterUsername          any     `json:"twitter_username"`
	Website                  any     `json:"website"`
	Telegram                 any     `json:"telegram"`
	TotalSupply              int64   `json:"total_supply"`
	OpenTimestamp            int     `json:"open_timestamp"`
	PriceChangePercent1M     float64 `json:"price_change_percent1m"`
	PriceChangePercent5M     float64 `json:"price_change_percent5m"`
	PriceChangePercent1H     float64 `json:"price_change_percent1h"`
	Buys                     int     `json:"buys"`
	Sells                    int     `json:"sells"`
	InitialLiquidity         float64 `json:"initial_liquidity"`
	IsShowAlert              bool    `json:"is_show_alert"`
	Top10HolderRate          float64 `json:"top_10_holder_rate"`
	RenouncedMint            int     `json:"renounced_mint"`
	RenouncedFreezeAccount   int     `json:"renounced_freeze_account"`
	BurnRatio                any     `json:"burn_ratio"`
	BurnStatus               any     `json:"burn_status"`
	Launchpad                any     `json:"launchpad"`
	DevTokenBurnAmount       any     `json:"dev_token_burn_amount"`
	DevTokenBurnRatio        any     `json:"dev_token_burn_ratio"`
	DexscrAd                 int     `json:"dexscr_ad"`
	DexscrUpdateLink         int     `json:"dexscr_update_link"`
	CtoFlag                  int     `json:"cto_flag"`
	TwitterChangeFlag        int     `json:"twitter_change_flag"`
	TwitterRenameCount       int     `json:"twitter_rename_count"`
	CreatorTokenStatus       string  `json:"creator_token_status"`
	CreatorClose             bool    `json:"creator_close,omitempty"`
	Creator                  string  `json:"creator,omitempty"`
	LaunchpadStatus          int     `json:"launchpad_status"`
	RatTraderAmountRate      float64 `json:"rat_trader_amount_rate"`
	CreatorCreatedInnerCount int     `json:"creator_created_inner_count"`
	CreatorCreatedOpenCount  int     `json:"creator_created_open_count"`
	CreatorCreatedOpenRatio  any     `json:"creator_created_open_ratio"`
	BluechipOwnerPercentage  float64 `json:"bluechip_owner_percentage"`
	RugRatio                 any     `json:"rug_ratio"`
	SniperCount              int     `json:"sniper_count"`
	SmartDegenCount          int     `json:"smart_degen_count"`
	RenownedCount            int     `json:"renowned_count"`
	IsWashTrading            bool    `json:"is_wash_trading"`
}

type STTokenSwaps struct {
	Rank []STSwapsRank `json:"rank"`
}

type GetTokenSwapsResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data STTokenSwaps `json:"data"`
}

type TokenSwapsTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenSwapsTool(baseUrl string, baseParam string) *TokenSwapsTool {
	return &TokenSwapsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (tst *TokenSwapsTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tst.proxyInfo = proxyInfo
}

func (tst *TokenSwapsTool) Get(chainType string, interval string, limit int) (*GetTokenSwapsResp, error) {
	url := "defi/quotation/v1/rank/" + chainType + "/swaps/" + interval + "?" + tst.baseParam
	url += "&limit=" + strconv.Itoa(limit)
	url += "&orderby=marketcap"
	url += "&direction=desc"
	url += "&filters=not_wash_trading"
	url += "&filters=renounced"
	url += "&filters=frozen"

	data, err := HttpGet(tst.baseUrl+url, tst.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenSwapsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
