package eth_trade

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Uniswap V3 SwapRouter ABI (仅包含 exactInputSingle 函数的简化版本)
const uniswapV3SwapRouterABI = `[
	{
		"inputs": [
			{
				"components": [
					{"name": "tokenIn", "type": "address"},
					{"name": "tokenOut", "type": "address"},
					{"name": "fee", "type": "uint24"},
					{"name": "recipient", "type": "address"},
					{"name": "deadline", "type": "uint256"},
					{"name": "amountIn", "type": "uint256"},
					{"name": "amountOutMinimum", "type": "uint256"},
					{"name": "sqrtPriceLimitX96", "type": "uint160"}
				],
				"internalType": "struct ISwapRouter.ExactInputSingleParams",
				"name": "params",
				"type": "tuple"
			}
		],
		"name": "exactInputSingle",
		"outputs": [
			{"name": "amountOut", "type": "uint256"}
		],
		"stateMutability": "payable",
		"type": "function"
	}
]`

type ExactInputSingleParams struct {
	TokenIn           string `json:"tokenIn"`
	TokenOut          string `json:"tokenOut"`
	Fee               int64  `json:"fee"`
	Recipient         string `json:"recipient"`
	Deadline          int64  `json:"deadline"`
	AmountIn          string `json:"amountIn"`
	AmountOutMinimum  string `json:"amountOutMinimum"`
	SqrtPriceLimitX96 string `json:"sqrtPriceLimitX96"`
}

// TradeConfig 定义交易配置
type TradeConfig struct {
	AmountIn     *big.Int       // 输入 ETH 数量 (Wei)
	TokenAddress common.Address // 目标 ERC20 代币地址
	Fee          int64          // 池费用 (e.g., 3000 for 0.3%)
	Slippage     int64          // 滑点保护百分比 (e.g., 95 表示 5% 滑点)
	Interval     time.Duration  // 交易检查间隔
}

type EthTradeTool struct {
	ethClient   *ethclient.Client
	routerAddr  common.Address
	abi         abi.ABI
	fromAddress common.Address
	privateKey  *ecdsa.PrivateKey
	chainID     *big.Int
}

func NewEthTradeTool(pubKey string, priKey string) *EthTradeTool {
	if pubKey == "" || priKey == "" {
		fmt.Println("pubKey == nil || priKey == nil")
		return nil
	}

	ethClient, err := ethclient.Dial(ETH_RCP_URL)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 解析 Uniswap V3 SwapRouter ABI
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV3SwapRouterABI))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ret := &EthTradeTool{
		ethClient:   ethClient,
		routerAddr:  common.HexToAddress(ETH_ROUTER_ADDRESS),
		abi:         parsedABI,
		fromAddress: common.HexToAddress(pubKey),
		privateKey:  privateKey,
		chainID:     chainID,
	}
	return ret
}

// ExecuteTrade 执行单次交易
func (et *EthTradeTool) Swap(inAddress string, outAddress string, amount *big.Int, fee int64, slippage int64) bool {
	auth, err := bind.NewKeyedTransactorWithChainID(et.privateKey, et.chainID)
	if err != nil {
		fmt.Println("Swap 1")
		fmt.Println(err)
		return false
	}
	auth.GasPrice, err = et.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Swap 2")
		fmt.Println(err)
		return false
	}
	auth.GasLimit = 350000 // Uniswap V3 交易通常需要更多 Gas
	auth.Value = amount    // 设置发送的 ETH 数量

	// 简化的价格估算（生产环境应使用 Quoter 合约）
	// 这里使用 amountIn 的比例作为 amountOutMin 的粗略估计
	amountOutMin := new(big.Int).Mul(amount, big.NewInt(slippage))
	amountOutMin.Div(amountOutMin, big.NewInt(100)) // 应用 5% 滑点保护

	// 设置交易截止时间（10 分钟后，转换为 UTC）
	deadline := big.NewInt(time.Now().UTC().Add(10 * time.Minute).Unix())
	abiParams := struct {
		TokenIn           common.Address
		TokenOut          common.Address
		Fee               *big.Int
		Recipient         common.Address
		Deadline          *big.Int
		AmountIn          *big.Int
		AmountOutMinimum  *big.Int
		SqrtPriceLimitX96 *big.Int
	}{
		TokenIn:           common.HexToAddress(inAddress),
		TokenOut:          common.HexToAddress(outAddress),
		Fee:               big.NewInt(fee),
		Recipient:         et.fromAddress,
		Deadline:          deadline,
		AmountIn:          amount,
		AmountOutMinimum:  amountOutMin,
		SqrtPriceLimitX96: big.NewInt(0),
	}

	// 编码函数调用
	data, err := et.abi.Pack("exactInputSingle", abiParams)
	if err != nil {
		fmt.Println("Swap 3")
		fmt.Println(err)
		return false
	}

	nonce, err := et.ethClient.PendingNonceAt(context.Background(), et.fromAddress)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 构造交易
	tx := types.NewTransaction(
		nonce,
		et.routerAddr,
		amount,
		auth.GasLimit,
		auth.GasPrice,
		data,
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(et.chainID), et.privateKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 发送交易
	err = et.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
