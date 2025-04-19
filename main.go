package main

import (
	"fmt"
	"time"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/geck_sdk"
)

func main() {
	GetTokenOhlcvInfo("3bHCTtfKwJv557P8EKv3nMunrbUhrgepJxu39m3ftMxU", "8pGeXS65kYBvS37jf3imYWkdPwM8seeTxrxMSSpbpump")
}

func GetTokenOhlcvInfo(poolAddress string, tokenAddress string) {
	ohlcvTool := geck_sdk.NewNetworkOhlcvTool()
	resp, err := ohlcvTool.GetNetworkOhlcv("solana", poolAddress, geck_sdk.OHLCV_TIME_FRAME_TYPE_HOUR, geck_sdk.OHLCV_AGREGATE_HOUR_1, tokenAddress)
	if err != nil {
		fmt.Println(err)
	}

	ohlcvInfoList := ohlcvTool.ParseOHLCVData(resp.Data.Attributes.OhlcvList)
	for i := 0; i < len(ohlcvInfoList); i++ {
		tm := time.Unix(ohlcvInfoList[i].Time, 0)
		fmt.Println("time = ", tm.Format("2006-01-02 15:04:05"))
		baseutils.ShowObjectValue(ohlcvInfoList[i])
	}
}
