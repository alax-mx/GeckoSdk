package gmgn_mobi

import "encoding/json"

type STRuggedTokens struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
	Name    string `json:"name"`
	Logo    string `json:"logo"`
}

type STTokenRugData struct {
	Address         string           `json:"address"`
	RugRatio        string           `json:"rug_ratio"`
	HolderRuggedNum int              `json:"holder_rugged_num"`
	HolderTokenNum  int              `json:"holder_token_num"`
	RuggedTokens    []STRuggedTokens `json:"rugged_tokens"`
}

type GetTokenRugInfoResp struct {
	Code    int            `json:"code"`
	Reason  string         `json:"reason"`
	Message string         `json:"message"`
	Data    STTokenRugData `json:"data"`
}

type TokenRugInfoTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenRugInfoTool(baseUrl string, baseParam string) *TokenRugInfoTool {
	return &TokenRugInfoTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tit *TokenRugInfoTool) Get(tokenAddress string) (*GetTokenRugInfoResp, error) {
	url := "api/v1/token_rug_info/sol/" + tokenAddress + "?" + tit.baseParam
	data, err := HttpGet(tit.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenRugInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
