package main

import (
	"fmt"

	"flyu.gecksdk/baseutils"
	"flyu.gecksdk/geck_sdk"
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
