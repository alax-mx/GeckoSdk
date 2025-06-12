package gmgn_mobi

const (
	CANDLES_5M  string = "5m"
	CANDLES_15M string = "15m"
	CANDLES_30M string = "30m"
	CANDLES_1H  string = "1h"
	CANDLES_4H  string = "4h"
	CANDLES_12H string = "12h"
	CANDLES_1D  string = "1D"
)

const (
	NEW_PAIR_PERIOD_1M string = "1m"
	NEW_PAIR_PERIOD_5M string = "5m"
	NEW_PAIR_PERIOD_1H string = "1h"
)

const (
	PUMP_RANK_PERIOD_1M  string = "1m"
	PUMP_RANK_PERIOD_5M  string = "5m"
	PUMP_RANK_PERIOD_1H  string = "1h"
	PUMP_RANK_PERIOD_6H  string = "6h"
	PUMP_RANK_PERIOD_24H string = "24h"
)

const (
	WALLET_ORDERBY_PNL_7D      string = "pnl_7d"
	WALLET_ORDERBY_WINRATE_7D  string = "winrate_7d"
	WALLET_ORDERBY_TXS         string = "txs"
	WALLET_ORDERBY_SOL_BALANCE string = "sol_balance"
	WALLET_ORDERBY_LAST_ACTIVE string = "last_active"

	WALLET_TAG_SMART_DEGEN string = "smart_degen"
	WALLET_TAG_PUMP_SMART  string = "pump_smart"
	WALLET_TAG_RENOWNED    string = "renowned"
	WALLET_TAG_FRESH       string = "fresh_wallet"
)

const (
	NEW_PAIR_DIRECTION_DESC string = "desc"

	NEW_PAIR_ORDER_BY_OPEN_TIMESTAMP   string = "open_timestamp"
	NEW_PAIR_ORDER_BY_CREATE_TIMESTAMP string = "creation_timestamp"
)

const (
	WALLET_TOKEN_ORDER_BY_LAST_ACTIVE_TIMESTAMP string = "last_active_timestamp"
	WALLET_TOKEN_DIRECTION_DESC                 string = "desc"
)

type DeviceInfo struct {
	DeviceID string `json:"device_id"`
	ClientID string `json:"client_id"`
	FromApp  string `json:"from_app"`
	AppVer   string `json:"app_ver"`
	Pkg      string `json:"pkg"`
	AppLang  string `json:"app_lang"`
	SysLang  string `json:"sys_lang"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Os       string `json:"os"`
	OsAPI    string `json:"os_api"`
	TzName   string `json:"tz_name"`
	TzOffset string `json:"tz_offset"`
	Gpv      string `json:"gpv"`
}
