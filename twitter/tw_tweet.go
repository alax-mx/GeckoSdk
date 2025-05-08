package twitter

import (
	"fmt"
	"strconv"
)

type TwTweetTool struct {
	apiKey string
}

func NewTwTweetTool(apiKey string) *TwTweetTool {
	return &TwTweetTool{
		apiKey: apiKey,
	}
}

func (ttt *TwTweetTool) GetTweetDetails(userID string, count int, needTime bool) {
	url := "tweet/" + userID + "?count=" + strconv.Itoa(count)
	if needTime {
		url += "&includeTimestamp=true"
	} else {
		url += "&includeTimestamp=false"
	}

	data, err := HttpGet(url, ttt.apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
