package main

import (
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/geck_sdk"
)

func main() {
	resp, err := geck_sdk.NewNetworkPoolsTool().GetNetworkPools("solana", "28LS9W8i6BAFUFSRJMFk2ttJsoXpoKNpkzdD7NYZtuT2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	baseutils.ShowObjectValue(resp)

	resp2, err2 := geck_sdk.NewNetworkTokensTool().GetNetworkPoolTokensInfo("solana", resp.Data.Attributes.Address)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	baseutils.ShowObjectValue(resp2)
}
