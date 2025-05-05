package gecko_sdk

const (
	INCLUDE_BASE_TOKEN  string = "base_token"
	INCLUDE_QUOTE_TOKEN string = "quote_token"
	INCLUDE_DEX         string = "dex"
)

const (
	SORY_BY_H24_TX_COUNT_DESC   string = "h24_tx_count_desc"
	SORY_BY_H24_VOLUME_USD_DESC string = "h24_volume_usd_desc"
)

const (
	TRENDING_DURATION_5M  string = "5m"
	TRENDING_DURATION_1H  string = "1h"
	TRENDING_DURATION_6H  string = "6h"
	TRENDING_DURATION_24H string = "24h"
)

const (
	OHLCV_TIME_FRAME_TYPE_MINUTE string = "minute"
	OHLCV_TIME_FRAME_TYPE_HOUR   string = "hour"
	OHLCV_TIME_FRAME_TYPE_DAY    string = "day"
)

const (
	OHLCV_AGREGATE_MINUTE_1  string = "1"
	OHLCV_AGREGATE_MINUTE_5  string = "5"
	OHLCV_AGREGATE_MINUTE_15 string = "15"
	OHLCV_AGREGATE_HOUR_1    string = "1"
	OHLCV_AGREGATE_HOUR_4    string = "4"
	OHLCV_AGREGATE_HOUR_12   string = "12"
	OHLCV_AGREGATE_DAY_1     string = "1"
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

type STIncludedAttributes struct {
	Address         string `json:"address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        int    `json:"decimals"`
	ImageURL        string `json:"image_url"`
	CoingeckoCoinID string `json:"coingecko_coin_id"`
}

type STIncluded struct {
	ID         string               `json:"id"`
	Type       string               `json:"type"`
	Attributes STIncludedAttributes `json:"attributes"`
}

type STGTScoreDetails struct {
	Pool        float64 `json:"pool"`
	Transaction float64 `json:"transaction"`
	Creation    float64 `json:"creation"`
	Info        float64 `json:"info"`
	Holders     float64 `json:"holders"`
}

type STDistributionPercentage struct {
	Top10    string `json:"top_10"`
	Top11_20 string `json:"11_20"`
	Top21_40 string `json:"21_40"`
	Rest     string `json:"rest"`
}

type STHolders struct {
	Count                  int                      `json:"count"`
	DistributionPercentage STDistributionPercentage `json:"distribution_percentage"`
	LastUpdated            string                   `json:"last_updated"`
}
