package gmgn_trade

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/alax-mx/geckosdk/proxy"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagliardetto/solana-go"
)

// Uniswap V3 SwapRouter ABI（仅包含 exactInputSingle 方法）
const GSwapRouterABI = `[{"inputs":[{"components":[{"internalType":"address","name":"tokenIn","type":"address"},{"internalType":"address","name":"tokenOut","type":"address"},{"internalType":"uint24","name":"fee","type":"uint24"},{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMinimum","type":"uint256"},{"internalType":"uint160","name":"sqrtPriceLimitX96","type":"uint160"}],"internalType":"struct ISwapRouter.ExactInputSingleParams","name":"params","type":"tuple"}],"name":"exactInputSingle","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"payable","type":"function"}]`
const contractABI = `[{"inputs":[{"internalType":"uint256","name":"ethAmount","type":"uint256"}],"name":"buyTokens","outputs":[],"stateMutability":"payable","type":"function"}]`

// ERC20 approve 函数的 ABI
const Erc20ApproveABI = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

// ERC20 allowance 函数的 ABI
const Erc20AllowanceABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

type ETHTadeTool2 struct {
	rpcURL string
	pubKey string
	priKey string

	proxy *proxy.ProxyTool
}

func NewETHTadeTool2(rpcURL string, pubKey string, priKey string, proxy *proxy.ProxyTool) *ETHTadeTool2 {
	return &ETHTadeTool2{
		rpcURL: rpcURL,
		pubKey: pubKey,
		priKey: priKey,
		proxy:  proxy,
	}
}

func (tool *ETHTadeTool2) Swap(inAddress string, outAddress string, amount float64) {
	// Connect to Sepolia RPC (replace with your Infura/Alchemy key)
	client, err := ethclient.Dial(tool.rpcURL)
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer client.Close()

	// Load private key from env
	privateKey, err := crypto.HexToECDSA("5af616d5aad4b117da457cf8719ee4b6a113dbbe987f431e72508c0549834aae")
	if err != nil {
		fmt.Println("Invalid private key")
		return
	}

	// 获取账户地址
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	// 智能合约地址
	contractAddress := common.HexToAddress("0xcf91b70017eabde82c9671e30e5502d312ea6eb2")

	contractABI2, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建合约实例
	contract := bind.NewBoundContract(contractAddress, contractABI2, client, client, client)
	if err != nil {
		fmt.Println("create contract err:", err)
	}

	// 设置购买的 ETH 数量（以 wei 为单位）
	// ethAmount := big.NewInt(1e18) // 1 ETH = 10^18 wei
	solDecimals := solana.DecimalsInBigInt(uint32(ETH_DECIMALS))
	ethAmount := big.NewInt(int64(0.005 * float64(solDecimals.Int64())))
	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// 设置 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	fmt.Println("gasPrice = ", gasPrice.Int64())

	// 创建交易授权
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // 1 是链 ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = ethAmount         // 发送的 ETH 数量
	auth.GasLimit = uint64(300000) // Gas 限制（根据合约复杂性调整）
	auth.GasPrice = gasPrice

	// 调用购买函数
	tx, err := contract.Transact(auth, "buyTokens", ethAmount)
	if err != nil {
		log.Fatalf("Failed to execute buyTokens: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to mine transaction: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("Transaction successful!")
	} else {
		fmt.Println("Transaction failed!")
	}
}

func (tool *ETHTadeTool2) Allowance(tokenAddress string) (*big.Int, error) {
	// 连接区块链客户端
	client, err := ethclient.Dial(tool.rpcURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// 配置地址
	ownerAddress := common.HexToAddress(tool.pubKey)             // 你的钱包地址
	spenderAddress := common.HexToAddress(ETH_ROUTER_V3_ADDRESS) // Uniswap V3 Router
	tokenAddress2 := common.HexToAddress(tokenAddress)           // 代币合约地址（例如 USDT）
	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(Erc20AllowanceABI))
	if err != nil {
		return nil, err
	}

	// 打包 approve 函数输入
	input, err := parsedABI.Pack("allowance", ownerAddress, spenderAddress)
	if err != nil {
		return nil, err
	}

	// 调用合约
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress2,
		Data: input,
	}, nil)
	if err != nil {
		return nil, err
	}

	// 解包返回结果
	outputs, err := parsedABI.Unpack("allowance", result)
	if err != nil {
		return nil, err
	}

	// 输出授权金额
	allowance := outputs[0].(*big.Int)
	return allowance, nil
}

func (tool *ETHTadeTool2) Approve(tokenAddress string, amount *big.Int) error {
	// 连接区块链客户端
	client, err := ethclient.Dial(tool.rpcURL)
	if err != nil {
		return err
	}
	defer client.Close()

	// Load private key from env
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(tool.priKey, "0x"))
	if err != nil {
		fmt.Println("Invalid private key")
		return err
	}

	// 获取账户地址
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	// 配置地址
	spenderAddress := common.HexToAddress(ETH_ROUTER_V3_ADDRESS) // Uniswap V3 Router
	tokenAddress2 := common.HexToAddress(tokenAddress)           // 代币合约地址（例如 USDT）

	// 解析 ABI
	parsedABI, err := abi.JSON(strings.NewReader(Erc20ApproveABI))
	if err != nil {
		return err
	}

	// 打包 approve 函数输入
	input, err := parsedABI.Pack("approve", spenderAddress, amount)
	if err != nil {
		return err
	}

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	// 设置 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// 创建交易授权
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // 1 是链 ID
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // 不发送 ETH
	auth.GasLimit = uint64(10000) // Gas 限制（根据合约复杂性调整）
	auth.GasPrice = gasPrice      // 使用建议的 gas 价格
	// 创建并签署交易
	tx := types.NewTransaction(auth.Nonce.Uint64(), tokenAddress2, big.NewInt(0), auth.GasLimit, auth.GasPrice, input)
	signedTx, err := auth.Signer(fromAddress, tx)
	if err != nil {
		return errors.New("无法签署交易: " + err.Error())
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return errors.New("无法发送交易: " + err.Error())
	}
	fmt.Printf("Approve 交易已发送: %s\n", signedTx.Hash().Hex())
	fmt.Printf("在 Etherscan 查看: https://etherscan.io/tx/%s\n", signedTx.Hash().Hex())

	// 可选：等待交易确认
	_, err = bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		return errors.New("交易确认失败: " + err.Error())
	}
	fmt.Println("Approve 交易已确认")
	return nil
}
