package gmgn

import (
	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_trade"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_web"
)

type GmgnTool struct {
	webTool   *gmgn_web.WebTool
	mobiTool  *gmgn_mobi.MobiTool
	tradeTool *gmgn_trade.TradeTool
}

func NewGmgnTool(pubKey string, priKey string, deviceInfo *gmgn_mobi.DeviceInfo) *GmgnTool {
	return &GmgnTool{
		webTool:   gmgn_web.NewWebTool(gmgn_define.G_BASE_GMGN_WEB_DEFI_URL),
		mobiTool:  gmgn_mobi.NewMobiTool(gmgn_define.G_BASE_GMGN_MOBI_URL, deviceInfo),
		tradeTool: gmgn_trade.NewTradeTool(gmgn_define.G_BASE_GMGN_TRADE_URL, pubKey, priKey),
	}
}

func (gt *GmgnTool) GetWebTool() *gmgn_web.WebTool {
	return gt.webTool
}

func (gt *GmgnTool) GetMobiTool() *gmgn_mobi.MobiTool {
	return gt.mobiTool
}

func (gt *GmgnTool) GetTradeTool() *gmgn_trade.TradeTool {
	return gt.tradeTool
}
