package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/1inch/1inch-sdk-go/constants"
	"github.com/1inch/1inch-sdk-go/sdk-clients/aggregation"
)

func main() {
	TestBuy()
}

func TestBuy() {
	rpcUrl := ""
	randomPrivateKey := ""

	config, err := aggregation.NewConfiguration(aggregation.ConfigurationParams{
		NodeUrl:    rpcUrl,
		PrivateKey: randomPrivateKey,
		ChainId:    constants.EthereumChainId,
		ApiUrl:     "https://api.1inch.dev",
		ApiKey:     "",
	})
	if err != nil {
		log.Fatalf("Failed to create configuration: %v\n", err)
	}
	client, err := aggregation.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}
	ctx := context.Background()

	swapData, err := client.GetSwap(ctx, aggregation.GetSwapParams{
		Src:             "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
		Dst:             "0xcf91b70017eabde82c9671e30e5502d312ea6eb2", // 1INCH
		Amount:          "1000000000000000",                           // 0.01 ETH
		From:            client.Wallet.Address().Hex(),
		Slippage:        1,
		DisableEstimate: true,
	})
	if err != nil {
		log.Fatalf("Failed to get swap data: %v\n", err)
	}

	tx, err := client.TxBuilder.New().SetData(swapData.TxNormalized.Data).SetTo(&swapData.TxNormalized.To).SetGas(swapData.TxNormalized.Gas).SetValue(swapData.TxNormalized.Value).Build(ctx)
	if err != nil {
		log.Fatalf("Failed to build transaction: %v\n", err)
	}
	signedTx, err := client.Wallet.Sign(tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v\n", err)
	}

	err = client.Wallet.BroadcastTransaction(ctx, signedTx)
	if err != nil {
		log.Fatalf("Failed to broadcast transaction: %v\n", err)
	}

	// Waiting for transaction, just an examples of it
	fmt.Printf("Transaction has been broadcast. View it on Polygonscan here: %v\n", fmt.Sprintf("https://polygonscan.com/tx/%v", signedTx.Hash().Hex()))
	for {
		receipt, err := client.Wallet.TransactionReceipt(ctx, signedTx.Hash())
		if receipt != nil {
			fmt.Println("Transaction complete!")
			return
		}
		if err != nil {
			fmt.Println("Waiting for transaction to be mined")
		}
		select {
		case <-time.After(1000 * time.Millisecond): // check again after a delay
		case <-ctx.Done():
			fmt.Println("Context cancelled")
			return
		}
	}
}
