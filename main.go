package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_trade"
)

func main() {
	// TestEvmTradeTool()
	deviceInfo := loaddevice("device.json")
	if deviceInfo == nil {
		return
	}
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo, nil)
	_, err := gmgnTool.GetMobiTool().GetTokenNewPairTool().Get("bsc", gmgn_mobi.NEW_PAIR_PERIOD_1M, 2, gmgn_mobi.NEW_PAIR_ORDER_BY_CREATE_TIMESTAMP)
	if err != nil {
		fmt.Println("TokenLaunchpadMonitor err:", err)
	}
}

func loaddevice(path string) *gmgn_mobi.DeviceInfo {
	data, err := baseutils.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cfg := &gmgn_mobi.DeviceInfo{}
	err2 := json.Unmarshal(data, cfg)
	if err2 != nil {
		fmt.Println(err)
		return nil
	}
	return cfg
}

func TestEvmTradeTool() {
	config := &gmgn_trade.STEvmConfig{
		ChainType: gmgn_define.CHAIN_TYPE_BSC,
		RpcURL:    "",
		OinchKey:  "",
		PriKey:    "",
		GasLegacy: gmgn_trade.GAS_PRICE_LEGACY_INSTANT,
		BuyNum:    0.1,
		Slippage:  1,
	}
	evmTradeTool := gmgn_trade.NewEvmTradeTool(config)
	if evmTradeTool == nil {
		return
	}

	tokenData, err := evmTradeTool.GetTokenData("0xcaa9ab6f25c6c5f3a997fc02788059b67f964444")
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(tokenData)
	// balanceStr, err := evmTradeTool.GetTokenBalance(gmgn_trade.MAIN_ETH20_ADDRESS)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(balanceStr)
	// tokenOut := "0x23D3F4EaaA515403C6765bb623F287a8Cca28F2b"
	// ethDecimals := solana.DecimalsInBigInt(uint32(gmgn_define.ETH20_MAIN_DECIMALS))
	// amount := big.NewInt(int64(0.01 * float64(ethDecimals.Int64()))) // 0.01个BNB
	// slippage := float32(1)                                           // 滑点1%
	// hash, err := evmTradeTool.Swap(gmgn_define.ETH20_MAIN_ADDRESS, tokenOut, amount, slippage)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("hash = ", hash)
	// SellToken(evmTradeTool, tokenOut)
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
	hash, err := evmTradeTool.Swap(tokenAddress, gmgn_define.ETH20_MAIN_ADDRESS, amount, slippage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hash = ", hash)
}
