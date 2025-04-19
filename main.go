package main

import (
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/geck_sdk"
)

func main() {
	GetTokenOhlcvInfo("3bHCTtfKwJv557P8EKv3nMunrbUhrgepJxu39m3ftMxU", "8pGeXS65kYBvS37jf3imYWkdPwM8seeTxrxMSSpbpump")
}

func GetTokenOhlcvInfo(poolAddress string, tokenAddress string) {
	resp, err := geck_sdk.NewNetworkOhlcvTool().GetNetworkOhlcv("solana", poolAddress, geck_sdk.OHLCV_TIME_FRAME_TYPE_MINUTE, geck_sdk.OHLCV_AGREGATE_HOUR_1, tokenAddress)
	if err != nil {
		fmt.Println(err)
	}

	baseutils.ShowObjectValue(resp)
}
