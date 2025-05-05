package token

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type STTokenMarketsData struct {
	PoolID             string  `json:"pool_id"`
	ProgramID          string  `json:"program_id"`
	Token1             string  `json:"token_1"`
	Token2             string  `json:"token_2"`
	TokenAccount1      string  `json:"token_account_1"`
	TokenAccount2      string  `json:"token_account_2"`
	TokenTrades24H     int     `json:"total_trades_24h"`
	TotalTradesPrev24H int     `json:"total_trades_prev_24h"`
	TotalVolume24H     float64 `json:"total_volume_24h"`
	TotalVolumePrev24H float64 `json:"total_volume_prev_24h"`
}

type STTokenMarketsResp struct {
	Success  bool                  `json:"success"`
	DataList []*STTokenMarketsData `json:"data"`
	Errors   basedef.STErrors      `json:"errors"`
}

type TokenMarketsTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenMarketsTool(solanaInfo *basedef.STSolanaDefine) *TokenMarketsTool {
	return &TokenMarketsTool{
		solanaInfo: solanaInfo,
	}
}
