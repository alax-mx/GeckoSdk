package gecko_sdk

import (
	"encoding/json"
	"strconv"
)

type STAttributes_Dexes struct {
	Name string `json:"name"`
}

type STDexesData struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes STAttributes_Dexes `json:"attributes"`
}

type STDexesResp struct {
	DataList []*STDexesData `json:"data"`
	Errors   []*STErrors    `json:"errors"`
}

type DexesTool struct {
	apiKey string
}

func NewDexesTool(apiKey string) *DexesTool {
	return &DexesTool{
		apiKey: apiKey,
	}
}

func (ndt *DexesTool) GetDexes(network string, page int) (*STDexesResp, error) {
	newUrl := "/networks/" + network + "/dexes?page=" + strconv.Itoa(page)
	data, err := HttpGet(ndt.apiKey, newUrl)
	if err != nil {
		return nil, err
	}

	ret := &STDexesResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
