package gecko_sdk

import (
	"encoding/json"
	"strconv"
)

type STAttributes_Supported struct {
	Name                     string `json:"name"`
	CoingeckoAssetPlatformID string `json:"coingecko_asset_platform_id"`
}

type STSupportedData struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes STAttributes_Supported `json:"attributes"`
}

type STSupportLinks struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type STSupportedResp struct {
	DataList []*STSupportedData `json:"data"`
	Links    STSupportLinks     `json:"links"`
	Errors   []*STErrors        `json:"errors"`
}

type SupportedTool struct {
	apiKey string
}

func NewSupportedTool(apiKey string) *SupportedTool {
	return &SupportedTool{
		apiKey: apiKey,
	}
}

func (ndt *SupportedTool) GetSupported(page int) (*STSupportedResp, error) {
	newUrl := "/networks?page=" + strconv.Itoa(page)
	data, err := HttpGet(ndt.apiKey, newUrl)
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
