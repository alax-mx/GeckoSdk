package geck_sdk

import (
	"encoding/json"
	"strconv"
)

type STAttributes_Supported struct {
	Name string `json:"name"`
}

type STSupportedData struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes STAttributes_Supported `json:"attributes"`
}

type STSupportedResp struct {
	DataList []*STSupportedData `json:"data"`
	Errors   []*STErrors        `json:"errors"`
}

type SupportedTool struct {
}

func NewSupportedTool() *SupportedTool {
	return &SupportedTool{}
}

func (ndt *SupportedTool) GetSupported(page int) (*STSupportedResp, error) {
	newUrl := "/networks?page=" + strconv.Itoa(page)
	data, err := HttpGet(newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STSupportedResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
