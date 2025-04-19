package geck_sdk

const (
	INCLUDE_BASE_TOKEN  string = "base_token"
	INCLUDE_QUOTE_TOKEN string = "quote_token"
	INCLUDE_DEX         string = "dex"
)

const (
	SORY_BY_H24_TX_COUNT_DESC   string = "h24_tx_count_desc"
	SORY_BY_H24_VOLUME_USD_DESC string = "h24_volume_usd_desc"
)

type STErrors struct {
	Status string `json:"status"`
	Title  string `json:"title"`
}

type STPriceChangePercentage struct {
	M5  string `json:"m5"`
	M15 string `json:"m15"`
	M30 string `json:"m30"`
	H1  string `json:"h1"`
	H6  string `json:"h6"`
	H24 string `json:"h24"`
}

type STTransactionsData struct {
	Buys    int `json:"buys"`
	Sells   int `json:"sells"`
	Buyers  int `json:"buyers"`
	Sellers int `json:"sellers"`
}

type STTransactions struct {
	M5  STTransactionsData `json:"m5"`
	M15 STTransactionsData `json:"m15"`
	M30 STTransactionsData `json:"m30"`
	H1  STTransactionsData `json:"h1"`
	H6  STTransactionsData `json:"h6"`
	H24 STTransactionsData `json:"h24"`
}

type STVolumeUSD struct {
	M5  string `json:"m5"`
	M15 string `json:"m15"`
	M30 string `json:"m30"`
	H1  string `json:"h1"`
	H6  string `json:"h6"`
	H24 string `json:"h24"`
}

type STRelationShipsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type STRelationShipsItem struct {
	Data STRelationShipsData `json:"data"`
}

type STRelationsShips struct {
	BaseToken  STRelationShipsItem `json:"base_token"`
	QuoteToken STRelationShipsItem `json:"quote_token"`
	Dex        STRelationShipsItem `json:"dex"`
}
