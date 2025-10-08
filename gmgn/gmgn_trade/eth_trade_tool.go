package gmgn_trade

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alax-mx/geckosdk/proxy"
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

type STSteps struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Tool string `json:"tool"`
}
type STPathLiquidity struct {
	Pair      string  `json:"pair"`
	Liquidity float64 `json:"liquidity"`
}
type STErrInfo struct {
	Code  int    `json:"code"`
	Mcode string `json:"mcode"`
	Msg   string `json:"msg"`
}
type STRoutes struct {
	ChainID            int               `json:"chain_id"`
	Type               string            `json:"type"`
	To                 string            `json:"to"`
	AmountIn           string            `json:"amount_in"`
	AmountOut          string            `json:"amount_out"`
	AmountMinOut       string            `json:"amount_min_out"`
	InputTokenAddress  string            `json:"input_token_address"`
	OutputTokenAddress string            `json:"output_token_address"`
	Fee                int               `json:"fee"`
	Path               []string          `json:"path"`
	PoolAddress        string            `json:"pool_address"`
	FactoryAddress     string            `json:"factory_address"`
	Steps              []STSteps         `json:"steps"`
	TokenInUsdPrice    string            `json:"token_in_usd_price"`
	AmountInUsd        string            `json:"amount_in_usd"`
	TokenOutUsdPrice   string            `json:"token_out_usd_price"`
	AmountOutUsd       string            `json:"amount_out_usd"`
	PriceImpact        string            `json:"price_impact"`
	PathBytes          string            `json:"path_bytes"`
	Liquidity          float64           `json:"liquidity"`
	PathLiquidity      []STPathLiquidity `json:"path_liquidity"`
	GasLimit           string            `json:"gas_limit"`
	ErrInfo            STErrInfo         `json:"err_info"`
}
type STVolatilities struct {
	TokenIn  int  `json:"token_in"`
	TokenOut int  `json:"token_out"`
	IsFomo   bool `json:"is_fomo"`
}

type STData struct {
	Routes       []STRoutes     `json:"routes"`
	Volatilities STVolatilities `json:"volatilities"`
	TimeTaken    float64        `json:"timeTaken"`
}

type GetEthRouterResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data STData `json:"data"`
	Tid  string `json:"tid"`
}
type EthTradeTool struct {
	ethClient *ethclient.Client
	baseUrl   string
	pubKey    string
	priKey    string
	abi       abi.ABI
	proxyInfo *proxy.STProxyInfo
}

func NewEthTradeTool(baseUrl string, pubKey string, priKey string) *EthTradeTool {
	// 解析 Uniswap V3 SwapRouter ABI
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV3SwapRouterABI))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ethClient, err := ethclient.Dial("https://ethereum.publicnode.com")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &EthTradeTool{
		ethClient: ethClient,
		baseUrl:   baseUrl,
		pubKey:    pubKey,
		priKey:    priKey,
		abi:       parsedABI,
		proxyInfo: nil,
	}
}

func (ett *EthTradeTool) GetAvailableRouter(chainType string, inAddress string, outAddress string, amount int) (*GetEthRouterResp, error) {
	tmpUrl := "/available_routes_exact_in?"
	tmpUrl += "token_in_chain=" + chainType
	tmpUrl += "&from_address=" + ett.pubKey
	tmpUrl += "&token_out_chain=" + chainType
	tmpUrl += "&token_in_address=" + inAddress
	tmpUrl += "&token_out_address=" + outAddress
	tmpUrl += "&in_amount=" + strconv.Itoa(amount)
	data, err := HttpGet(ett.baseUrl+tmpUrl, ett.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetEthRouterResp{}
	err2 := json.Unmarshal(data, ret)
	if err2 != nil {
		return nil, err2
	}

	return ret, nil
}

func (ett *EthTradeTool) Swap(router *STRoutes, slipper float64) bool {
	amountInt, err := strconv.ParseInt(router.AmountIn, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}

	amountMinOut, err := strconv.ParseInt(router.AmountMinOut, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
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
		TokenIn:           common.HexToAddress(router.InputTokenAddress),
		TokenOut:          common.HexToAddress(router.OutputTokenAddress),
		Fee:               big.NewInt(int64(router.Fee)),
		Recipient:         common.HexToAddress(ett.pubKey),
		Deadline:          big.NewInt(time.Now().UTC().Add(10 * time.Minute).Unix()),
		AmountIn:          big.NewInt(amountInt),
		AmountOutMinimum:  big.NewInt(amountMinOut),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	data, err := ett.abi.Pack("exactInputSingle", abiParams)
	if err != nil {
		fmt.Println("Swap 3")
		fmt.Println(err)
		return false
	}

	nonce, err := ett.ethClient.PendingNonceAt(context.Background(), common.HexToAddress(ett.pubKey))
	if err != nil {
		fmt.Println(err)
		return false
	}

	routerAddr := common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")

	gasLimit, err := strconv.ParseInt(router.GasLimit, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 构造交易
	tx := types.NewTransaction(
		nonce,
		routerAddr,
		big.NewInt(amountInt),
		uint64(gasLimit),
		nil,
		data,
	)

	privateKey, err := crypto.HexToECDSA(ett.priKey)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(router.ChainID))), privateKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 发送交易
	err = ett.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
