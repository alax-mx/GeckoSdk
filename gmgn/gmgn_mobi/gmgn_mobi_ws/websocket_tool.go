package gmgn_mobi_ws

type STWebsocketTool struct {
	baseParam string
}

func NewSTWebsocketTool(baseParam string) *STWebsocketTool {
	return &STWebsocketTool{
		baseParam: baseParam,
	}
}
