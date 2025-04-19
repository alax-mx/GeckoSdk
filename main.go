package main

import (
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/geck_sdk"
	"github.com/alax-mx/geckosdk/solana_sdk"
)

func main() {
	// resp, err := geck_sdk.NewNetworkPoolsTool().GetNetworkPools("solana", "28LS9W8i6BAFUFSRJMFk2ttJsoXpoKNpkzdD7NYZtuT2", "")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// baseutils.ShowObjectValue(resp)

	// resp2, err2 := geck_sdk.NewNetworkTokensTool().GetNetworkPoolTokensInfo("solana", resp.Data.Attributes.Address)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return
	// }
	// baseutils.ShowObjectValue(resp2)
}

func CheckPoolLocked(poolAddress string) (float64, error) {
	resp, err := geck_sdk.NewNetworkTokensTool().GetNetworkPoolTokensInfo("solana", poolAddress)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	baseutils.ShowObjectValue(resp)

	data, err2 := solana_sdk.GetTokenRugInfo(resp.Data[1].Attributes.Address)
	if err2 != nil {
		fmt.Println(err2)
		return 0, err2
	}
	baseutils.ShowObjectValue(data)
	// resp, err := geck_sdk.NewNetworkPoolsTool().GetNetworkPools("solana", "28LS9W8i6BAFUFSRJMFk2ttJsoXpoKNpkzdD7NYZtuT2", "")
	// if err != nil {
	// 	return 0, err
	// }

	// data, err2 := solana_sdk.GetTokenRugInfo(poolAddress)

	// baseutils.ShowObjectValue(resp)
	// percent, err2 := strconv.ParseFloat(resp.Data.Attributes.LockedLiquidityPercentage, 32)
	// if err2 != nil {
	// 	return 0, err2
	// }

	return 0, nil
}
