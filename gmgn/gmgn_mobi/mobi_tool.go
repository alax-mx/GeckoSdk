package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
)

type MobiTool struct {
	baseUrl    string
	deviceInfo *DeviceInfo
}

func NewMobiTool(baseUrl string, deviceInfo *DeviceInfo) *MobiTool {
	ret := &MobiTool{
		baseUrl: baseUrl,
	}

	// 没指定就用默认的
	ret.deviceInfo = deviceInfo
	if deviceInfo == nil {
		ret.deviceInfo = &DeviceInfo{}
		data, _ := baseutils.ReadFile("device.json")
		json.Unmarshal(data, ret.deviceInfo)
	}
	return ret
}
