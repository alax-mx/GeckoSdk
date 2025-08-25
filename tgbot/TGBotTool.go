package tgbot

type TGBotTool struct {
	TGToken string
	GroupID string
}

func NewTGBotTool(token string, groupID string) *TGBotTool {
	return &TGBotTool{
		TGToken: token,
		GroupID: groupID,
	}
}

func (tgb *TGBotTool) SendText(msg string) error {
	url := "https://api.telegram.org/bot" + tgb.TGToken + "/sendMessage?chat_id=" + tgb.GroupID
	url += "&text=" + msg
	_, err := HttpGetFullURL(url)
	if err != nil {
		return err
	}

	return nil
}
