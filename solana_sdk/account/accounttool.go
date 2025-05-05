package account

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type AccountTool struct {
	solanaInfo               *basedef.STSolanaDefine
	accountDetailTool        *AccountDetailTool
	accountStakeTool         *AccountStakeTool
	accountTokenAccountsTool *AccountTokenAccountsTool
	accountPortfolioTool     *AccountPortfolioTool
}

func NewAccountTool(solanaInfo *basedef.STSolanaDefine) *AccountTool {
	return &AccountTool{
		solanaInfo:               solanaInfo,
		accountDetailTool:        NewAccountDetailTool(solanaInfo),
		accountStakeTool:         NewAccountStakeTool(solanaInfo),
		accountTokenAccountsTool: NewAccountTokenAccountsTool(solanaInfo),
		accountPortfolioTool:     NewAccountPortfolioTool(solanaInfo),
	}
}

func (at *AccountTool) GetAccountDetail(address string) (*STAccountDetailResp, error) {
	return at.accountDetailTool.GetAccountDetail(address)
}

func (at *AccountTool) GetAccountStake(address string, page int, pageSize int) (*STAccountStakeResp, error) {
	return at.accountStakeTool.GetAccountStake(address, page, pageSize)
}

func (at *AccountTool) GetAccountTokenAccounts(address string, stype string, page int, pageSize int, hideZero bool) (*STAccountTokenAccountsResp, error) {
	return at.accountTokenAccountsTool.GetAccountTokenAccounts(address, stype, page, pageSize, hideZero)
}

func (at *AccountTool) GetAccountPortfolio(address string) (*STAccountPortfolioResp, error) {
	return at.accountPortfolioTool.GetAccountPortfolio(address)
}
