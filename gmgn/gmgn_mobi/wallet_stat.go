package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STWalletStatData struct {
	Buy                 int       `json:"buy"`
	Buy1D               int       `json:"buy_1d"`
	Buy7D               int       `json:"buy_7d"`
	Buy30D              int       `json:"buy_30d"`
	Sell                int       `json:"sell"`
	Sell1D              int       `json:"sell_1d"`
	Sell7D              int       `json:"sell_7d"`
	Sell30D             int       `json:"sell_30d"`
	Pnl                 float64   `json:"pnl"`
	Pnl1D               float64   `json:"pnl_1d"`
	Pnl7D               float64   `json:"pnl_7d"`
	Pnl30D              float64   `json:"pnl_30d"`
	AllPnl              float64   `json:"all_pnl"`
	RealizedProfit      float64   `json:"realized_profit"`
	RealizedProfit1D    float64   `json:"realized_profit_1d"`
	RealizedProfit7D    float64   `json:"realized_profit_7d"`
	RealizedProfit30D   float64   `json:"realized_profit_30d"`
	UnrealizedProfit    any       `json:"unrealized_profit"`
	UnrealizedPnl       any       `json:"unrealized_pnl"`
	TotalProfit         float64   `json:"total_profit"`
	TotalProfitPnl      float64   `json:"total_profit_pnl"`
	Balance             string    `json:"balance"`
	EthBalance          string    `json:"eth_balance"`
	SolBalance          string    `json:"sol_balance"`
	TrxBalance          string    `json:"trx_balance"`
	BnbBalance          string    `json:"bnb_balance"`
	TotalValue          float64   `json:"total_value"`
	Winrate             float64   `json:"winrate"`
	TokenSoldAvgProfit  float64   `json:"token_sold_avg_profit"`
	HistoryBoughtCost   float64   `json:"history_bought_cost"`
	TokenAvgCost        float64   `json:"token_avg_cost"`
	TokenNum            int       `json:"token_num"`
	ProfitNum           int       `json:"profit_num"`
	PnlLtMinusDot5Num   int       `json:"pnl_lt_minus_dot5_num"`
	PnlMinusDot50XNum   int       `json:"pnl_minus_dot5_0x_num"`
	PnlLt2XNum          int       `json:"pnl_lt_2x_num"`
	Pnl2X5XNum          int       `json:"pnl_2x_5x_num"`
	PnlGt5XNum          int       `json:"pnl_gt_5x_num"`
	GasCost             float64   `json:"gas_cost"`
	Bind                bool      `json:"bind"`
	Avatar              string    `json:"avatar"`
	Name                string    `json:"name"`
	Ens                 string    `json:"ens"`
	Tags                []string  `json:"tags"`
	TagRank             STTagRank `json:"tag_rank"`
	TwitterName         string    `json:"twitter_name"`
	TwitterUsername     string    `json:"twitter_username"`
	TwitterBind         bool      `json:"twitter_bind"`
	TwitterFansNum      int       `json:"twitter_fans_num"`
	FollowersCount      int       `json:"followers_count"`
	IsContract          bool      `json:"is_contract"`
	LastActiveTimestamp int       `json:"last_active_timestamp"`
	Risk                STRisk    `json:"risk"`
	AvgHoldingPeroid    float64   `json:"avg_holding_peroid"`
	UpdatedAt           int       `json:"updated_at"`
	RefreshRequestedAt  any       `json:"refresh_requested_at"`
	FollowCount         int       `json:"follow_count"`
	RemarkCount         int       `json:"remark_count"`
	TotalVolume         float64   `json:"total_volume"`
}

type GetWalletStatResp struct {
	Code    int              `json:"code"`
	Reason  string           `json:"reason"`
	Message string           `json:"message"`
	Data    STWalletStatData `json:"data"`
}

type WalletStatTool struct {
	baseUrl   string
	baseParam string
	authStr   string
	proxyInfo *proxy.STProxyInfo
}

func NewWalletStatTool(baseUrl string, baseParam string, authStr string) *WalletStatTool {
	return &WalletStatTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		authStr:   authStr,
		proxyInfo: nil,
	}
}

func (wht *WalletStatTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	wht.proxyInfo = proxyInfo
}

func (tpt *WalletStatTool) SetAuthString(authStr string) {
	tpt.authStr = authStr
}

func (wht *WalletStatTool) Get(chainType string, walletAddress string, period string) (*GetWalletStatResp, error) {
	url := "api/v1/wallet_stat/" + chainType + "/" + walletAddress + "/" + period + "?" + wht.baseParam
	url += "&period=" + period
	data, err := HttpGet(wht.baseUrl+url, wht.authStr, wht.proxyInfo)
	if err != nil {
		return nil, err
	}

	ret := &GetWalletStatResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
