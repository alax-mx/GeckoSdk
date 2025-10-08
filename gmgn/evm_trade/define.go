package evm_trade

const (
	ETH_ADDRESS        string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	ETH_DECIMALS       int    = 18
	USDC_ETH_ADDRESS   string = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	ETH_ROUTER_ADDRESS string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
)

type EvmConfig struct {
	// 钱包配置
	PrivateKey    string
	WalletAddress string

	// 网络配置
	RPCURL  string
	ChainID int64

	// Uniswap配置
	Slippage      float64
	UniswapRouter string
}
