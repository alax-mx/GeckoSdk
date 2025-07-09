package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	data, _ := baseutils.ReadFile("device.json")
	deviceInfo := &gmgn_mobi.DeviceInfo{}
	json.Unmarshal(data, deviceInfo)
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo)
	resp, err := gmgnTool.GetMobiTool().GetTokenPoolTool().Get("2mU4wMdQGmugVvTZJ7GcSY2orr7h4bA9VHBX7W6zbonk")
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
	poolTime := time.Unix(int64(resp.Data.CreationTimestamp), 0)
	fmt.Println("poolTime := ", poolTime)
}
