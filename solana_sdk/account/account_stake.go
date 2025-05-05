package account

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STAccountStakeData struct {
	Amount               int      `json:"amount"`
	Role                 []string `json:"role"`
	Status               string   `json:"status"`
	Type                 string   `json:"type"`
	Voter                string   `json:"voter"`
	ActiveStakeAmount    int      `json:"active_stake_amount"`
	DelegatedStakeAmount int      `json:"delegated_stake_amount"`
	SolBalance           int      `json:"sol_balance"`
	TotalReward          string   `json:"total_reward"`
	StakeAccount         string   `json:"stake_account"`
	ActivationEpoch      int      `json:"activation_epoch"`
	StakeType            int      `json:"stake_type"`
}

type STAccountStakeResp struct {
	Success  bool                  `json:"success"`
	DataList []*STAccountStakeData `json:"data"`
	Errors   basedef.STErrors      `json:"errors"`
}

type AccountStakeTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewAccountStakeTool(solanaInfo *basedef.STSolanaDefine) *AccountStakeTool {
	return &AccountStakeTool{
		solanaInfo: solanaInfo,
	}
}

func (ast *AccountStakeTool) GetAccountStake(address string, page int, pageSize int) (*STAccountStakeResp, error) {
	newUrl := basedef.G_ACCOUNT_STAKE_URL + "address=" + address
	if page >= 0 {
		newUrl += "&page=" + strconv.Itoa(page)
	}
	if pageSize >= 0 {
		newUrl += "&page_size=" + strconv.Itoa(pageSize)
	}

	data, err := httptool.HttpGet(newUrl, ast.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STAccountStakeResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
