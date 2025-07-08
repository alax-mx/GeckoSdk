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
	resp, err := gmgnTool.GetMobiTool().GetTokenTradesTool().Get("9Ertjp8ZpBwuEt1xoqz8gAPe6tvLCTkXvJYNqK35bonk", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
}
