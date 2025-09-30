package gmgn_trade

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/alax-mx/geckosdk/proxy"
)

type STTokenBalanceInfo struct {
	Amount   int64
	FAmount  float64
	Decimals uint8
}

type STTradeInfo struct {
	Hash       string
	RouterInfo STRouterInfo
}

type STSwapInfo struct {
	Ammkey       string      `json:"ammKey"`
	Label        string      `json:"label"`
	InputMint    string      `json:"inputMint"`
	OutputMint   string      `json:"outputMint"`
	InAmount     interface{} `json:"inAmount"`
	OutputAmount interface{} `json:"outAmount"`
	FeeAmount    interface{} `json:"feeAmount"`
	FeeMint      string      `json:"feeMint"`
}

type STRoutePlan struct {
	SwapInfo *STSwapInfo `json:"swapInfo"`
	Percent  int         `json:"percent"`
}

type STRawTx struct {
	SwapTransaction           string      `json:"swapTransaction"`
	LastValidBlockHeight      int         `json:"lastValidBlockHeight"`
	PrioritizationFeeLamports interface{} `json:"prioritizationFeeLamports"`
	RecentBlockhash           string      `json:"recentBlockhash"`
	Version                   string      `json:"version"`
}

type STQuoteInfo struct {
	InputMint            string         `json:"inputMint"`
	InAmount             any            `json:"inAmount"`
	OutputMint           string         `json:"outputMint"`
	OutputAmount         any            `json:"outAmount"`
	OtherAmountThreshold any            `json:"otherAmountThreshold"`
	InDecimals           int            `json:"inDecimals"`
	OutDecimals          int            `json:"outDecimals"`
	SwapMode             string         `json:"swapMode"`
	SlippageBps          any            `json:"slippageBps"`
	PlatformFee          any            `json:"platformFee"`
	PriceImpactPct       any            `json:"priceImpactPct"`
	RoutePlan            []*STRoutePlan `json:"routePlan"`
	TimeTaken            any            `json:"timeTaken"`
}

type STRouterInfo struct {
	Quote        STQuoteInfo `json:"quote"`
	RawTX        STRawTx     `json:"raw_tx"`
	AmountInUsd  interface{} `json:"amount_in_usd"`
	AmountOutUsd interface{} `json:"amount_out_usd"`
	JitoOrderID  interface{} `json:"jito_order_id"`
}

type GetRouterResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Tid  string       `json:"tid"`
	Data STRouterInfo `json:"data"`
}

type STTransactionData struct {
	Hash      string `json:"hash"`
	TimeTaken int    `json:"timeTaken"`
}

type STBundledTransactionData struct {
	OrderID              string `json:"order_id"`
	BundleID             string `json:"bundle_id"`
	LastValidBlockNumber int    `json:"last_valid_block_number"`
	TxHash               string `json:"tx_hash"`
}

type SendTransactionResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data STTransactionData `json:"data"`
}

type SendBundledTransactionResp struct {
	Code int                      `json:"code"`
	Msg  string                   `json:"msg"`
	Data STBundledTransactionData `json:"data"`
}

type STTradeStatusInfo struct {
	Success bool        `json:"success"`
	Failed  bool        `json:"failed"`
	Expired bool        `json:"expired"`
	Err     interface{} `json:"err"`
	ErrCode interface{} `json:"err_code"`
}

type GetTransactionStatusResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data STTradeStatusInfo `json:"data"`
}

type SolTradeTool struct {
	baseUrl   string
	pubKey    solana.PublicKey
	priKey    solana.PrivateKey
	rpcClinet *rpc.Client
	proxyInfo *proxy.STProxyInfo
}

func NewSolTradeTool(baseUrl string, pubKey string, priKey string) *SolTradeTool {
	if pubKey == "" || priKey == "" {
		fmt.Println("pubKey == nil || priKey == nil")
		return nil
	}
	return &SolTradeTool{
		baseUrl:   baseUrl,
		pubKey:    solana.MustPublicKeyFromBase58(pubKey),
		priKey:    solana.MustPrivateKeyFromBase58(priKey),
		rpcClinet: rpc.New(rpc.MainNetBeta_RPC),
		proxyInfo: nil,
	}
}
func (gtt *SolTradeTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	gtt.proxyInfo = proxyInfo
}

func (gtt *SolTradeTool) Swap(inAddress string, outAddress string, amount int, slippage float64, isAntiMev bool, fee string) (*STTradeInfo, error) {
	// GetRouter
	resp, err := gtt.GetSwapRouter(inAddress, outAddress, amount, gtt.pubKey.String(), slippage, fee)
	if err != nil {
		return nil, err
	}

	// Sign
	signStr, err := gtt.signTransaction(resp.Data.RawTX)
	if err != nil {
		return nil, err
	}

	transResp, err := gtt.sendTransaction(signStr, isAntiMev)
	if err != nil {
		return nil, err
	}
	if transResp.Code != 0 {
		return nil, errors.New("sendTransaction err: " + transResp.Msg)
	}

	return &STTradeInfo{
		Hash:       transResp.Data.Hash,
		RouterInfo: resp.Data,
	}, nil
}

func (gtt *SolTradeTool) SwapByRouter(routerResp *GetRouterResp, isAntiMev bool) (*STTradeInfo, error) {
	// Sign
	signStr, err := gtt.signTransaction(routerResp.Data.RawTX)
	if err != nil {
		return nil, err
	}

	transResp, err := gtt.sendTransaction(signStr, isAntiMev)
	if err != nil {
		return nil, err
	}
	if transResp.Code != 0 {
		return nil, errors.New("sendTransaction err: " + transResp.Msg)
	}

	return &STTradeInfo{
		Hash:       transResp.Data.Hash,
		RouterInfo: routerResp.Data,
	}, nil
}

func (gtt *SolTradeTool) GetSwapRouter(inAddress string, outAddress string, amount int,
	walletPubkey string, slippage float64, fee string) (*GetRouterResp, error) {
	tmpUrl := "/get_swap_route?token_in_address=" + inAddress
	tmpUrl += "&token_out_address=" + outAddress
	tmpUrl += "&in_amount=" + strconv.Itoa(amount)
	tmpUrl += "&from_address=" + walletPubkey
	tmpUrl += "&fee=" + fee
	tmpUrl += "&slippage=" + strconv.FormatFloat(slippage, 'f', 2, 64)
	data, err := HttpGet(gtt.baseUrl+tmpUrl, gtt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetRouterResp{}
	err2 := json.Unmarshal(data, ret)
	if err2 != nil {
		return nil, err2
	}

	return ret, nil
}

func (gtt *SolTradeTool) signTransaction(rawTx STRawTx) (string, error) {
	// Decode base64 transaction
	txBytes, err := base64.StdEncoding.DecodeString(rawTx.SwapTransaction)
	if err != nil {
		return "", err
	}

	// Deserialize the transaction
	tx, err := solana.TransactionFromBytes(txBytes)
	if err != nil {
		return "", err
	}

	// Sign the transaction
	tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(gtt.priKey.PublicKey()) {
			return &gtt.priKey
		}
		return nil
	})

	return tx.MustToBase64(), nil
}

func (gtt *SolTradeTool) sendTransaction(signedTx string, isAntiMev bool) (*SendTransactionResp, error) {
	tmpUrl := "https://gmgn.ai/txproxy/v1/send_transaction"
	param := make(map[string]interface{})
	param["chain"] = "sol"
	param["signedTx"] = signedTx
	param["isAntiMev"] = isAntiMev
	bytesData, _ := json.Marshal(param)
	data, err := HttpPostRouter(tmpUrl, bytesData, gtt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &SendTransactionResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (gtt *SolTradeTool) GetTransactionStatus(hash string, lastValidHeight int) (*GetTransactionStatusResp, error) {
	tmpUrl := "/get_transaction_status?"
	tmpUrl += "hash=" + hash
	tmpUrl += "&last_valid_height" + strconv.Itoa(lastValidHeight)
	data, err := HttpGet(gtt.baseUrl+tmpUrl, gtt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetTransactionStatusResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (gtt *SolTradeTool) GetSolBalance() (*rpc.GetBalanceResult, error) {
	out, err := gtt.rpcClinet.GetBalance(context.Background(), gtt.pubKey, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (gtt *SolTradeTool) GetTokenBalance(tokenAddress string) (*STTokenBalanceInfo, error) {
	tokenmint := solana.MustPublicKeyFromBase58(tokenAddress)                  //token 地址
	tokenacc, _, _ := solana.FindAssociatedTokenAddress(gtt.pubKey, tokenmint) //算出token账号地址
	outtbl, err := gtt.rpcClinet.GetTokenAccountBalance(context.Background(), tokenacc, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	ret := &STTokenBalanceInfo{}
	ret.Amount, _ = strconv.ParseInt(outtbl.Value.Amount, 10, 64)
	ret.Decimals = outtbl.Value.Decimals
	ret.FAmount = *outtbl.Value.UiAmount
	return ret, nil
}
