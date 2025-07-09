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
	resp, err := gmgnTool.GetMobiTool().GetTokenRecommendSlippageTool().Get("2mU4wMdQGmugVvTZJ7GcSY2orr7h4bA9VHBX7W6zbonk")
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
}
