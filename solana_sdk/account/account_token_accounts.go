package account

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

const (
	TYPE_TOKEN string = "token"
	TYPE_NFT   string = "nft"
)

type STAccountTokenAccountsData struct {
	TokenAccount  string `json:"token_account"`
	TokenAddress  string `json:"token_address"`
	Amount        int    `json:"amount"`
	TokenDecimals int    `json:"token_decimals"`
	Owner         string `json:"owner"`
}

type STAccountTokenAccountsResp struct {
	Success  bool                          `json:"success"`
	DataList []*STAccountTokenAccountsData `json:"data"`
	Errors   basedef.STErrors              `json:"errors"`
}

type AccountTokenAccountsTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewAccountTokenAccountsTool(solanaInfo *basedef.STSolanaDefine) *AccountTokenAccountsTool {
	return &AccountTokenAccountsTool{
		solanaInfo: solanaInfo,
	}
}

func (atat *AccountTokenAccountsTool) GetAccountTokenAccounts(address string, stype string, page int, pageSize int, hideZero bool) (*STAccountTokenAccountsResp, error) {
	newUrl := basedef.G_ACCOUNT_TOKEN_ACCOUNTS_URL + "address=" + address + "&type=" + stype
	if page >= 0 {
		newUrl += "&page=" + strconv.Itoa(page)
	}
	if pageSize >= 0 {
		newUrl += "&page_size=" + strconv.Itoa(pageSize)
	}
	if hideZero {
		newUrl += "&hide_zero=true"
	} else {
		newUrl += "&hide_zero=false"
	}

	data, err := httptool.HttpGet(newUrl, atat.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STAccountTokenAccountsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
