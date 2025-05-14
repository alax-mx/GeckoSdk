package gmgn_web

import (
	"encoding/json"
	"strconv"

	"github.com/alax-mx/geckosdk/gmgn/gmgn_define"
)

type WebTokenTool struct {
	baseUrl string
}

func NewWebTokenTool(baseUrl string) *WebTokenTool {
	return &WebTokenTool{
		baseUrl: baseUrl,
	}
}

func (gtt *WebTokenTool) GetTokenInfo(tokenAddress string) (*gmgn_define.STTokenInfo, error) {
	urlAddr := "v1/tokens/sol/" + tokenAddress
	data, err := HttpGet(gtt.baseUrl + urlAddr)
	if err != nil {
		return nil, err
	}

	tokenInfo := &gmgn_define.STTokenInfo{}
	err2 := json.Unmarshal(data, tokenInfo)
	if err2 != nil {
		return nil, err2
	}

	return tokenInfo, nil
}

func (gtt *WebTokenTool) GetNewTokens(limit int, orderBy string, direction string) (*gmgn_define.GetTokenPairResp, error) {
	urlAddr := "v1/pairs/sol/new_pairs?limit=" + strconv.Itoa(limit)
	urlAddr += "&orderby=" + orderBy + "&direction=" + direction
	data, err := HttpGet(gtt.baseUrl + urlAddr)
	if err != nil {
		return nil, err
	}

	resp := &gmgn_define.GetTokenPairResp{}

	err2 := json.Unmarshal(data, resp)
	if err2 != nil {
		return nil, err2
	}

	return resp, nil
}

func (gtt *WebTokenTool) GetTokenPrice(tokenAddress string) (*gmgn_define.GetTokenPriceResp, error) {
	urlAddr := "v1/sol/tokens/realtime_token_price?address=" + tokenAddress
	data, err := HttpGet(gtt.baseUrl + urlAddr)
	if err != nil {
		return nil, err
	}

	ret := &gmgn_define.GetTokenPriceResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
