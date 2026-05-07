package wss_mobi

type STNewPoolInfo struct {
	C  string       `json:"c"`
	Rg string       `json:"rg"`
	P  []STPoolInfo `json:"p"`
}

type STPoolInfo struct {
	ID                  any             `json:"id"`
	Address             string          `json:"a"`
	Exchange            string          `json:"ex"`
	PoolAddress         string          `json:"pa"`
	BaseAddress         string          `json:"ba"`
	QuoteAddress        string          `json:"qa"`
	QuoteReserve        string          `json:"qr"`
	InitialLiquidity    string          `json:"il"`
	InitialQuoteReserve string          `json:"iqr"`
	Liquidity           string          `json:"l"`
	OpenTimestamp       int64           `json:"ot"`
	PoolType            int64           `json:"pt"`
	QuoteSymbol         string          `json:"qs"`
	BaseTokenInfo       STBaseTokenInfo `json:"bti"`
}

type STBaseTokenInfo struct {
	Symbol                  string  `json:"s"`
	Name                    string  `json:"n"`
	Logo                    string  `json:"l"`
	TotalSupply             int64   `json:"ts"`
	Valume24H               float64 `json:"v24h"`
	Sell24H                 int64   `json:"s24h"`
	Price                   float64 `json:"p"`
	HolderCount             int64   `json:"hc"`
	BurnRatio               string  `json:"br"`
	BurnStatus              string  `json:"bs"`
	IsShowAlert             bool    `json:"isa"`
	HotLevel                int64   `json:"hl"`
	Top10HolderRate         int64   `json:"t10hr"`
	RenouncedMint           int64   `json:"rm"`
	RenouncedFreezeAccount  int64   `json:"rfa"`
	MarketCap               float64 `json:"mc"`
	Creator                 string  `json:"ctr"`
	CreatorBalanceRate      int64   `json:"cbr"`
	CreatorTokenStatus      string  `json:"cts"`
	RatTraderAmountRate     int64   `json:"rtar"`
	BluechipOwnerPercentage int64   `json:"bop"`
	SmartDegenCount         int64   `json:"sdc"`
	RenownedCount           int64   `json:"rc"`
	PriceGrowth             float64 `json:"pg"`
	DexscrAd                int64   `json:"dx_ad"`
	DexscrUpdateLink        int64   `json:"dx_ul"`
	CtoFlag                 bool    `json:"cto"`
}
