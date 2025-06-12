package gmgn_mobi

import "encoding/json"

type STTokens struct {
	TokenAddress             string `json:"token_address"`
	Symbol                   string `json:"symbol"`
	Chain                    string `json:"chain"`
	CreateTimestamp          int    `json:"create_timestamp"`
	IsOpen                   bool   `json:"is_open"`
	MarcketCap               string `json:"marcket_cap"`
	BiggestPool              string `json:"biggest_pool"`
	PoolLiquidity            string `json:"pool_liquidity"`
	Holders                  int    `json:"holders"`
	Swap1H                   int    `json:"swap_1h"`
	Volume1H                 string `json:"volume_1h"`
	Logo                     string `json:"logo"`
	LiquidityLess4K          bool   `json:"liquidity_less_4k"`
	CtoFlag                  bool   `json:"cto_flag"`
	TwitterNameChangeHistory []any  `json:"twitter_name_change_history"`
	DexscrAd                 bool   `json:"dexscr_ad"`
	DexscrUpdateLink         bool   `json:"dexscr_update_link"`
	IsPump                   bool   `json:"is_pump"`
	LaunchpadPlatform        string `json:"launchpad_platform"`
}

type STTokenDevCreatedData struct {
	LastCreateTimestamp int        `json:"last_create_timestamp"`
	InnerCount          int        `json:"inner_count"`
	OpenCount           int        `json:"open_count"`
	OpenRatio           string     `json:"open_ratio"`
	Tokens              []STTokens `json:"tokens"`
}

type GetTokenDevCreatedResp struct {
	Code    int                   `json:"code"`
	Reason  string                `json:"reason"`
	Message string                `json:"message"`
	Data    STTokenDevCreatedData `json:"data"`
}

type TokenDevCreatedTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenDevCreatedTool(baseUrl string, baseParam string) *TokenDevCreatedTool {
	return &TokenDevCreatedTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tdt *TokenDevCreatedTool) Get(tokenAddress string) (*GetTokenDevCreatedResp, error) {
	url := "api/v1/dev_created_tokens/sol/" + tokenAddress + "?" + tdt.baseParam
	data, err := HttpGet(tdt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenDevCreatedResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
