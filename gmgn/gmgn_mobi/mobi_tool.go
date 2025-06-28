package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/proxy"
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
	smartMoneyCardsTool    *SmartMoneyCardsTool
	kolCardsTool           *KolCardsTool
	tokenWalletMonitorTool *TokenWalletMonitorTool
	tokenDevCreatedTool    *TokenDevCreatedTool
}

func NewMobiTool(baseUrl string, deviceInfo *DeviceInfo) *MobiTool {
	ret := &MobiTool{}

	// 没指定就用默认的
	if deviceInfo == nil {
		deviceInfo = &DeviceInfo{}
		data, _ := baseutils.ReadFile("device.json")
		json.Unmarshal(data, deviceInfo)
	}
	baseGetParam := ret.GetBaseGetParam(deviceInfo)
	basePostParam, _ := json.Marshal(deviceInfo)
	ret.tokenSecurityTool = NewTokenSecurityTool(baseUrl, baseGetParam)
	ret.tokenStatTool = NewTokenStatTool(baseUrl, baseGetParam)
	ret.tokenPriceTool = NewTokenPriceTool(baseUrl, baseGetParam)
	ret.tokenPoolTool = NewTokenPoolTool(baseUrl, baseGetParam)
	ret.tokenWalletTagStatTool = NewTokenWalletTagStatTool(baseUrl, baseGetParam)
	ret.tokenDevTool = NewTokenDevTool(baseUrl, baseGetParam)
	ret.tokenCandlesTool = NewTokenCandlesTool(baseUrl, baseGetParam)
	ret.gasPriceTool = NewGasPriceTool(baseUrl, baseGetParam)
	ret.tokenHoldersTool = NewTokenHoldersTool(baseUrl, baseGetParam)
	ret.tokenHolderStatTool = NewTokenHolderStatTool(baseUrl, baseGetParam)
	ret.tokenTopBuyersTool = NewTokenTopBuyersTool(baseUrl, baseGetParam)
	ret.tokenRugInfoTool = NewTokenRugInfoTool(baseUrl, baseGetParam)
	ret.tokenNewPairTool = NewTokenNewPairTool(baseUrl, baseGetParam)
	ret.walletHoldingsTool = NewWalletHoldingsTool(baseUrl, baseGetParam)
	ret.tokenLinkTool = NewTokenLinkTool(baseUrl, baseGetParam)
	ret.tokenLaunchpadInfoTool = NewTokenLaunchpadInfoTool(baseUrl, baseGetParam)
	ret.tokenPumpRankTool = NewTokenPumpRankTool(baseUrl, baseGetParam)
	ret.tokenBluchipRankTool = NewTokenBluchipRankTool(baseUrl, baseGetParam)
	ret.tokenPumpTool = NewTokenPumpTool(baseUrl, baseGetParam)
	ret.tokenSwapsTool = NewTokenSwapsTool(baseUrl, baseGetParam)
	ret.tokenDevCreatedTool = NewTokenDevCreatedTool(baseUrl, baseGetParam)
	ret.tokenWalletMonitorTool = NewTokenWalletMonitorTool(baseUrl, baseGetParam)
	ret.smartMoneyCardsTool = NewSmartMoneyCardsTool(baseUrl, baseGetParam, basePostParam)
	ret.kolCardsTool = NewKolCardsTool(baseUrl, baseGetParam, basePostParam)

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

func (mt *MobiTool) GetTokenDevCreatedTool() *TokenDevCreatedTool {
	return mt.tokenDevCreatedTool
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

func (mt *MobiTool) GetTokenWalletMonitorTool() *TokenWalletMonitorTool {
	return mt.tokenWalletMonitorTool
}

func (mt *MobiTool) GetSmartMoneyCardsTool() *SmartMoneyCardsTool {
	return mt.smartMoneyCardsTool
}

func (mt *MobiTool) GetKolCardsTool() *KolCardsTool {
	return mt.kolCardsTool
}

func (mt *MobiTool) GetBaseGetParam(deviceInfo *DeviceInfo) string {
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

func (mt *MobiTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	mt.tokenSecurityTool.SetProxy(proxyInfo)
	mt.tokenStatTool.SetProxy(proxyInfo)
	mt.tokenPriceTool.SetProxy(proxyInfo)
	mt.tokenPoolTool.SetProxy(proxyInfo)
	mt.tokenWalletTagStatTool.SetProxy(proxyInfo)
	mt.tokenDevTool.SetProxy(proxyInfo)
	mt.tokenCandlesTool.SetProxy(proxyInfo)
	mt.gasPriceTool.SetProxy(proxyInfo)
	mt.tokenHoldersTool.SetProxy(proxyInfo)
	mt.tokenHolderStatTool.SetProxy(proxyInfo)
	mt.tokenTopBuyersTool.SetProxy(proxyInfo)
	mt.tokenRugInfoTool.SetProxy(proxyInfo)
	mt.tokenNewPairTool.SetProxy(proxyInfo)
	mt.walletHoldingsTool.SetProxy(proxyInfo)
	mt.tokenLinkTool.SetProxy(proxyInfo)
	mt.tokenLaunchpadInfoTool.SetProxy(proxyInfo)
	mt.tokenPumpRankTool.SetProxy(proxyInfo)
	mt.tokenBluchipRankTool.SetProxy(proxyInfo)
	mt.tokenPumpTool.SetProxy(proxyInfo)
	mt.tokenSwapsTool.SetProxy(proxyInfo)
	mt.smartMoneyCardsTool.SetProxy(proxyInfo)
	mt.kolCardsTool.SetProxy(proxyInfo)
	mt.tokenWalletMonitorTool.SetProxy(proxyInfo)
	mt.tokenDevCreatedTool.SetProxy(proxyInfo)
}
