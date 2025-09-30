package main

import (
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/eth_trade"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	data, _ := baseutils.ReadFile("device.json")
	deviceInfo := &gmgn_mobi.DeviceInfo{}
	json.Unmarshal(data, deviceInfo)
	gmgnTool := gmgn.NewGmgnTool("",
		"",
		deviceInfo, gmgn_define.CHAIN_TYPE_ETH)

	solDecimals := solana.DecimalsInBigInt(uint32(eth_trade.ETH_DECIMALS))
	tmpMount := int64(0.005 * float64(solDecimals.Int64()))
	routerResp, err := gmgnTool.GetEthTradeTool().GetAvailableRouter("eth", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		"0xcf91b70017eabde82c9671e30e5502d312ea6eb2", int(tmpMount))
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(routerResp)

	routerData, err := json.Marshal(routerResp.Data.Routes[0])
	fmt.Println(string(routerData))
	gmgnTool.GetEthTradeTool().Swap(&routerResp.Data.Routes[0], 10)
	// resp, err := gmgnTool.GetMobiTool().GetTokenPoolTool().GetPoolInfoEvm(gmgn_define.CHAIN_TYPE_ETH, "0xcf91b70017eabde82c9671e30e5502d312ea6eb2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	time.Sleep(50 * time.Second)
	// 	return
	// }

	// baseutils.ShowObjectValue(resp)
	// count := 0
	// for {
	// 	count++
	// 	_, err := gmgnTool.GetMobiTool().GetWalletStatTool().Get("4S9U8HckRngscHWrW418cG6Suw62dhEZzmyrT2hxSye5", "7d")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// baseutils.ShowObjectValue(resp)
	// 	fmt.Println("count = ", count)
	// }
}
