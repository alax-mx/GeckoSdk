package main

import (
	"encoding/json"
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	TestFreshToken()
}

func TestFreshToken() {
	deviceInfo := loaddevice("device.json")
	if deviceInfo == nil {
		return
	}
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo, nil, "")
	_, err := gmgnTool.GetMobiTool().GetTokenNewPairTool().Get("sol", gmgn_mobi.NEW_PAIR_PERIOD_1M, 5, gmgn_mobi.NEW_PAIR_ORDER_BY_CREATE_TIMESTAMP)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func loaddevice(path string) *gmgn_mobi.DeviceInfo {
	data, err := baseutils.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cfg := &gmgn_mobi.DeviceInfo{}
	err2 := json.Unmarshal(data, cfg)
	if err2 != nil {
		fmt.Println(err)
		return nil
	}
	return cfg
}
