package gmgn_define

const (
	ORDER_BY_OPEN_TIME   = "open_timestamp"
	ORDER_BY_CREATE_TIME = "creation_timestamp"

	DIRECTION_DESC = "desc"
)

const (
	G_BASE_GMGN_WEB_DEFI_URL string = "https://gmgn.ai/defi/quotation/"
	G_BASE_GMGN_TRADE_URL    string = "https://gmgn.ai/defi/router/v1/sol/tx"
	G_BASE_GMGN_MOBI_URL     string = "https://gmgn.mobi/"
)

type STTokenInfo struct {
	Name               string        `json:"name"`
	Symbol             string        `json:"symbol"`
	Address            string        `json:"address"`
	Decimals           int           `json:"decimals"`
	HolderCount        int           `json:"holder_count"`
	Price              float64       `json:"price"`
	Price1M            float64       `json:"price_1m"`
	Price5M            float64       `json:"price_5m"`
	Price1H            float64       `json:"price_1h"`
	Price6H            float64       `json:"price_6h"`
	Price24H           float64       `json:"price_24h"`
	Swaps5M            int           `json:"swaps_5m"`
	Swaps1H            int           `json:"swaps_1h"`
	Swaps6H            int           `json:"swaps_6h"`
	Swaps24H           int           `json:"swaps_24h"`
	Liquidity          float64       `json:"liquidity"`
	MaxSupply          int           `json:"max_supply"`
	TotalSupply        int           `json:"total_supply"`
	BiggestPoolAddress string        `json:"biggest_pool_address"`
	Chain              string        `json:"chain"`
	CreationTimestamp  int           `json:"creation_timestamp"`
	OpenTimestamp      int           `json:"open_timestamp"`
	CirculatingSupply  int           `json:"circulating_supply"`
	HighPrice          float64       `json:"high_price"`
	HighPriceTimestamp int           `json:"high_price_timestamp"`
	LowPrice           float64       `json:"low_price"`
	LowPriceTimestamp  int           `json:"low_price_timestamp"`
	Buys1M             int           `json:"buys_1m"`
	Sells1M            int           `json:"sells_1m"`
	Swaps1M            int           `json:"swaps_1m"`
	Volume1M           float64       `json:"volume_1m"`
	BuyVolume1M        float64       `json:"buy_volume_1m"`
	SellVolume1M       float64       `json:"sell_volume_1m"`
	NetInVolume1M      float64       `json:"net_in_volume_1m"`
	Buys5M             int           `json:"buys_5m"`
	Sells5M            int           `json:"sells_5m"`
	Volume5M           float64       `json:"volume_5m"`
	BuyVolume5M        float64       `json:"buy_volume_5m"`
	SellVolume5M       float64       `json:"sell_volume_5m"`
	NetInVolume5M      float64       `json:"net_in_volume_5m"`
	Buys1H             int           `json:"buys_1h"`
	Sells1H            int           `json:"sells_1h"`
	Volume1H           float64       `json:"volume_1h"`
	BuyVolume1H        float64       `json:"buy_volume_1h"`
	SellVolume1H       float64       `json:"sell_volume_1h"`
	NetInVolume1H      float64       `json:"net_in_volume_1h"`
	Buys6H             int           `json:"buys_6h"`
	Sells6H            int           `json:"sells_6h"`
	Volume6H           float64       `json:"volume_6h"`
	BuyVolume6H        float64       `json:"buy_volume_6h"`
	SellVolume6H       float64       `json:"sell_volume_6h"`
	NetInVolume6H      float64       `json:"net_in_volume_6h"`
	Buys24H            int           `json:"buys_24h"`
	Sells24H           int           `json:"sells_24h"`
	Volume24H          float64       `json:"volume_24h"`
	BuyVolume24H       float64       `json:"buy_volume_24h"`
	SellVolume24H      float64       `json:"sell_volume_24h"`
	NetInVolume24H     float64       `json:"net_in_volume_24h"`
	Fdv                float64       `json:"fdv"`
	MarketCap          float64       `json:"market_cap"`
	Link               STLink        `json:"link"`
	SocialLinks        STSocialLinks `json:"social_links"`
}

type STLink struct {
	Geckoterminal string `json:"geckoterminal"`
	Gmgn          string `json:"gmgn"`
}

type STSocialLinks struct {
	ID              int    `json:"id"`
	Chain           string `json:"chain"`
	Address         string `json:"address"`
	TwitterUsername string `json:"twitter_username"`
	WebSite         string `json:"website"`
	Telegram        string `json:"telegram"`
	Bitbucket       string `json:"bitbucket"`
	Discord         string `json:"discord"`
	Description     string `json:"description"`
	FaceBook        string `json:"facebook"`
	Github          string `json:"github"`
	Instagram       string `json:"instagram"`
	LinkedIN        string `json:"linkedin"`
	Medium          string `json:"medium"`
	Reddit          string `json:"reddit"`
	Tiktok          string `json:"tiktok"`
	Youtube         string `json:"youtube"`
	UpdatedAt       int    `json:"updated_at"`
	VerifyStatus    int    `json:"verify_status"`
}

type STTokenPairs struct {
	ID                string          `json:"verify_status"`
	Address           string          `json:"address"`
	BaseAddress       string          `json:"base_address"`
	QuoteAddress      string          `json:"quote_address"`
	CreationTimestamp int             `json:"creation_timestamp"`
	PoolType          int             `json:"pool_type"`
	QuoteSymbol       string          `json:"quote_symbol"`
	BurnRatio         string          `json:"burn_ratio"`
	BurnStatus        string          `json:"burn_status"`
	OpenTimestamp     int             `json:"open_timestamp"`
	Launchpad         string          `json:"launchpad"`
	QuoteReserveUsd   string          `json:"quote_reserve_usd"`
	BaseTokenInfo     STBaseTokenInfo `json:"base_token_info"`
	BotDegenCount     string          `json:"bot_degen_count"`
}

type STBaseTokenInfo struct {
	Name                    string        `json:"name"`
	Address                 string        `json:"address"`
	Symbol                  string        `json:"symbol"`
	HolderCount             int           `json:"holder_count"`
	Logo                    string        `json:"logo"`
	TotalSupply             int           `json:"total_supply"`
	Creator                 string        `json:"creator"`
	CreatorClose            bool          `json:"creator_close"`
	CreatorTokenStatus      string        `json:"creator_token_status"`
	BurnRatio               any           `json:"burn_ratio"`
	BurnStatus              string        `json:"burn_status"`
	PoolID                  string        `json:"pool_id"`
	BiggestPoolAddress      string        `json:"biggest_pool_address"`
	Liquidity               any           `json:"liquidity"`
	LaunchpadStatus         int           `json:"launchpad_status"`
	IsShowAlert             bool          `json:"is_show_alert"`
	HotLevel                int           `json:"hot_level"`
	Buys                    int           `json:"buys"`
	Sells                   int           `json:"sells"`
	Renounced               int           `json:"renounced"`
	RenouncedMint           int           `json:"renounced_mint"`
	RenouncedFreezeAccount  int           `json:"renounced_freeze_account"`
	SocialLinks             STSocialLinks `json:"social_links"`
	SniperCount             int           `json:"sniper_count"`
	DevTokenBurnRatio       any           `json:"dev_token_burn_ratio"`
	DexscrAd                int           `json:"dexscr_ad"`
	DexscrUpdateLink        int           `json:"dexscr_update_link"`
	CtoFlag                 int           `json:"cto_flag"`
	TwitterChangeFlag       int           `json:"twitter_change_flag"`
	TwitterRenameCount      int           `json:"twitter_rename_count"`
	RatTraderAmountRate     any           `json:"rat_trader_amount_rate"`
	Top10HolderRate         any           `json:"top_10_holder_rate"`
	RugRatio                any           `json:"rug_ratio"`
	BluechipOwnerPercentage any           `json:"bluechip_owner_percentage"`
	IsHoneypot              any           `json:"is_honeypot"`
}

type STTokenPairsData struct {
	Pairs []*STTokenPairs `json:"pairs"`
}

type GetTokenPairResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data STTokenPairsData `json:"Data"`
}

type STTokenPrice struct {
	UsdPrice float64 `json:"usd_price"`
}

type GetTokenPriceResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data STTokenPrice `json:"Data"`
}
