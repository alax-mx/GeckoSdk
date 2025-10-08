package gmgn_trade

import (
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
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
