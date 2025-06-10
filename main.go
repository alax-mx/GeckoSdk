package main

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	data, _ := baseutils.ReadFile("device.json")
	deviceInfo := &gmgn_mobi.DeviceInfo{}
	json.Unmarshal(data, deviceInfo)
	gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo)
	resp, _ := gmgnTool.GetMobiTool().GetTokenLaunchpadInfoTool().Get("Asuu49BSxjz4Qj9hkdrwt3GpxdjvYUfbt5rXnsNWpump")
	baseutils.ShowObjectValue(resp)
}
