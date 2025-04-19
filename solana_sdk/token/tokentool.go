package token

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type TokenTool struct {
	solanaInfo              *basedef.STSolanaDefine
	tokenListTool           *TokenListTool
	tokenMetaTool           *TokenMetaTool
	tokenTopTool            *TokenTopTool
	tokenPriceTool          *TokenPriceTool
	tokenTransferTool       *TokenTransferTool
	tokenDefiactivitiesTool *TokenDefiactivitiesTool
	tokenMarketsTool        *TokenMarketsTool
	tokenTrendingTool       *TokenTrendingTool
	tokenHoldersTool        *TokenHoldersTool
}

func NewTokenTool(solanaInfo *basedef.STSolanaDefine) *TokenTool {
	return &TokenTool{
		solanaInfo:              solanaInfo,
		tokenListTool:           NewTokenListTool(solanaInfo),
		tokenMetaTool:           NewTokenMetaTool(solanaInfo),
		tokenTopTool:            NewTokenTopTool(solanaInfo),
		tokenPriceTool:          NewTokenPriceTool(solanaInfo),
		tokenTransferTool:       NewTokenTransferTool(solanaInfo),
		tokenDefiactivitiesTool: NewTokenDefiactivitiesTool(solanaInfo),
		tokenMarketsTool:        NewTokenMarketsTool(solanaInfo),
		tokenTrendingTool:       NewTokenTrendingTool(solanaInfo),
		tokenHoldersTool:        NewTokenHoldersTool(solanaInfo),
	}
}

func (tt *TokenTool) GetTokenList(sortBy string, sortOrder string, page int, pageSize int) (*STTokenListResp, error) {
	return tt.tokenListTool.GetTokenList(sortBy, sortOrder, page, pageSize)
}

func (tt *TokenTool) GetTokenPrice(address string, fromTime int, toTime int) (*STTokenPriceResp, error) {
	return tt.tokenPriceTool.GetTokenPrice(address, fromTime, toTime)
}

func (tt *TokenTool) GetTokenHolders(address string, page int, pageSize int, fromAmount string, toAmount string) (*STTokenHoldersResp, error) {
	return tt.tokenHoldersTool.GetTokenHolders(address, page, pageSize, fromAmount, toAmount)
}

func (tt *TokenTool) GetTokenMeta(address string) (*STTokenMetaResp, error) {
	return tt.tokenMetaTool.GetTokenMeta(address)
}

func (tt *TokenTool) GetTokenTop() (*STTokenTopResp, error) {
	return tt.tokenTopTool.GetTokenTop()
}

func (tt *TokenTool) GetTokenTrending(limit int) (*STTokenTrendingResp, error) {
	return tt.tokenTrendingTool.GetTokenTrending(limit)
}
