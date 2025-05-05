package gmgn

import (
	"encoding/json"
	"strconv"
)

type GmgnTokenTool struct {
}

func NewGmgnTokenTool() *GmgnTokenTool {
	return &GmgnTokenTool{}
}

func (gtt *GmgnTokenTool) GetTokenInfo(tokenAddress string) (*STTokenInfo, error) {
	urlAddr := "/v1/tokens/sol/" + tokenAddress
	data, err := HttpGet(urlAddr)
	if err != nil {
		return nil, err
	}

	tokenInfo := &STTokenInfo{}
	err2 := json.Unmarshal(data, tokenInfo)
	if err2 != nil {
		return nil, err2
	}

	return tokenInfo, nil
}

func (gtt *GmgnTokenTool) GetNewTokens(limit int, orderBy string, direction string) (*GetTokenPairResp, error) {
	urlAddr := "/v1/pairs/sol/new_pairs?limit=" + strconv.Itoa(limit)
	urlAddr += "&orderby=" + orderBy + "&direction=" + direction
	data, err := HttpGet(urlAddr)
	if err != nil {
		return nil, err
	}

	resp := &GetTokenPairResp{}

	err2 := json.Unmarshal(data, resp)
	if err2 != nil {
		return nil, err2
	}

	return resp, nil
}
