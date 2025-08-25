package tgbot

import "fmt"

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

func (tgb *TGBotTool) SendText(msg string) {
	url := "https://api.telegram.org/bot" + tgb.TGToken + "/sendMessage?chat_id=" + tgb.GroupID
	url += "&text=" + msg
	body, err := HttpGetFullURL(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
