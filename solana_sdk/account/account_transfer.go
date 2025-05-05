package account

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type STAccountTransferData struct {
	BlockID       int    `json:"block_id"`
	TransID       string `json:"trans_id"`
	BlockTime     int    `json:"block_time"`
	Time          string `json:"time"`
	ActivityType  string `json:"activity_type"`
	FromAddress   string `json:"from_address"`
	ToAddress     string `json:"to_address"`
	TokenAddress  string `json:"token_address"`
	TokenDecimals int    `json:"token_decimals"`
	Amount        int    `json:"amount"`
	Flow          string `json:"flow"`
}

type STAccountTransferResp struct {
	Success  bool                     `json:"success"`
	DataList []*STAccountTransferData `json:"data"`
	Errors   basedef.STErrors         `json:"errors"`
}
