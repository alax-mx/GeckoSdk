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
	"github.com/1inch/1inch-sdk-go/sdk-clients/gasprices"
	"github.com/1inch/1inch-sdk-go/sdk-clients/tokens"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
)

type EvmTradeTool struct {
	evmConfig     *STEvmConfig
	client        *aggregation.Client
	gasClient     *gasprices.Client
	balanceClient *balances.Client
	ethClient     *ethclient.Client
	tokenClient   *tokens.Client
	ctx           context.Context
}

func NewEvmTradeTool(evmConfig *STEvmConfig) *EvmTradeTool {
	param := aggregation.ConfigurationParams{
		NodeUrl:    evmConfig.RpcURL,
		PrivateKey: evmConfig.PriKey,
		ApiUrl:     "https://api.1inch.dev",
		ApiKey:     evmConfig.OinchKey,
	}

	chainID, err := GetChainIdByType(evmConfig.ChainType)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to get chain ID ", err)
		return nil
	}
	param.ChainId = uint64(chainID)

	// aggregation client
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

	// balance client
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

	// gasprices client
	gasClientConfig, err := gasprices.NewConfiguration(gasprices.ConfigurationParams{
		ChainId: param.ChainId, // 或其他链 ID，如 constants.PolygonChainId
		ApiUrl:  "https://api.1inch.dev",
		ApiKey:  evmConfig.OinchKey,
	})
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create gasprices configuration ", err)
		return nil
	}
	gasClient, err := gasprices.NewClient(gasClientConfig)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create gasprices client: ", err)
		return nil
	}

	tokenClientConfig, err := tokens.NewConfiguration(tokens.ConfigurationParams{
		ChainId: uint64(chainID),
		ApiUrl:  "https://api.1inch.dev",
		ApiKey:  evmConfig.OinchKey,
	})
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create token configuration ", err)
		return nil
	}
	tokenClient, err := tokens.NewClient(tokenClientConfig)
	if err != nil {
		fmt.Println("NewEvmTradeTool err: Failed to create token client: ", err)
		return nil
	}
	return &EvmTradeTool{
		evmConfig:     evmConfig,
		client:        client,
		ctx:           context.Background(),
		ethClient:     ethClient,
		balanceClient: balanceClient,
		gasClient:     gasClient,
		tokenClient:   tokenClient,
	}
}

func GetChainIdByType(chainType string) (int, error) {
	switch chainType {
	case gmgn_define.CHAIN_TYPE_ETH:
		return constants.EthereumChainId, nil
	case gmgn_define.CHAIN_TYPE_BSC:
		return constants.BscChainId, nil
	case gmgn_define.CHAIN_TYPE_BASE:
		return constants.BaseChainId, nil
	case gmgn_define.CHAIN_TYPE_POLYGON:
		return constants.PolygonChainId, nil
	default:
		return 0, errors.New("GetChainIdByType err: unknow chaintype " + chainType)
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

	// 如果支持 permit 则设置
	if tokenIn != gmgn_define.ETH20_MAIN_ADDRESS {
		spender, err := ett.client.GetApproveSpender(ett.ctx)
		if err == nil {
			now := time.Now()
			twoDaysLater := now.Add(time.Minute * 10)
			permitData, err := ett.client.Wallet.GetContractDetailsForPermit(ett.ctx, common.HexToAddress(tokenIn), common.HexToAddress(spender.Address), amount, twoDaysLater.Unix())
			if err == nil {
				permit, err := ett.client.Wallet.TokenPermit(*permitData)
				if err == nil {
					swapParams.Permit = permit
				}
			}
		}
	}

	swapData, err := ett.client.GetSwap(ett.ctx, swapParams)
	if err != nil {
		return common.Hash{}, errors.New("Swap: Failed to get swap data: " + err.Error())
	}

	builder := ett.client.TxBuilder.New()
	builder.SetData(swapData.TxNormalized.Data).
		SetTo(&swapData.TxNormalized.To).
		SetGas(1000000).
		SetValue(swapData.TxNormalized.Value)

	// 根据不同的链设置GasPrice
	switch ett.evmConfig.ChainType {
	case gmgn_define.CHAIN_TYPE_ETH:
		maxFeePerGas, maxPriorityFeePerGas, err := ett.GetGasLegacyEIP1559(ett.evmConfig.GasLegacy)
		if err != nil {
			return common.Hash{}, err
		}
		builder.SetGasFeeCap(maxFeePerGas).SetGasTipCap(maxPriorityFeePerGas)
	case gmgn_define.CHAIN_TYPE_BSC:
		gasPrice, err := ett.GetGasLegacy(ett.evmConfig.GasLegacy)
		if err != nil {
			return common.Hash{}, err
		}
		builder.SetGasPrice(gasPrice)
	case gmgn_define.CHAIN_TYPE_BASE:
		maxFeePerGas, maxPriorityFeePerGas, err := ett.GetGasLegacyEIP1559(ett.evmConfig.GasLegacy)
		if err != nil {
			return common.Hash{}, err
		}
		builder.SetGasFeeCap(maxFeePerGas).SetGasTipCap(maxPriorityFeePerGas)
	}

	tx, err := builder.Build(ett.ctx)
	if err != nil {
		return common.Hash{}, errors.New("Swap: Failed to build transaction: " + err.Error())
	}
	signedTx, err := ett.client.Wallet.Sign(tx)
	if err != nil {
		return common.Hash{}, errors.New("Swap: Failed to sign transaction: " + err.Error())
	}

	err = ett.client.Wallet.BroadcastTransaction(ett.ctx, signedTx)
	if err != nil {
		return common.Hash{}, errors.New("Swap: Failed to broadcast transaction: " + err.Error())
	}

	return signedTx.Hash(), nil
}

// GetGasByLegacy 获取gas优先级信息
func (ett *EvmTradeTool) GetGasLegacyEIP1559(legacy string) (*big.Int, *big.Int, error) {
	ctx := context.Background()
	gasPriceLegacy, err := ett.gasClient.GetGasPriceEIP1559(ctx)
	if err != nil {
		return nil, nil, errors.New("GetGasByLegacy err: " + err.Error())
	}

	maxFeePerGasStr := gasPriceLegacy.Low.MaxFeePerGas
	maxPriorityFeePerGasStr := gasPriceLegacy.Low.MaxPriorityFeePerGas
	switch legacy {
	case GAS_PRICE_LEGACY_MEDIUM:
		maxFeePerGasStr = gasPriceLegacy.Medium.MaxFeePerGas
		maxPriorityFeePerGasStr = gasPriceLegacy.Medium.MaxPriorityFeePerGas
	case GAS_PRICE_LEGACY_HIGH:
		maxFeePerGasStr = gasPriceLegacy.High.MaxFeePerGas
		maxPriorityFeePerGasStr = gasPriceLegacy.High.MaxPriorityFeePerGas
	case GAS_PRICE_LEGACY_INSTANT:
		maxFeePerGasStr = gasPriceLegacy.Instant.MaxFeePerGas
		maxPriorityFeePerGasStr = gasPriceLegacy.Instant.MaxPriorityFeePerGas
	}

	maxFeePerGas := new(big.Int)
	maxFeePerGas.SetString(maxFeePerGasStr, 10)
	maxPriorityFeePerGas := new(big.Int)
	maxPriorityFeePerGas.SetString(maxPriorityFeePerGasStr, 10)

	return maxFeePerGas, maxPriorityFeePerGas, nil
}

// GetGasByLegacy 获取gas优先级信息
func (ett *EvmTradeTool) GetGasLegacy(legacy string) (*big.Int, error) {
	ctx := context.Background()
	gasPriceLegacyResp, err := ett.gasClient.GetGasPriceLegacy(ctx)
	if err != nil {
		return nil, errors.New("GetGasLegacy err: " + err.Error())
	}

	gasPriceStr := gasPriceLegacyResp.Standard
	switch legacy {
	case GAS_PRICE_LEGACY_HIGH:
		gasPriceStr = gasPriceLegacyResp.Fast
	case GAS_PRICE_LEGACY_INSTANT:
		gasPriceStr = gasPriceLegacyResp.Instant
	}

	gasPrice := new(big.Int)
	gasPrice.SetString(gasPriceStr, 10)

	return gasPrice, nil
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
	_, err = ett.Approve(tokenAddress)
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

// Approve 申请批准额度
func (ett *EvmTradeTool) Approve(tokenIn string) (common.Hash, error) {
	approveData, err := ett.client.GetApproveTransaction(ett.ctx, aggregation.GetApproveParams{
		TokenAddress: tokenIn,
	})
	if err != nil {
		return common.Hash{}, errors.New("Approve: Failed to get approve data: " + err.Error())
	}
	data, err := hexutil.Decode(approveData.Data)
	if err != nil {
		return common.Hash{}, errors.New("Approve: Failed to decode approve data: " + err.Error())
	}

	to := common.HexToAddress(approveData.To)
	builder := ett.client.TxBuilder.New()
	builder.SetData(data).SetTo(&to).SetGas(1000000)

	switch ett.evmConfig.ChainType {
	case gmgn_define.CHAIN_TYPE_ETH:
		maxFeePerGas, maxPriorityFeePerGas, err := ett.GetGasLegacyEIP1559(ett.evmConfig.GasLegacy)
		if err != nil {
			return common.Hash{}, err
		}
		builder.SetGasFeeCap(maxFeePerGas).SetGasTipCap(maxPriorityFeePerGas)
	case gmgn_define.CHAIN_TYPE_BSC:
		gasPrice, err := ett.GetGasLegacy(ett.evmConfig.GasLegacy)
		if err != nil {
			return common.Hash{}, err
		}
		builder.SetGasPrice(gasPrice)
	}

	tx, err := builder.Build(ett.ctx)
	if err != nil {
		return common.Hash{}, errors.New("Approve: Failed to build approve transaction: " + err.Error())
	}

	signedTx, err := ett.client.Wallet.Sign(tx)
	if err != nil {
		return common.Hash{}, errors.New("Approve: Failed to sign approve transaction: " + err.Error())
	}

	err = ett.client.Wallet.BroadcastTransaction(ett.ctx, signedTx)
	if err != nil {
		return common.Hash{}, errors.New("Approve: Failed to broadcast approve transaction: " + err.Error())
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

func (ett *EvmTradeTool) CheckPermitSupport(tokenAddr string) error {
	// 代币 ABI，仅包含 nonces 和 name
	tokenABI, err := abi.JSON(strings.NewReader(`[
        {"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"nonces","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
        {"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"stateMutability":"view","type":"function"}
    ]`))
	if err != nil {
		return err
	}

	contract := bind.NewBoundContract(common.HexToAddress(tokenAddr), tokenABI, ett.ethClient, ett.ethClient, ett.ethClient)

	// 查询 nonce
	var nonceResults []interface{}
	err = contract.Call(nil, &nonceResults, "nonces", ett.client.Wallet.Address)
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

func (ett *EvmTradeTool) GetTokenData(tokenAddress string) (*tokens.ProviderTokenDtoFixed, error) {
	return ett.tokenClient.GetCustomToken(ett.ctx, tokens.CustomTokensControllerGetTokenInfoParams{
		Address: tokenAddress,
	})
}
