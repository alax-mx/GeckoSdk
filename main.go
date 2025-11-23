package main

import (
	"encoding/json"
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
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
