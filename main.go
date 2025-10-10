package main

import (
	"fmt"
	"math/big"

	"github.com/gagliardetto/solana-go"

	"github.com/alax-mx/geckosdk/gmgn/gmgn_trade"
)

func main() {
	TestEvmTradeTool()
}

func TestEvmTradeTool() {
	config := &gmgn_trade.STEvmConfig{
		ChainType: gmgn_trade.CHAIN_TYPE_ETH,
		RpcURL:    "https://mainnet.infura.io/v3/your-project-id",
		OinchKey:  "1inch-key",
		PriKey:    "wallet-pri-key",
	}
	evmTradeTool := gmgn_trade.NewEvmTradeTool(config)
	if evmTradeTool == nil {
		return
	}

	tokenOut := "0xcf91b70017eabde82c9671e30e5502d312ea6eb2"
	ethDecimals := solana.DecimalsInBigInt(uint32(gmgn_trade.ETH_DECIMALS))
	amount := big.NewInt(int64(0.01 * float64(ethDecimals.Int64()))) // 0.01个ETH
	slippage := float32(1)                                           // 滑点1%
	hash, err := evmTradeTool.Swap(gmgn_trade.ETH_ADDRESS, tokenOut, amount, slippage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hash = ", hash)
}
