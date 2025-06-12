package gmgn_mobi

import (
	"encoding/json"
	"strconv"
)

type STTagRank struct {
	FreshWallet int `json:"fresh_wallet"`
}
type STDailyProfit7D struct {
	Timestamp int     `json:"timestamp"`
	Profit    float64 `json:"profit"`
}
type STRisk struct {
	TokenActive        string  `json:"token_active"`
	TokenHoneypot      string  `json:"token_honeypot"`
	TokenHoneypotRatio float64 `json:"token_honeypot_ratio"`
	NoBuyHold          string  `json:"no_buy_hold"`
	NoBuyHoldRatio     float64 `json:"no_buy_hold_ratio"`
	SellPassBuy        string  `json:"sell_pass_buy"`
	SellPassBuyRatio   float64 `json:"sell_pass_buy_ratio"`
	FastTx             string  `json:"fast_tx"`
	FastTxRatio        float64 `json:"fast_tx_ratio"`
}
type STWalletRank struct {
	WalletAddress            string            `json:"wallet_address"`
	Address                  string            `json:"address"`
	RealizedProfit           float64           `json:"realized_profit"`
	Buy                      int               `json:"buy"`
	Sell                     int               `json:"sell"`
	LastActive               int               `json:"last_active"`
	RealizedProfit1D         float64           `json:"realized_profit_1d"`
	RealizedProfit7D         float64           `json:"realized_profit_7d"`
	RealizedProfit30D        float64           `json:"realized_profit_30d"`
	Pnl30D                   float64           `json:"pnl_30d"`
	Pnl7D                    float64           `json:"pnl_7d"`
	Pnl1D                    any               `json:"pnl_1d"`
	Txs30D                   int               `json:"txs_30d"`
	Buy30D                   int               `json:"buy_30d"`
	Sell30D                  int               `json:"sell_30d"`
	Balance                  float64           `json:"balance"`
	EthBalance               float64           `json:"eth_balance"`
	SolBalance               float64           `json:"sol_balance"`
	TrxBalance               float64           `json:"trx_balance"`
	FollowCount              int               `json:"follow_count"`
	RemarkCount              int               `json:"remark_count"`
	TwitterUsername          any               `json:"twitter_username"`
	Avatar                   any               `json:"avatar"`
	Ens                      any               `json:"ens"`
	Tag                      any               `json:"tag"`
	TagRank                  STTagRank         `json:"tag_rank"`
	Nickname                 any               `json:"nickname"`
	Tags                     []string          `json:"tags"`
	TwitterName              any               `json:"twitter_name"`
	FollowersCount           int               `json:"followers_count"`
	IsBlueVerified           any               `json:"is_blue_verified"`
	TwitterDescription       any               `json:"twitter_description"`
	Name                     any               `json:"name"`
	AvgHoldTime              int               `json:"avg_hold_time"`
	RecentBuyTokens          []any             `json:"recent_buy_tokens"`
	Winrate7D                float64           `json:"winrate_7d"`
	AvgCost7D                float64           `json:"avg_cost_7d"`
	PnlLtMinusDot5Num7D      int               `json:"pnl_lt_minus_dot5_num_7d"`
	PnlMinusDot50XNum7D      int               `json:"pnl_minus_dot5_0x_num_7d"`
	PnlLt2XNum7D             int               `json:"pnl_lt_2x_num_7d"`
	Pnl2X5XNum7D             int               `json:"pnl_2x_5x_num_7d"`
	PnlGt5XNum7D             int               `json:"pnl_gt_5x_num_7d"`
	PnlLtMinusDot5Num7DRatio float64           `json:"pnl_lt_minus_dot5_num_7d_ratio"`
	PnlMinusDot50XNum7DRatio float64           `json:"pnl_minus_dot5_0x_num_7d_ratio"`
	PnlLt2XNum7DRatio        float64           `json:"pnl_lt_2x_num_7d_ratio"`
	Pnl2X5XNum7DRatio        float64           `json:"pnl_2x_5x_num_7d_ratio"`
	PnlGt5XNum7DRatio        float64           `json:"pnl_gt_5x_num_7d_ratio"`
	DailyProfit7D            []STDailyProfit7D `json:"daily_profit_7d"`
	Risk                     STRisk            `json:"risk"`
	Txs                      int               `json:"txs"`
	TokenNum7D               int               `json:"token_num_7d"`
	AvgHoldingPeriod7D       float64           `json:"avg_holding_period_7d"`
}
type STWalletMonitorData struct {
	Rank []STWalletRank `json:"rank"`
}

type GetTokenWalletMonitorResp struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data STWalletMonitorData `json:"data"`
}

type TokenWalletMonitorTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenWalletMonitorTool(baseUrl string, baseParam string) *TokenWalletMonitorTool {
	return &TokenWalletMonitorTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tst *TokenWalletMonitorTool) Get(orderBy string, limit int) (*GetTokenWalletMonitorResp, error) {
	url := "defi/quotation/v1/rank/sol/wallets/7d?" + tst.baseParam
	url += "&orderby=" + orderBy
	url += "&direction=desc"
	url += "&limit=" + strconv.Itoa(limit)
	data, err := HttpGet(tst.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenWalletMonitorResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
