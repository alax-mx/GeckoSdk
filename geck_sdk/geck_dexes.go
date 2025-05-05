package geck_sdk

import (
	"encoding/json"
	"fmt"
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
}

func NewDexesTool() *DexesTool {
	return &DexesTool{}
}

func (ndt *DexesTool) GetDexes(network string, page int) (*STDexesResp, error) {
	newUrl := "/networkss/" + network + "/dexes?page=" + strconv.Itoa(page)
	fmt.Println("newUrl = ", newUrl)
	data, err := HttpGet(newUrl)
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
