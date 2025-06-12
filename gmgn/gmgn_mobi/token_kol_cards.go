package gmgn_mobi

import "encoding/json"

type GetKolCardsResp struct {
	Code    int         `json:"code"`
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Data    STCardsData `json:"data"`
}

type KolCardsTool struct {
	baseUrl   string
	baseParam string
	postData  []byte
}

func NewKolCardsTool(baseUrl string, baseParam string, postData []byte) *KolCardsTool {
	return &KolCardsTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		postData:  postData,
	}
}

func (tdt *KolCardsTool) Get(interval string) (*GetKolCardsResp, error) {
	url := "api/v1/kol_cards/cards/sol/" + interval + "?" + tdt.baseParam
	data, err := HttpPost(tdt.baseUrl+url, tdt.postData)
	if err != nil {
		return nil, err
	}

	ret := &GetKolCardsResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
