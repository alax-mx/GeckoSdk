package gmgn_trade

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/1inch/1inch-sdk-go/constants"
	"github.com/1inch/1inch-sdk-go/sdk-clients/aggregation"
	"github.com/1inch/1inch-sdk-go/sdk-clients/balances"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alax-mx/geckosdk/baseutils"
)

type EvmTradeTool struct {
	client        *aggregation.Client
	balanceClient *balances.Client
	ethClient     *ethclient.Client
	ctx           context.Context
}

func NewEvmTradeTool(evmConfig *STEvmConfig) *EvmTradeTool {
	param := aggregation.ConfigurationParams{
		NodeUrl:    evmConfig.RpcURL,
		PrivateKey: evmConfig.PriKey,
		ApiUrl:     "https://api.1inch.dev",
		ApiKey:     evmConfig.OinchKey,
	}
	switch evmConfig.ChainType {
	case CHAIN_TYPE_ETH:
		param.ChainId = constants.EthereumChainId
	case CHAIN_TYPE_BSC:
		param.ChainId = constants.BscChainId
	case CHAIN_TYPE_BASE:
		param.ChainId = constants.BaseChainId
	case CHAIN_TYPE_POLYGON:
		param.ChainId = constants.PolygonChainId
	default:
		fmt.Println("NewEvmTradeTool err: unknow chaintype ", evmConfig.ChainType)
		return nil
	}

	config, err := aggregation.NewConfiguration(param)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create configuration ", err)
		return nil
	}

	client, err := aggregation.NewClient(config)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create client ", err)
		return nil
	}
	ethClient, err := ethclient.Dial(evmConfig.RpcURL)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create eth client ", err)
		return nil
	}

	balanceConfig, err := balances.NewConfiguration(balances.ConfigurationParams{
		ChainId: param.ChainId,
		ApiUrl:  "https://api.1inch.dev",
		ApiKey:  evmConfig.OinchKey,
	})
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to NewConfiguration client ", err)
		return nil
	}
	balanceClient, err := balances.NewClient(balanceConfig)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
		fmt.Println("NewEvmTradeTool err: Failed to create balanceClient ", err)
		return nil
	}

	return &EvmTradeTool{
		client:        client,
		ctx:           context.Background(),
		ethClient:     ethClient,
		balanceClient: balanceClient,
	}
}

func (ett *EvmTradeTool) Swap(tokenIn string, tokenOut string, amount *big.Int, slippage float32) (common.Hash, error) {
	swapParams := aggregation.GetSwapParams{
		Src:             tokenIn,
		Dst:             tokenOut,
		Amount:          amount.String(),
		From:            ett.client.Wallet.Address().Hex(),
		Slippage:        slippage,
		DisableEstimate: false,
	}

	swapData, err := ett.client.GetSwap(ett.ctx, swapParams)
	if err != nil {
		return common.Hash{}, errors.New("StartSwap: Failed to get swap data: " + err.Error())
	}
	tx, err := ett.client.TxBuilder.New().SetData(swapData.TxNormalized.Data).
		SetTo(&swapData.TxNormalized.To).
		SetGas(1000000).
		SetValue(swapData.TxNormalized.Value).
		Build(ett.ctx)
	fmt.Println("swapData.TxNormalized.Gas = ", swapData.TxNormalized.GasPrice)
	if err != nil {
		return common.Hash{}, errors.New("StartSwap: Failed to build transaction: " + err.Error())
	}
	signedTx, err := ett.client.Wallet.Sign(tx)
	if err != nil {
		return common.Hash{}, errors.New("StartSwap: Failed to sign transaction: " + err.Error())
	}

	err = ett.client.Wallet.BroadcastTransaction(ett.ctx, signedTx)
	if err != nil {
		return common.Hash{}, errors.New("StartSwap: Failed to broadcast transaction: " + err.Error())
	}

	return signedTx.Hash(), nil
}

// CheckTokenAllAllowance 检测token是否所有额度都批准了
func (ett *EvmTradeTool) CheckTokenAllAllowance(tokenAddress string) (bool, error) {
	balanceStr, err := ett.GetTokenBalance(tokenAddress)
	if err != nil {
		return false, err
	}

	allowanceInt := new(big.Int)
	allowanceInt.SetString(balanceStr, 10)
	return ett.CheckAllowance(tokenAddress, allowanceInt)
}

// CheckTokenAllAllowance 批准token所有额度
func (ett *EvmTradeTool) MakeTokenApporveAll(tokenAddress string) error {
	balanceStr, err := ett.GetTokenBalance(tokenAddress)
	if err != nil {
		return err
	}

	allowanceInt := new(big.Int)
	allowanceInt.SetString(balanceStr, 10)
	_, err = ett.Approve(tokenAddress, allowanceInt)
	if err != nil {
		return err
	}
	return nil
}

// CheckAllowance 检查批准额度
func (ett *EvmTradeTool) CheckAllowance(tokenIn string, amount *big.Int) (bool, error) {
	allowanceParams := aggregation.GetAllowanceParams{
		TokenAddress:  tokenIn,
		WalletAddress: ett.client.Wallet.Address().Hex(),
	}
	allowance, err := ett.client.GetApproveAllowance(ett.ctx, allowanceParams)
	if err != nil {
		return false, err
	}

	allowanceInt := new(big.Int)
	allowanceInt.SetString(allowance.Allowance, 10)
	if allowanceInt.Cmp(amount) < 0 {
		return false, nil
	}
	return true, nil
}

func (ett *EvmTradeTool) GetPermit(tokenIn string, amount *big.Int) (string, error) {
	spender, err := ett.client.GetApproveSpender(ett.ctx)
	if err != nil {
		return "", errors.New("GetPermit err: " + err.Error())
	}
	fmt.Println("spender = ", spender)
	now := time.Now()
	twoDaysLater := now.Add(time.Hour * 24 * 2)
	permitData, err := ett.client.Wallet.GetContractDetailsForPermit(ett.ctx,
		common.HexToAddress(tokenIn),
		common.HexToAddress(spender.Address),
		amount, twoDaysLater.Unix())
	if err != nil {
		return "", errors.New("GetPermit err: Failed to get permit data:" + err.Error())
	}

	permit, err := ett.client.Wallet.TokenPermit(*permitData)
	if err != nil {
		return "", errors.New("GetPermit err: Failed to sign permit:" + err.Error())
	}

	return permit, nil
}

// Approve 申请批准额度
func (ett *EvmTradeTool) Approve(tokenIn string, amount *big.Int) (common.Hash, error) {
	approveData, err := ett.client.GetApproveTransaction(ett.ctx, aggregation.GetApproveParams{
		TokenAddress: tokenIn,
		Amount:       amount.String(),
	})
	if err != nil {
		return common.Hash{}, errors.New("Failed to get approve data: " + err.Error())
	}
	data, err := hexutil.Decode(approveData.Data)
	if err != nil {
		return common.Hash{}, errors.New("Failed to decode approve data: " + err.Error())
	}

	to := common.HexToAddress(approveData.To)
	tx, err := ett.client.TxBuilder.New().SetData(data).SetTo(&to).SetGas(1000000).Build(ett.ctx)
	if err != nil {
		return common.Hash{}, errors.New("Failed to build approve transaction: " + err.Error())
	}

	signedTx, err := ett.client.Wallet.Sign(tx)
	if err != nil {
		return common.Hash{}, errors.New("Failed to sign approve transaction: " + err.Error())
	}

	err = ett.client.Wallet.BroadcastTransaction(ett.ctx, signedTx)
	if err != nil {
		return common.Hash{}, errors.New("Failed to broadcast approve transaction: " + err.Error())
	}

	return signedTx.Hash(), nil
}

// TransactionReceipt 查询交易hash状态
func (ett *EvmTradeTool) TransactionReceipt(hash common.Hash) {
	for {
		receipt, err := ett.client.Wallet.TransactionReceipt(ett.ctx, hash)
		if receipt != nil {
			fmt.Println("Transaction complete!")
			break
		}
		if err != nil {
			fmt.Println("Waiting for transaction to be mined: ", err, " hash = ", hash)
		}
		time.Sleep(1 * time.Second)
	}
}

func (ett *EvmTradeTool) CheckPermitSupport(tokenAddr, owner common.Address) error {
	// 代币 ABI，仅包含 nonces 和 name
	tokenABI, err := abi.JSON(strings.NewReader(`[
        {"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"nonces","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
        {"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"stateMutability":"view","type":"function"}
    ]`))
	if err != nil {
		return err
	}

	contract := bind.NewBoundContract(tokenAddr, tokenABI, ett.ethClient, ett.ethClient, ett.ethClient)

	// 查询 nonce
	var nonceResults []interface{}
	err = contract.Call(nil, &nonceResults, "nonces", owner)
	if err != nil {
		return err
	}
	if len(nonceResults) == 0 {
		return errors.New("nonceResults 返回为null")
	}

	// 查询 name
	var nameResults []interface{}
	err = contract.Call(nil, &nameResults, "name")
	if err != nil {
		return err
	}
	if len(nameResults) == 0 {
		return errors.New("nameResults 返回为null")
	}

	return nil
}

func (ett *EvmTradeTool) GetTokenBalance(tokenAddress string) (string, error) {
	response, err := ett.balanceClient.GetBalancesOfCustomTokensByWalletAddress(ett.ctx, balances.BalancesOfCustomTokensByWalletAddressParams{
		Wallet: ett.client.Wallet.Address().String(),
		Tokens: []string{tokenAddress},
	})
	if err != nil {
		return "", err
	}

	retMap := *response
	return retMap[baseutils.ToLowerHex(tokenAddress)], nil
}
