package basedef

import (
	"encoding/json"
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
)

type STSolanaDefine struct {
	APIKey       string `json:"api_key"`
	WalletPriKey string `json:"wallet_pri_key"`
}

type STErrors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSTSolanaDefine(cfgPath string) *STSolanaDefine {
	data, err := baseutils.ReadFile(cfgPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ret := &STSolanaDefine{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return ret
}
