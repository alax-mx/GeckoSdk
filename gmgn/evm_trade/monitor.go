package evm_trade

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PriceMonitor struct {
	client    *ethclient.Client
	router    *UniswapV2Router
	pairCache map[string]*big.Float
}

func NewPriceMonitor(rpcURL string, routerAddress common.Address) (*PriceMonitor, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	router, err := NewUniswapV2Router(routerAddress)
	if err != nil {
		return nil, err
	}

	return &PriceMonitor{
		client:    client,
		router:    router,
		pairCache: make(map[string]*big.Float),
	}, nil
}

// GetPriceWithPath 通过指定路径获取价格
func (pm *PriceMonitor) GetPriceWithPath(amountIn *big.Int, path []common.Address) (*big.Float, error) {
	data, err := pm.router.ABI.Pack("getAmountsOut", amountIn, path)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &pm.router.Address,
		Data: data,
	}

	result, err := pm.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	var amounts []*big.Int
	err = pm.router.ABI.UnpackIntoInterface(&amounts, "getAmountsOut", result)
	if err != nil {
		return nil, err
	}

	if len(amounts) < 2 {
		return nil, fmt.Errorf("invalid amounts array length")
	}

	price := new(big.Float).Quo(
		new(big.Float).SetInt(amounts[len(amounts)-1]),
		new(big.Float).SetInt(amountIn),
	)

	return price, nil
}

// MonitorPriceChanges 监控价格变化
func (pm *PriceMonitor) MonitorPriceChanges(path []common.Address, interval time.Duration, callback func(price *big.Float)) {
	ticker := time.NewTicker(interval)
	baseAmount := big.NewInt(1e18) // 1 ETH or 1 token with 18 decimals

	for range ticker.C {
		price, err := pm.GetPriceWithPath(baseAmount, path)
		if err != nil {
			log.Printf("获取价格失败: %v", err)
			continue
		}

		callback(price)
	}
}
