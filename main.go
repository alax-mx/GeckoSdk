package main

import (
	"encoding/json"
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	data, _ := baseutils.ReadFile("device.json")
	deviceInfo := &gmgn_mobi.DeviceInfo{}
	json.Unmarshal(data, deviceInfo)
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo)
	// resp, err := gmgnTool.GetMobiTool().GetTokenBluchipRankTool().Get("1h", 1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	resp, err := gmgnTool.GetMobiTool().GetWalletStatTool().Get("DfMxre4cKmvogbLrPigxmibVTTQDuzjdXojWzjCXXhzj", "7d")
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
}
