package gmgn_web

type WebTool struct {
	tokenTool    *WebTokenTool
	trendingTool *WebTrendingTool
}

func NewWebTool(baseUrl string) *WebTool {
	return &WebTool{
		tokenTool:    NewWebTokenTool(baseUrl),
		trendingTool: NewWebTrendingTool(baseUrl),
	}
}

func (gt *WebTool) GetTokenTool() *WebTokenTool {
	return gt.tokenTool
}

func (gt *WebTool) GetTrendingTool() *WebTrendingTool {
	return gt.trendingTool
}
