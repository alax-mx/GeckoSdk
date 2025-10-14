package gmgn_trade

const (
	GAS_PRICE_LEGACY_LOW     string = "low"
	GAS_PRICE_LEGACY_MEDIUM  string = "medium"
	GAS_PRICE_LEGACY_HIGH    string = "high"
	GAS_PRICE_LEGACY_INSTANT string = "instant"
)

type STEvmConfig struct {
	ChainType string  `json:"chain_type"`
	RpcURL    string  `json:"rpc_url"`
	PriKey    string  `json:"pri_key"`
	OinchKey  string  `json:"oinch_key"`
	GasLegacy string  `json:"gas_legacy"`
	BuyNum    float64 `json:"buy_num"`
	Slippage  float64 `json:"slippage"`
}
