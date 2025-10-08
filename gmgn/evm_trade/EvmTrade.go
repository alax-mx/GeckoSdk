package evm_trade

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmTrade struct {
	client *ethclient.Client
	config *EvmConfig
	trader *UniswapTrader
	wallet *Wallet
}

func NewEvmTrade(walletPrivateKey string, rcpURL string, chainID int64, routerAddress string) *EvmTrade {
	client, err := ethclient.Dial(rcpURL)
	if err != nil {
		fmt.Println("Dial err:", err)
		return nil
	}
	config := &EvmConfig{
		PrivateKey:    walletPrivateKey,
		RPCURL:        rcpURL,
		ChainID:       chainID,
		UniswapRouter: routerAddress,
		Slippage:      0.5, // 默认滑点0.5%
	}

	wallet, err := NewWallet(client, walletPrivateKey)
	if err != nil {
		fmt.Println("NewWallet err:", err)
		return nil
	}

	trader, err := NewUniswapTrader(client, config, wallet)
	if err != nil {
		fmt.Println("NewUniswapTrader err:", err)
		return nil
	}
	return &EvmTrade{
		client: client,
		config: config,
		wallet: wallet,
		trader: trader,
	}
}

func (et *EvmTrade) BuyToken(tokenAddress string, amountEth float64) (string, error) {
	return et.trader.Buy(tokenAddress, amountEth)
}

func (et *EvmTrade) SellToken(tokenAddress string, amountToken float64) (string, error) {
	return et.trader.Sell(tokenAddress, amountToken)
}
