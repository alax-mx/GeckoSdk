package evm_trade

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	publicKey  common.Address
}

func NewWallet(rpcURL, privateKeyStr string) (*Wallet, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &Wallet{
		client:     client,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

func (w *Wallet) GetAuth(chainID *big.Int) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(w.privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (w *Wallet) GetBalance(tokenAddress common.Address) (*big.Float, error) {
	ctx := context.Background()

	// ETH余额查询
	if tokenAddress == (common.Address{}) {
		balance, err := w.client.BalanceAt(ctx, w.publicKey, nil)
		if err != nil {
			return nil, err
		}
		ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
		return ethValue, nil
	}

	// ERC20代币余额查询
	// 这里需要实现ERC20合约调用
	return big.NewFloat(0), nil
}

func (w *Wallet) GetAddress() common.Address {
	return w.publicKey
}
