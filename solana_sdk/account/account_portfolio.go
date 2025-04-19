package account

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STNativeBalance struct {
	Amount        int     `json:"amount"`
	Balance       float64 `json:"balance"`
	TokenPrice    float64 `json:"token_price"`
	TokenDecimals int     `json:"token_decimals"`
	TokenName     string  `json:"token_name"`
	TokenSymbol   string  `json:"token_symbol"`
	TokenIcon     string  `json:"token_icon"`
	Value         float64 `json:"value"`
}

type STTokens struct {
	TokenAddress  string  `json:"token_address"`
	Amount        int     `json:"amount"`
	Balance       float64 `json:"balance"`
	TokenPrice    float64 `json:"token_price"`
	TokenDecimals int     `json:"token_decimals"`
	TokenName     string  `json:"token_name"`
	TokenSymbol   string  `json:"token_symbol"`
	TokenIcon     string  `json:"token_icon"`
	Value         float64 `json:"value"`
}

type STAccountPortfolioData struct {
	TotalValue    float64          `json:"total_value"`
	NativeBalance *STNativeBalance `json:"native_balance"`
	Tokens        []*STTokens      `json:"tokens"`
}

type STAccountPortfolioResp struct {
	Success  bool                    `json:"success"`
	DataList *STAccountPortfolioData `json:"data"`
	Errors   basedef.STErrors        `json:"errors"`
}

type AccountPortfolioTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewAccountPortfolioTool(solanaInfo *basedef.STSolanaDefine) *AccountPortfolioTool {
	return &AccountPortfolioTool{
		solanaInfo: solanaInfo,
	}
}

func (apt *AccountPortfolioTool) GetAccountPortfolio(address string) (*STAccountPortfolioResp, error) {
	newUrl := basedef.G_ACCOUNT_PORTFOLIO_URL + "address=" + address
	data, err := httptool.HttpGet(newUrl, apt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STAccountPortfolioResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
