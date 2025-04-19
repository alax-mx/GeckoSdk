package main

import (
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/geck_sdk"
)

func main() {
	// resp, err := geck_sdk.NewNetworkPoolsTool().GetNetworkPools("solana", "GSg2nNSm4zSMxjVQ7x3KxRWJ7JpTtRGyaSMsX4PwtH5N", "")
	resp, err := geck_sdk.NewNetworkPoolsTool().GetNetworkNewPools("solana", "", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	baseutils.ShowObjectValue(resp)
}
