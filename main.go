package main

import (
	"encoding/json"
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	data, _ := baseutils.ReadFile("device.json")
	deviceInfo := &gmgn_mobi.DeviceInfo{}
	json.Unmarshal(data, deviceInfo)
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo)

	resp, err := gmgnTool.GetMobiTool().GetTokenStatTool().Get(gmgn_define.CHAIN_TYPE_BSC, "0x5448dff2d2564534befab1bb9266632fd67c4444")
	if err != nil {
		fmt.Println(err)
		return
	}

	baseutils.ShowObjectValue(resp)
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
