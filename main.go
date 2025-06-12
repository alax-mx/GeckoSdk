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
	tag := make([]string, 0)
	tag = append(tag, gmgn_mobi.WALLET_TAG_PUMP_SMART)
	resp, err := gmgnTool.GetMobiTool().GetTokenWalletMonitorTool().Get(gmgn_mobi.WALLET_ORDERBY_PNL_7D, 1000, tag)
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
}
