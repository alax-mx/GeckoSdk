package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
)

type MobiTool struct {
	tokenSecurityTool *TokenSecurityTool
	tokenStatTool     *TokenStatTool
	tokenPriceTool    *TokenPriceTool
}

func NewMobiTool(baseUrl string, deviceInfo *DeviceInfo) *MobiTool {
	ret := &MobiTool{}

	// 没指定就用默认的
	if deviceInfo == nil {
		deviceInfo = &DeviceInfo{}
		data, _ := baseutils.ReadFile("device.json")
		json.Unmarshal(data, deviceInfo)
	}
	baseParam := ret.GetBaseParam(deviceInfo)
	ret.tokenSecurityTool = NewTokenSecurityTool(baseUrl, baseParam)
	ret.tokenStatTool = NewTokenStatTool(baseUrl, baseParam)
	ret.tokenPriceTool = NewTokenPriceTool(baseUrl, baseParam)
	return ret
}

func (mt *MobiTool) GetTokenSecurityTool() *TokenSecurityTool {
	return mt.tokenSecurityTool
}

func (mt *MobiTool) GetTokenStatTool() *TokenStatTool {
	return mt.tokenStatTool
}

func (mt *MobiTool) GetTokenPriceTool() *TokenPriceTool {
	return mt.tokenPriceTool
}

func (mt *MobiTool) GetBaseParam(deviceInfo *DeviceInfo) string {
	retStr := "device_id=" + deviceInfo.DeviceID
	retStr += "&client_id=" + deviceInfo.ClientID
	retStr += "&from_app=" + deviceInfo.FromApp
	retStr += "&app_ver=" + deviceInfo.AppVer
	retStr += "&pkg=" + deviceInfo.Pkg
	retStr += "&app_lang=" + deviceInfo.AppLang
	retStr += "&sys_lang=" + deviceInfo.SysLang
	retStr += "&brand=" + deviceInfo.Brand
	retStr += "&model=" + deviceInfo.Model
	retStr += "&os=" + deviceInfo.Os
	retStr += "&os_api=" + deviceInfo.OsAPI
	retStr += "&tz_name=" + deviceInfo.TzName
	retStr += "&tz_offset=" + deviceInfo.TzOffset
	retStr += "&gpv=" + deviceInfo.Gpv
	return retStr
}
