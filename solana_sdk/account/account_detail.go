package account

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/solana_sdk/basedef"
	"github.com/alax-mx/geckosdk/solana_sdk/httptool"
)

type STAccountDetailData struct {
	Account      string `json:"account"`
	Lamports     int    `json:"lamports"`
	Type         string `json:"type"`
	Executable   bool   `json:"executable"`
	OwnerProgram string `json:"owner_program"`
	// RentEpoch    int    `json:"rent_epoch"`
	ISOncurve bool `json:"is_oncurve"`
}

type STAccountDetailResp struct {
	Success  bool                 `json:"success"`
	DataList *STAccountDetailData `json:"data"`
	Errors   basedef.STErrors     `json:"errors"`
}

type AccountDetailTool struct {
	solanaInfo *basedef.STSolanaDefine
}

func NewAccountDetailTool(solanaInfo *basedef.STSolanaDefine) *AccountDetailTool {
	return &AccountDetailTool{
		solanaInfo: solanaInfo,
	}
}

func (adt *AccountDetailTool) GetAccountDetail(address string) (*STAccountDetailResp, error) {
	newUrl := basedef.G_ACCOUNT_DETAIL_URL + "address=" + address
	data, err := httptool.HttpGet(newUrl, adt.solanaInfo.APIKey)
	if err != nil {
		return nil, err
	}

	ret := &STAccountDetailResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
