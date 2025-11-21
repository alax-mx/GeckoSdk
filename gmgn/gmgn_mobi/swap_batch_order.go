package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STSwapOrderData struct {
	Chain                string  `json:"chain"`
	FromAddress          string  `json:"from_address"`
	InputToken           string  `json:"input_token"`
	OutputToken          string  `json:"output_token"`
	InputAmount          string  `json:"input_amount"`
	Slippage             int     `json:"slippage"`
	AutoSlippage         bool    `json:"auto_slippage"`
	RetryOnSubmitFailed  int     `json:"retry_on_submit_failed"`
	SimulateBeforeSubmit bool    `json:"simulate_before_submit"`
	IsAntiMev            bool    `json:"is_anti_mev"`
	AutoApproveAfterBuy  bool    `json:"auto_approve_after_buy"`
	Source               string  `json:"source"`
	CreateStrategy       bool    `json:"create_strategy"`
	Decimals             int     `json:"decimals"`
	Standard             string  `json:"standard"`
	PriorityGasPrice     string  `json:"priority_gas_price"`
	GasPrice             string  `json:"gas_price"`
	PriorityFee          string  `json:"priority_fee"`
	TipFee               string  `json:"tip_fee"`
	Fee                  float64 `json:"fee"`
}

type STConfirmation struct {
	State  string `json:"state"`
	Detail any    `json:"detail"`
}

type STSwapRespData struct {
	State        int            `json:"state"`
	Hash         string         `json:"hash"`
	OrderID      string         `json:"order_id"`
	ErrorCode    string         `json:"error_code"`
	ErrorStatus  string         `json:"error_status"`
	Confirmation STConfirmation `json:"confirmation"`
}

type SwapOrderResp struct {
	Code    int            `json:"code"`
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Data    STSwapRespData `json:"data"`
}

type SwapBatchOrderTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewSwapBatchOrderTool(baseUrl string, baseParam string, authStr string) *SwapBatchOrderTool {
	return &SwapBatchOrderTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (tdt *SwapBatchOrderTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	tdt.proxyInfo = proxyInfo
}

func (tdt *SwapBatchOrderTool) Swap(orderData *STSwapOrderData) (*SwapOrderResp, error) {
	postData, err := json.Marshal(orderData)
	if err != nil {
		return nil, err
	}

	url := "mrtapi/v2/swap_batch_order" + "?" + tdt.baseParam
	data, err := HttpPost(tdt.baseUrl+url, postData, tdt.authStr, tdt.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &SwapOrderResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
