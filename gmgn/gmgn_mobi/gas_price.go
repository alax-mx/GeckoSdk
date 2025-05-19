package gmgn_mobi

import "encoding/json"

type STGasPriceData struct {
	Auto                string  `json:"auto"`
	AutoMev             string  `json:"auto_mev"`
	LastBlock           int     `json:"last_block"`
	High                string  `json:"high"`
	Average             string  `json:"average"`
	Low                 string  `json:"low"`
	SuggestBaseFee      string  `json:"suggest_base_fee"`
	HighPrioFee         string  `json:"high_prio_fee"`
	AveragePrioFee      string  `json:"average_prio_fee"`
	LowPrioFee          string  `json:"low_prio_fee"`
	HighPrioFeeMixed    string  `json:"high_prio_fee_mixed"`
	AveragePrioFeeMixed string  `json:"average_prio_fee_mixed"`
	LowPrioFeeMixed     string  `json:"low_prio_fee_mixed"`
	NativeTokenUsdPrice float64 `json:"native_token_usd_price"`
	EthUsdPrice         float64 `json:"eth_usd_price"`
	HighEstimateTime    int     `json:"high_estimate_time"`
	AverageEstimateTime int     `json:"average_estimate_time"`
	LowEstimateTime     int     `json:"low_estimate_time"`
	HighOrign           string  `json:"high_orign"`
	AverageOrign        string  `json:"average_orign"`
	LowOrign            string  `json:"low_orign"`
}

type GetGasPriceResp struct {
	Code    int            `json:"code"`
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Data    STGasPriceData `json:"data"`
}

type GasPriceTool struct {
	baseUrl   string
	baseParam string
}

func NewGasPriceTool(baseUrl string, baseParam string) *GasPriceTool {
	return &GasPriceTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (gpt *GasPriceTool) Get() (*GetGasPriceResp, error) {
	url := "api/v1/gas_price/sol?" + gpt.baseParam
	data, err := HttpGet(gpt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetGasPriceResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
