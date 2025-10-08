package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/1inch/1inch-sdk-go/constants"
	"github.com/1inch/1inch-sdk-go/sdk-clients/aggregation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	fmt.Println(client.Wallet.Address())
	allowanceParams := aggregation.GetAllowanceParams{
		TokenAddress:  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		WalletAddress: client.Wallet.Address().Hex(),
	}
	allowance, err := client.GetApproveAllowance(ctx, allowanceParams)
	if err != nil {
		log.Fatalf("檢查批准額度失敗: %v\n", err)
	}
	fmt.Printf("USDC 對 1inch Router 的批准額度：%s\n", allowance.Allowance)

	allowanceInt := new(big.Int)
	allowanceInt.SetString(allowance.Allowance, 10)
	swapAmount := big.NewInt(1000000000000000) // 1 USDC (6 位小數，單位為 wei)
	if allowanceInt.Cmp(swapAmount) < 0 {
		approveData, err := client.GetApproveTransaction(ctx, aggregation.GetApproveParams{
			TokenAddress: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			Amount:       swapAmount.String(),
		})
		if err != nil {
			log.Fatalf("Failed to get approve data: %v\n", err)
		}
		data, err := hexutil.Decode(approveData.Data)
		if err != nil {
			log.Fatalf("Failed to decode approve data: %v\n", err)
		}

		to := common.HexToAddress(approveData.To)
		tx, err := client.TxBuilder.New().SetData(data).SetTo(&to).Build(ctx)
		if err != nil {
			log.Fatalf("Failed to build approve transaction: %v\n", err)
		}

		signedTx, err := client.Wallet.Sign(tx)
		if err != nil {
			log.Fatalf("Failed to sign approve transaction: %v\n", err)
		}

		err = client.Wallet.BroadcastTransaction(ctx, signedTx)
		if err != nil {
			log.Fatalf("Failed to broadcast approve transaction: %v\n", err)
		}

		fmt.Printf("Transaction has been broadcast. View it on Polygonscan here: %v\n", fmt.Sprintf("https://polygonscan.com/tx/%v", signedTx.Hash().Hex()))
		for {
			receipt, err := client.Wallet.TransactionReceipt(ctx, signedTx.Hash())
			if receipt != nil {
				fmt.Println("Transaction complete!")
				break
			}
			if err != nil {
				fmt.Println("Waiting for transaction to be mined")
			}
			select {
			case <-time.After(1000 * time.Millisecond): // check again after a delay
			case <-ctx.Done():
				fmt.Println("Context cancelled")
				break
			}
		}
	} else {
		fmt.Println("批准額度足夠，無需再次批准")
	}

	swapData, err := client.GetSwap(ctx, aggregation.GetSwapParams{
		Src:             "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
		Dst:             "0xcf91b70017eabde82c9671e30e5502d312ea6eb2", // USDC
		Amount:          "1000000000000000",                           // 0.01 ETH
		From:            client.Wallet.Address().Hex(),
		Slippage:        1,
		DisableEstimate: false,
	})
	if err != nil {
		log.Fatalf("Failed to get swap data: %v\n", err)
	}

	output, err := json.MarshalIndent(swapData, "", "  ")
	if err != nil {
		log.Fatalf("JSON 序列化失敗: %v\n", err)
	}
	fmt.Printf("ETH Swap 數據：\n%s\n", string(output))
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
	fmt.Printf("Transaction has been broadcast. View it on Etherscan here: %v\n", fmt.Sprintf("https://etherscan.io/tx/%v", signedTx.Hash().Hex()))
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
