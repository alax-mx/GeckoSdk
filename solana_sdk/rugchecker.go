package solana_sdk

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type TopHolder struct {
	Owner string  `json:"owner"`
	Pct   float64 `json:"pct"`
}

type TokenData struct {
	Mint                 string      `json:"mint"`
	Token                TokenInfo   `json:"token"`
	TokenMeta            TokenMeta   `json:"tokenMeta"`
	TopHolders           []TopHolder `json:"topHolders"`
	Risks                []Risk      `json:"risks"`
	FileMeta             FileMeta    `json:"fileMeta"`
	Rugged               bool        `json:"rugged"`
	Markets              []Market    `json:"markets"`
	Price                float64     `json:"price"`
	TotalMarketLiquidity float64     `json:"totalMarketLiquidity"`
}

type TokenInfo struct {
	MintAuthority   string  `json:"mintAuthority"`
	Supply          float64 `json:"supply"` // Change type to float64
	Decimals        uint8   `json:"decimals"`
	FreezeAuthority string  `json:"freezeAuthority"`
}

type TokenMeta struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Mutable         bool   `json:"mutable"`
	UpdateAuthority string `json:"updateAuthority"`
}

type Risk struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Score       int    `json:"score"`
	Level       string `json:"level"`
}

type FileMeta struct {
	Image string `json:"image"`
}

type Market struct {
	MarketType        string            `json:"marketType"`
	MintA             string            `json:"mintA"`
	MintB             string            `json:"mintB"`
	LiquidityAAccount LiquidityAAccount `json:"liquidityAAccount"`
	LiquidityBAccount LiquidityBAccount `json:"liquidityBAccount"`
	Lp                LP                `json:"lp"`
}

type LP struct {
	BaseMint      string  `json:"baseMint"`  // 打哪个区代币
	QuoteMint     string  `json:"quoteMint"` // 交易代币
	QuotePrice    float64 `json:"quotePrice"`
	QuoteUSD      float64 `json:"quoteUSD"`    // 上市池子数量
	LpLocked      int     `json:"lpLocked"`    // 锁定的资金池子数量
	LpUnlocked    int     `json:"lpUnlocked"`  // 未锁定的资金池数量
	LpLockedPct   float64 `json:"lpLockedPct"` // 锁定的资金池百分比
	LpLockedUSD   float64 `json:"lpLockedUSD"` // 锁定的资金池
	LpTotalSupply int     `json:"lpTotalSupply"`
}

type LiquidityAAccount struct {
	Mint   string `json:"mint"`
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type LiquidityBAccount struct {
	Mint   string `json:"mint"`
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

func GetTokenRugInfo(token string) (*TokenData, error) {
	data, err := httptool.HttpGetRug(token)
	if err != nil {
		return nil, err
	}

	tokenData := &TokenData{}
	err = json.Unmarshal(data, tokenData)
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}
