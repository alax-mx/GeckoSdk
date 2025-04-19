package token

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type STChildRouter struct {
	Token1         string `json:"token1"`
	Token1Decimals int    `json:"token1_decimals"`
	Amount1        int    `json:"amount1"`
	Token2         string `json:"token2"`
	Token2Decimals int    `json:"token2_decimals"`
	Amount2        int    `json:"amount2"`
}

type STRouters struct {
	Token1         string          `json:"token1"`
	Token1Decimals int             `json:"token1_decimals"`
	Amount1        int             `json:"amount1"`
	Token2         string          `json:"token2"`
	Token2Decimals int             `json:"token2_decimals"`
	Amount2        int             `json:"amount2"`
	ChildRouters   []STChildRouter `json:"child_routers"`
}

type STTokenDefiActivitiesData struct {
	BlockID      int       `json:"block_id"`
	TransID      string    `json:"trans_id"`
	BlockTime    int       `json:"block_time"`
	ActivityType string    `json:"activity_type"`
	FromAddress  string    `json:"from_address"`
	ToAddress    string    `json:"to_address"`
	Sources      []string  `json:"sources"`
	Platform     string    `json:"platform"`
	Routers      STRouters `json:"routers"`
}

type STTokenDefiActivitiesResp struct {
	Success  bool                         `json:"success"`
	DataList []*STTokenDefiActivitiesData `json:"data"`
	Errors   basedef.STErrors             `json:"errors"`
}

type TokenDefiactivitiesTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewTokenDefiactivitiesTool(solanaInfo *basedef.STSolanaDefine) *TokenDefiactivitiesTool {
	return &TokenDefiactivitiesTool{
		solanaInfo: solanaInfo,
	}
}
