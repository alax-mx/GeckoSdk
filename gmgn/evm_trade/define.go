package evm_trade

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	ETH_RCP_URL        string = "https://eth.llamarpc.com"
	ETH_ADDRESS        string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	ETH_DECIMALS       int    = 18
	USDC_ETH_ADDRESS   string = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	ETH_ROUTER_ADDRESS string = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
)

type EvmConfig struct {
	// 钱包配置
	PrivateKey    string
	WalletAddress common.Address

	// 网络配置
	RPCURL  string
	ChainID int64

	// 交易对配置
	TokenIn  common.Address // ETH地址用 zero address
	TokenOut common.Address // 目标代币地址

	// Uniswap配置
	UniswapRouter common.Address
	Slippage      float64

	// 交易策略
	BuyThreshold  float64 // 买入价格阈值
	SellThreshold float64 // 卖出价格阈值
	TradeAmount   float64 // 每次交易金额(ETH)

	// 监控配置
	CheckInterval int // 检查间隔(秒)
}
