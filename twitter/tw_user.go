package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/alax-mx/geckosdk/twitter/basedef"
)

type TwUserTool struct {
	apiKey string
}

func NewTwUserTool(apiKey string) *TwUserTool {
	return &TwUserTool{
		apiKey: apiKey,
	}
}

func (tut *TwUserTool) GetUserDetailsByUserName(userName string) (*basedef.STUserDetailResp, error) {
	url := "user/by/username/" + userName

	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	ret := &basedef.STUserDetailResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) GetUserDetailsByUserID(userID string) (*basedef.STUserDetailResp, error) {
	url := "user/" + userID

	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	ret := &basedef.STUserDetailResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (tut *TwUserTool) GetTwttersByUserName(userName string, count int) (*basedef.STUserDetailResp, error) {
	url := "user/-1/tweets?count=" + strconv.Itoa(count) + "&username=" + userName

	data, err := HttpGet(url, tut.apiKey)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(data))
	// ret := &STUserDetailsResp{}
	// err = json.Unmarshal(data, ret)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
