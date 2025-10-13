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
		ChainType: gmgn_trade.CHAIN_TYPE_BSC,
		RpcURL:    "",
		OinchKey:  "",
		PriKey:    "",
		GasLegacy: gmgn_trade.GAS_PRICE_LEGACY_INSTANT,
	}
	evmTradeTool := gmgn_trade.NewEvmTradeTool(config)
	if evmTradeTool == nil {
		return
	}

	// balanceStr, err := evmTradeTool.GetTokenBalance(gmgn_trade.MAIN_ETH20_ADDRESS)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(balanceStr)
	tokenOut := "0x23D3F4EaaA515403C6765bb623F287a8Cca28F2b"
	ethDecimals := solana.DecimalsInBigInt(uint32(gmgn_trade.ETH_DECIMALS))
	amount := big.NewInt(int64(0.01 * float64(ethDecimals.Int64()))) // 0.01个BNB
	slippage := float32(1)                                           // 滑点1%
	hash, err := evmTradeTool.Swap(gmgn_trade.MAIN_ETH20_ADDRESS, tokenOut, amount, slippage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hash = ", hash)
	SellToken(evmTradeTool, tokenOut)
}

func SellToken(evmTradeTool *gmgn_trade.EvmTradeTool, tokenAddress string) {
	balanceStr, err := evmTradeTool.GetTokenBalance(tokenAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balanceStr)

	amount := big.NewInt(0)
	amount.SetString(balanceStr, 10)
	slippage := float32(1)
	hash, err := evmTradeTool.Swap(tokenAddress, gmgn_trade.MAIN_ETH20_ADDRESS, amount, slippage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hash = ", hash)
}
