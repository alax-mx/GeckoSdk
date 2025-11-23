package gmgn

import (
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_trade"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_web"
	"github.com/alax-mx/geckosdk/proxy"
)

type GmgnTool struct {
	webTool      *gmgn_web.WebTool
	mobiTool     *gmgn_mobi.MobiTool
	solTradeTool *gmgn_trade.SolTradeTool
	evmTradeTool *gmgn_trade.EvmTradeTool
}

func NewGmgnTool(pubKey string, priKey string, deviceInfo *gmgn_mobi.DeviceInfo, evmConfig *gmgn_trade.STEvmConfig, authStr string) *GmgnTool {
	ret := &GmgnTool{
		webTool:      gmgn_web.NewWebTool(gmgn_define.G_BASE_GMGN_WEB_DEFI_URL),
		mobiTool:     gmgn_mobi.NewMobiTool(gmgn_define.G_BASE_GMGN_MOBI_URL, deviceInfo, authStr),
		solTradeTool: nil,
		evmTradeTool: nil,
	}
	if evmConfig != nil {
		ret.evmTradeTool = gmgn_trade.NewEvmTradeTool(evmConfig)
	} else {
		ret.solTradeTool = gmgn_trade.NewSolTradeTool(gmgn_define.G_BASE_GMGN_SOL_TRADE_URL, pubKey, priKey)
	}
	return ret
}

func (gt *GmgnTool) GetWebTool() *gmgn_web.WebTool {
	return gt.webTool
}

func (gt *GmgnTool) GetMobiTool() *gmgn_mobi.MobiTool {
	return gt.mobiTool
}

func (gt *GmgnTool) GetSolTradeTool() *gmgn_trade.SolTradeTool {
	return gt.solTradeTool
}

func (gt *GmgnTool) GetEvmTradeTool() *gmgn_trade.EvmTradeTool {
	return gt.evmTradeTool
}

func (gt *GmgnTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	gt.mobiTool.SetProxy(proxyInfo)
	if gt.solTradeTool != nil {
		gt.solTradeTool.SetProxy(proxyInfo)
	}
}
