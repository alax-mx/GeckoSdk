package gmgn_trade

const (
	ETH_ADDRESS           string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	ETH_DECIMALS          int    = 18
	USDC_ETH_ADDRESS      string = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	ETH_ROUTER_V3_ADDRESS string = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
)

const (
	CHAIN_TYPE_ETH     string = "eth"
	CHAIN_TYPE_BSC     string = "bsc"
	CHAIN_TYPE_BASE    string = "base"
	CHAIN_TYPE_POLYGON string = "polygon"
)

type STEvmConfig struct {
	ChainType string
	RpcURL    string
	PriKey    string
	OinchKey  string
}
