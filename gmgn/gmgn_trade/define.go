package gmgn_trade

const (
	MAIN_ETH20_ADDRESS string = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
	ETH_DECIMALS       int    = 18
	USDC_ETH_ADDRESS   string = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	UADC_ETH_DECIMALS  int    = 6
)

const (
	GAS_PRICE_LEGACY_LOW     string = "low"
	GAS_PRICE_LEGACY_MEDIUM  string = "medium"
	GAS_PRICE_LEGACY_HIGH    string = "high"
	GAS_PRICE_LEGACY_INSTANT string = "instant"
)

type STEvmConfig struct {
	ChainType string `json:"chain_type"`
	RpcURL    string `json:"rpc_url"`
	PriKey    string `json:"pri_key"`
	OinchKey  string `json:"oinch_key"`
	GasLegacy string `json:"gas_legacy"`
}
