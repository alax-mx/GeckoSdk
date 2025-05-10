package tw

import (
	"encoding/json"
	"errors"
)

type GetUserInfoByIDsResp struct {
	Status string       `json:"status"`
	Msg    string       `json:"msg"`
	Users  []STUserInfo `json:"users"`
}

type GetUserInfoResp struct {
	Status string     `json:"status"`
	Msg    string     `json:"msg"`
	Data   STUserInfo `json:"data"`
}

type STTweetData struct {
	Tweets []STTweetInfo `json:"tweets"`
}

type GetUserLastTweetsResp struct {
	Status      string      `json:"status"`
	Code        int         `json:"code"`
	Msg         string      `json:"msg"`
	Data        STTweetData `json:"data"`
	HasNextPage bool        `json:"has_next_page"`
	NextCursor  string      `json:"next_cursor"`
}

type GetUserFollwingsResp struct {
	Followings []STFollowingInfo `json:"followings"`
}

type AdvancedSearchResp struct {
	Tweets      []STTweetInfo `json:"tweets"`
	HasNextPage bool          `json:"has_next_page"`
	NextCursor  string        `json:"next_cursor"`
}

type TwUserTool struct {
	apiKey string
}

func NewTwUserTool(apiKey string) *TwUserTool {
	return &TwUserTool{
		apiKey: apiKey,
	}
}

func (tut *TwUserTool) GetUserInfoByIDs(userIds []string) (*GetUserInfoByIDsResp, error) {
	if len(userIds) <= 0 {
		return nil, errors.New("GetUserInfoByIDs err: userIds is null")
	}
	userIdstr := userIds[0]
	for i := 1; i < len(userIds); i++ {
		userIdstr += "," + userIds[i]
	}
	url := "user/batch_info_by_ids?userIds=" + userIdstr
	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &GetUserInfoByIDsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) GetUserInfoByName(userName string) (*GetUserInfoResp, error) {
	url := "user/info?userName=" + userName
	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &GetUserInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) GetUserLastTweets(userID string, userName string, cursor string) (*GetUserLastTweetsResp, error) {
	if len(userID) <= 0 && len(userName) <= 0 {
		return nil, errors.New("GetUserLastTweets err: userId and userName all is null")
	}

	url := "user/last_tweets"
	if len(userID) > 0 {
		url += "?userId=" + userID
	} else {
		url += "?userName=" + userName
	}

	if len(cursor) > 0 {
		url += "&cursor=" + cursor
	}
	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &GetUserLastTweetsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) GetUserFollwings(userName string, cursor string) (*GetUserFollwingsResp, error) {
	url := "user/followings?userName=" + userName

	if len(cursor) > 0 {
		url += "&cursor=" + cursor
	}
	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &GetUserFollwingsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) AdvancedSearch(query string, queryType string, cursor string) (*AdvancedSearchResp, error) {
	url := "tweet/advanced_search?query=" + query + "&queryType=" + queryType
	if len(cursor) > 0 {
		url += "&cursor=" + cursor
	}
	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &AdvancedSearchResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
