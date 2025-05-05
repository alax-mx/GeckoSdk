package account

import "github.com/alax-mx/geckosdk/solana_sdk/basedef"

type STParsedInstructions struct {
	Type      string `json:"type"`
	Program   string `json:"program"`
	ProgramID string `json:"program_id"`
}

type STAccountTransactionsData struct {
	Slot               int                     `json:"slot"`
	Fee                int                     `json:"fee"`
	Status             string                  `json:"status"`
	Signer             []string                `json:"signer"`
	BlockTime          int                     `json:"block_time"`
	TXHash             string                  `json:"tx_hash"`
	ParsedInstructions []*STParsedInstructions `json:"parsed_instructions"`
	ProgramIDS         []string                `json:"program_ids"`
	Time               string                  `json:"time"`
}

type STAccountTransactionsResp struct {
	Success  bool                         `json:"success"`
	DataList []*STAccountTransactionsData `json:"data"`
	Errors   basedef.STErrors             `json:"errors"`
}
