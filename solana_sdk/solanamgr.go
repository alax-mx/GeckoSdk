package solana_sdk

import (
	"fmt"

	"github.com/alax-mx/geckosdk/solana_sdk/account"
	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/token"
)

type SolanaMgr struct {
	solanaInfo  *basedef.STSolanaDefine
	accountTool *account.AccountTool
	tokenTool   *token.TokenTool
}

func NewSolanaMgr(cfgPath string) *SolanaMgr {
	solanaInfo := basedef.NewSTSolanaDefine(cfgPath)
	if solanaInfo == nil {
		fmt.Println("NewSolanaMgr failed!! solanainfo == nil")
		return nil
	}

	return &SolanaMgr{
		solanaInfo:  solanaInfo,
		accountTool: account.NewAccountTool(solanaInfo),
		tokenTool:   token.NewTokenTool(solanaInfo),
	}
}

func (sm *SolanaMgr) GetAccountTool() *account.AccountTool {
	return sm.accountTool
}

func (sm *SolanaMgr) GetTokenTool() *token.TokenTool {
	return sm.tokenTool
}
