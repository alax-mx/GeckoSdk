package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/proxy"
)

type MobiTool struct {
	tokenSecurityTool          *TokenSecurityTool
	tokenStatTool              *TokenStatTool
	tokenPriceTool             *TokenPriceTool
	tokenPoolTool              *TokenPoolTool
	tokenWalletTagStatTool     *TokenWalletTagStatTool
	tokenDevTool               *TokenDevTool
	tokenCandlesTool           *TokenCandlesTool
	gasPriceTool               *GasPriceTool
	tokenHoldersTool           *TokenHoldersTool
	tokenHolderStatTool        *TokenHolderStatTool
	tokenTopBuyersTool         *TokenTopBuyersTool
	tokenRugInfoTool           *TokenRugInfoTool
	tokenNewPairTool           *TokenNewPairTool
	walletHoldingsTool         *WalletHoldingsTool
	walletStatTool             *WalletStatTool
	walletActivity             *WalletActivity
	walletTokenActivity        *WalletTokenActivity
	tokenLinkTool              *TokenLinkTool
	tokenLaunchpadInfoTool     *TokenLaunchpadInfoTool
	tokenPumpRankTool          *TokenPumpRankTool
	tokenBluchipRankTool       *TokenBluchipRankTool
	tokenPumpTool              *TokenPumpTool
	tokenSwapsTool             *TokenSwapsTool
	smartMoneyCardsTool        *SmartMoneyCardsTool
	kolCardsTool               *KolCardsTool
	tokenWalletMonitorTool     *TokenWalletMonitorTool
	tokenDevCreatedTool        *TokenDevCreatedTool
	tokenTradesTool            *TokenTradesTool
	tokenLiquidityTrendTool    *TokenLiquidityTrendTool
	tokenLiquidityDetailTool   *TokenLiquidityDetailTool
	tokenLiquidityStatsTool    *TokenLiquidityStatsTool
	tokenRecommendSlippageTool *TokenRecommendSlippageTool
	swapBatchOrderTool         *SwapBatchOrderTool
	accountTool                *AccountTool
}

func NewMobiTool(baseUrl string, deviceInfo *DeviceInfo, authStr string) *MobiTool {
	ret := &MobiTool{}

	// 没指定就用默认的
	if deviceInfo == nil {
		deviceInfo = &DeviceInfo{}
		data, _ := baseutils.ReadFile("device.json")
		json.Unmarshal(data, deviceInfo)
	}
	baseGetParam := ret.GetBaseGetParam(deviceInfo)
	basePostParam, _ := json.Marshal(deviceInfo)
	ret.tokenSecurityTool = NewTokenSecurityTool(baseUrl, baseGetParam, authStr)
	ret.tokenStatTool = NewTokenStatTool(baseUrl, baseGetParam, authStr)
	ret.tokenPriceTool = NewTokenPriceTool(baseUrl, baseGetParam, authStr)
	ret.tokenPoolTool = NewTokenPoolTool(baseUrl, baseGetParam, authStr)
	ret.tokenWalletTagStatTool = NewTokenWalletTagStatTool(baseUrl, baseGetParam, authStr)
	ret.tokenDevTool = NewTokenDevTool(baseUrl, baseGetParam, authStr)
	ret.tokenCandlesTool = NewTokenCandlesTool(baseUrl, baseGetParam, authStr)
	ret.gasPriceTool = NewGasPriceTool(baseUrl, baseGetParam, authStr)
	ret.tokenHoldersTool = NewTokenHoldersTool(baseUrl, baseGetParam, authStr)
	ret.tokenHolderStatTool = NewTokenHolderStatTool(baseUrl, baseGetParam, authStr)
	ret.tokenTopBuyersTool = NewTokenTopBuyersTool(baseUrl, baseGetParam, authStr)
	ret.tokenRugInfoTool = NewTokenRugInfoTool(baseUrl, baseGetParam, authStr)
	ret.tokenNewPairTool = NewTokenNewPairTool(baseUrl, baseGetParam, authStr)
	ret.walletHoldingsTool = NewWalletHoldingsTool(baseUrl, baseGetParam, authStr)
	ret.walletStatTool = NewWalletStatTool(baseUrl, baseGetParam, authStr)
	ret.walletActivity = NewWalletActivity(baseUrl, baseGetParam, authStr)
	ret.walletTokenActivity = NewWalletTokenActivity(baseUrl, baseGetParam, authStr)
	ret.tokenLinkTool = NewTokenLinkTool(baseUrl, baseGetParam, authStr)
	ret.tokenLaunchpadInfoTool = NewTokenLaunchpadInfoTool(baseUrl, baseGetParam, authStr)
	ret.tokenPumpRankTool = NewTokenPumpRankTool(baseUrl, baseGetParam, authStr)
	ret.tokenBluchipRankTool = NewTokenBluchipRankTool(baseUrl, baseGetParam, authStr)
	ret.tokenPumpTool = NewTokenPumpTool(baseUrl, baseGetParam, authStr)
	ret.tokenSwapsTool = NewTokenSwapsTool(baseUrl, baseGetParam, authStr)
	ret.tokenDevCreatedTool = NewTokenDevCreatedTool(baseUrl, baseGetParam, authStr)
	ret.tokenWalletMonitorTool = NewTokenWalletMonitorTool(baseUrl, baseGetParam, authStr)
	ret.tokenTradesTool = NewTokenTradesTool(baseUrl, baseGetParam, authStr)
	ret.smartMoneyCardsTool = NewSmartMoneyCardsTool(baseUrl, baseGetParam, basePostParam, authStr)
	ret.kolCardsTool = NewKolCardsTool(baseUrl, baseGetParam, basePostParam, authStr)
	ret.tokenLiquidityTrendTool = NewTokenLiquidityTrendTool(baseUrl, baseGetParam, authStr)
	ret.tokenLiquidityDetailTool = NewTokenLiquidityDetailTool(baseUrl, baseGetParam, authStr)
	ret.tokenLiquidityStatsTool = NewTokenLiquidityStatsTool(baseUrl, baseGetParam, authStr)
	ret.tokenRecommendSlippageTool = NewTokenRecommendSlippageTool(baseUrl, baseGetParam, authStr)
	ret.swapBatchOrderTool = NewSwapBatchOrderTool(baseUrl, baseGetParam, authStr)
	ret.accountTool = NewAccountTool(baseUrl, baseGetParam)

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

func (mt *MobiTool) GetTokenTradesTool() *TokenTradesTool {
	return mt.tokenTradesTool
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

func (mt *MobiTool) GetWalletStatTool() *WalletStatTool {
	return mt.walletStatTool
}

func (mt *MobiTool) GetWalletActivity() *WalletActivity {
	return mt.walletActivity
}

func (mt *MobiTool) GetWalletTokenActivity() *WalletTokenActivity {
	return mt.walletTokenActivity
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

func (mt *MobiTool) GetTokenLiquidityTrendTool() *TokenLiquidityTrendTool {
	return mt.tokenLiquidityTrendTool
}

func (mt *MobiTool) GetTokenLiquidityDetailTool() *TokenLiquidityDetailTool {
	return mt.tokenLiquidityDetailTool
}

func (mt *MobiTool) GetTokenLiquidityStatsTool() *TokenLiquidityStatsTool {
	return mt.tokenLiquidityStatsTool
}

func (mt *MobiTool) GetTokenRecommendSlippageTool() *TokenRecommendSlippageTool {
	return mt.tokenRecommendSlippageTool
}

func (mt *MobiTool) GetSwapBatchOrderTool() *SwapBatchOrderTool {
	return mt.swapBatchOrderTool
}

func (mt *MobiTool) GetAccountTool() *AccountTool {
	return mt.accountTool
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
	mt.walletStatTool.SetProxy(proxyInfo)
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
	mt.swapBatchOrderTool.SetProxy(proxyInfo)
	mt.accountTool.SetProxy(proxyInfo)
}

func (mt *MobiTool) SetAuthString(authStr string) {
	mt.tokenSecurityTool.SetAuthString(authStr)
	mt.tokenStatTool.SetAuthString(authStr)
	mt.tokenPriceTool.SetAuthString(authStr)
	mt.tokenPoolTool.SetAuthString(authStr)
	mt.tokenWalletTagStatTool.SetAuthString(authStr)
	mt.tokenDevTool.SetAuthString(authStr)
	mt.tokenCandlesTool.SetAuthString(authStr)
	mt.gasPriceTool.SetAuthString(authStr)
	mt.tokenHoldersTool.SetAuthString(authStr)
	mt.tokenHolderStatTool.SetAuthString(authStr)
	mt.tokenTopBuyersTool.SetAuthString(authStr)
	mt.tokenRugInfoTool.SetAuthString(authStr)
	mt.tokenNewPairTool.SetAuthString(authStr)
	mt.walletHoldingsTool.SetAuthString(authStr)
	mt.walletStatTool.SetAuthString(authStr)
	mt.tokenLinkTool.SetAuthString(authStr)
	mt.tokenLaunchpadInfoTool.SetAuthString(authStr)
	mt.tokenPumpRankTool.SetAuthString(authStr)
	mt.tokenBluchipRankTool.SetAuthString(authStr)
	mt.tokenPumpTool.SetAuthString(authStr)
	mt.tokenSwapsTool.SetAuthString(authStr)
	mt.smartMoneyCardsTool.SetAuthString(authStr)
	mt.kolCardsTool.SetAuthString(authStr)
	mt.tokenWalletMonitorTool.SetAuthString(authStr)
	mt.tokenDevCreatedTool.SetAuthString(authStr)
	mt.swapBatchOrderTool.SetAuthString(authStr)
}
