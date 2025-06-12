package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
)

type MobiTool struct {
	tokenSecurityTool      *TokenSecurityTool
	tokenStatTool          *TokenStatTool
	tokenPriceTool         *TokenPriceTool
	tokenPoolTool          *TokenPoolTool
	tokenWalletTagStatTool *TokenWalletTagStatTool
	tokenDevTool           *TokenDevTool
	tokenCandlesTool       *TokenCandlesTool
	gasPriceTool           *GasPriceTool
	tokenHoldersTool       *TokenHoldersTool
	tokenHolderStatTool    *TokenHolderStatTool
	tokenTopBuyersTool     *TokenTopBuyersTool
	tokenRugInfoTool       *TokenRugInfoTool
	tokenNewPairTool       *TokenNewPairTool
	walletHoldingsTool     *WalletHoldingsTool
	tokenLinkTool          *TokenLinkTool
	tokenLaunchpadInfoTool *TokenLaunchpadInfoTool
	tokenPumpRankTool      *TokenPumpRankTool
	tokenBluchipRankTool   *TokenBluchipRankTool
	tokenPumpTool          *TokenPumpTool
	tokenSwapsTool         *TokenSwapsTool
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
	ret.tokenPoolTool = NewTokenPoolTool(baseUrl, baseParam)
	ret.tokenWalletTagStatTool = NewTokenWalletTagStatTool(baseUrl, baseParam)
	ret.tokenDevTool = NewTokenDevTool(baseUrl, baseParam)
	ret.tokenCandlesTool = NewTokenCandlesTool(baseUrl, baseParam)
	ret.gasPriceTool = NewGasPriceTool(baseUrl, baseParam)
	ret.tokenHoldersTool = NewTokenHoldersTool(baseUrl, baseParam)
	ret.tokenHolderStatTool = NewTokenHolderStatTool(baseUrl, baseParam)
	ret.tokenTopBuyersTool = NewTokenTopBuyersTool(baseUrl, baseParam)
	ret.tokenRugInfoTool = NewTokenRugInfoTool(baseUrl, baseParam)
	ret.tokenNewPairTool = NewTokenNewPairTool(baseUrl, baseParam)
	ret.walletHoldingsTool = NewWalletHoldingsTool(baseUrl, baseParam)
	ret.tokenLinkTool = NewTokenLinkTool(baseUrl, baseParam)
	ret.tokenLaunchpadInfoTool = NewTokenLaunchpadInfoTool(baseUrl, baseParam)
	ret.tokenPumpRankTool = NewTokenPumpRankTool(baseUrl, baseParam)
	ret.tokenBluchipRankTool = NewTokenBluchipRankTool(baseUrl, baseParam)
	ret.tokenPumpTool = NewTokenPumpTool(baseUrl, baseParam)
	ret.tokenSwapsTool = NewTokenSwapsTool(baseUrl, baseParam)
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

func (mt *MobiTool) GetTokenPoolTool() *TokenPoolTool {
	return mt.tokenPoolTool
}

func (mt *MobiTool) GetTokenWalletTagStatTool() *TokenWalletTagStatTool {
	return mt.tokenWalletTagStatTool
}

func (mt *MobiTool) GetTokenDevTool() *TokenDevTool {
	return mt.tokenDevTool
}

func (mt *MobiTool) GetTokenLinkTool() *TokenLinkTool {
	return mt.tokenLinkTool
}

func (mt *MobiTool) GetTokenLaunchpadInfoTool() *TokenLaunchpadInfoTool {
	return mt.tokenLaunchpadInfoTool
}

func (mt *MobiTool) GetTokenCandlesTool() *TokenCandlesTool {
	return mt.tokenCandlesTool
}

func (mt *MobiTool) GetGasPriceTool() *GasPriceTool {
	return mt.gasPriceTool
}

func (mt *MobiTool) GetTokenHoldersTool() *TokenHoldersTool {
	return mt.tokenHoldersTool
}

func (mt *MobiTool) GetTokenHolderStatTool() *TokenHolderStatTool {
	return mt.tokenHolderStatTool
}

func (mt *MobiTool) GetTokenTopBuyersTool() *TokenTopBuyersTool {
	return mt.tokenTopBuyersTool
}

func (mt *MobiTool) GetTokenRugInfoTool() *TokenRugInfoTool {
	return mt.tokenRugInfoTool
}

func (mt *MobiTool) GetTokenNewPairTool() *TokenNewPairTool {
	return mt.tokenNewPairTool
}

func (mt *MobiTool) GetWalletHoldingsTool() *WalletHoldingsTool {
	return mt.walletHoldingsTool
}

func (mt *MobiTool) GetTokenPumpRankTool() *TokenPumpRankTool {
	return mt.tokenPumpRankTool
}

func (mt *MobiTool) GetTokenBluchipRankTool() *TokenBluchipRankTool {
	return mt.tokenBluchipRankTool
}

func (mt *MobiTool) GetTokenPumpTool() *TokenPumpTool {
	return mt.tokenPumpTool
}

func (mt *MobiTool) GetTokenSwapsTool() *TokenSwapsTool {
	return mt.tokenSwapsTool
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
