package evm_trade

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapTrader struct {
	client *ethclient.Client
	config *EvmConfig
	wallet *Wallet
	router *UniswapV2Router
}

func NewUniswapTrader(client *ethclient.Client, cfg *EvmConfig, wallet *Wallet) (*UniswapTrader, error) {
	router, err := NewUniswapV2Router(common.HexToAddress(cfg.UniswapRouter))
	if err != nil {
		return nil, err
	}

	return &UniswapTrader{
		client: client,
		config: cfg,
		wallet: wallet,
		router: router,
	}, nil
}

// getAmountsOut 调用Uniswap Router的getAmountsOut函数
func (ut *UniswapTrader) getAmountsOut(path []common.Address, amountIn *big.Int) ([]*big.Int, error) {
	// 打包调用数据
	data, err := ut.router.ABI.Pack("getAmountsOut", amountIn, path)
	if err != nil {
		return nil, err
	}

	// 调用合约
	msg := ethereum.CallMsg{
		To:   &ut.router.Address,
		Data: data,
	}

	result, err := ut.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}

	// 解析返回结果
	var amounts []*big.Int
	err = ut.router.ABI.UnpackIntoInterface(&amounts, "getAmountsOut", result)
	if err != nil {
		return nil, err
	}

	return amounts, nil
}

// Buy 执行买入交易 (ETH -> Token)
func (ut *UniswapTrader) Buy(tokenAddress string, amountEth float64) (string, error) {
	// 转换金额为wei
	amountWei := new(big.Int)
	big.NewFloat(amountEth).Mul(big.NewFloat(amountEth), big.NewFloat(1e18)).Int(amountWei)

	path := []common.Address{
		common.HexToAddress(ETH_ADDRESS),  // WETH
		common.HexToAddress(tokenAddress), // 目标代币
	}

	// 获取预计输出数量
	amounts, err := ut.getAmountsOut(path, amountWei)
	if err != nil {
		return "", err
	}

	if len(amounts) < 2 {
		return "", fmt.Errorf("invalid amounts array length")
	}

	// 计算考虑滑点的最小输出数量
	minAmountOut := ut.calculateMinAmountOut(amounts[len(amounts)-1])

	// 构建交易路径
	// path := ut.getTradePath()
	// 获取交易授权
	auth, err := ut.wallet.GetAuth(big.NewInt(ut.config.ChainID))
	if err != nil {
		return "", err
	}

	auth.Value = amountWei
	auth.GasLimit = uint64(300000) // 适当的gas limit

	// 设置交易截止时间 (20分钟后)
	deadline := big.NewInt(time.Now().Add(20 * time.Minute).Unix())

	// 打包交易数据
	data, err := ut.router.ABI.Pack("swapExactETHForTokens", minAmountOut, path, ut.wallet.GetAddress(), deadline)
	if err != nil {
		return "", err
	}

	// 发送交易
	tx := types.NewTransaction(
		auth.Nonce.Uint64(),
		ut.router.Address,
		amountWei,
		auth.GasLimit,
		auth.GasPrice,
		data,
	)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", err
	}

	err = ut.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// Sell 执行卖出交易 (Token -> ETH)
func (ut *UniswapTrader) Sell(tokenAddress string, amountToken float64) (string, error) {
	// 转换金额为代币的最小单位
	amountIn := new(big.Int)
	big.NewFloat(amountToken).Mul(big.NewFloat(amountToken), big.NewFloat(1e18)).Int(amountIn)

	// 构建交易路径
	path := []common.Address{
		common.HexToAddress(tokenAddress), // 目标代币
		common.HexToAddress(ETH_ADDRESS),  // WETH
	}
	// 获取预计输出数量
	amounts, err := ut.getAmountsOut(path, amountIn)
	if err != nil {
		return "", err
	}

	if len(amounts) < 2 {
		return "", fmt.Errorf("invalid amounts array length")
	}

	// 计算考虑滑点的最小输出数量
	minAmountOut := ut.calculateMinAmountOut(amounts[len(amounts)-1])

	// 获取交易授权
	auth, err := ut.wallet.GetAuth(big.NewInt(ut.config.ChainID))
	if err != nil {
		return "", err
	}

	auth.GasLimit = uint64(300000)

	// 设置交易截止时间
	deadline := big.NewInt(time.Now().Add(20 * time.Minute).Unix())

	// 首先需要授权Router合约使用我们的代币
	err = ut.approveTokenIfNeeded(amountIn, auth)
	if err != nil {
		return "", err
	}

	// 打包交易数据
	data, err := ut.router.ABI.Pack("swapExactTokensForETH", amountIn, minAmountOut, path, ut.wallet.GetAddress(), deadline)
	if err != nil {
		return "", err
	}

	// 发送交易
	tx := types.NewTransaction(
		auth.Nonce.Uint64(),
		ut.router.Address,
		big.NewInt(0),
		auth.GasLimit,
		auth.GasPrice,
		data,
	)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", err
	}

	err = ut.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// calculateMinAmountOut 计算考虑滑点的最小输出数量
func (ut *UniswapTrader) calculateMinAmountOut(amountOut *big.Int) *big.Int {
	// 应用滑点容忍度
	slippageFactor := big.NewInt(10000 - int64(ut.config.Slippage*100))
	minAmountOut := new(big.Int).Mul(amountOut, slippageFactor)
	minAmountOut = new(big.Int).Div(minAmountOut, big.NewInt(10000))
	return minAmountOut
}

// approveTokenIfNeeded 授权Router合约使用代币
func (ut *UniswapTrader) approveTokenIfNeeded(amount *big.Int, auth *bind.TransactOpts) error {
	// 这里需要实现ERC20 approve调用
	// 检查当前授权额度，如果不足则进行授权

	// 简化实现：总是进行授权
	// 实际项目中应该检查当前授权额度
	return nil
}
