package gmgn_mobi

import "encoding/json"

type STTokenLaunchpadInfo struct {
	Address           string `json:"address"`
	Launchpad         string `json:"launchpad"`
	LaunchpadStatus   int    `json:"launchpad_status"`
	LaunchpadProgress string `json:"launchpad_progress"`
	Description       string `json:"description"`
	LaunchpadPlatform string `json:"launchpad_platform"`
}

type GetTokenLaunchpadInfoResp struct {
	Code    int                  `json:"code"`
	Reason  string               `json:"reason"`
	Message string               `json:"message"`
	Data    STTokenLaunchpadInfo `json:"data"`
}

type TokenLaunchpadInfoTool struct {
	baseUrl   string
	baseParam string
}

func NewTokenLaunchpadInfoTool(baseUrl string, baseParam string) *TokenLaunchpadInfoTool {
	return &TokenLaunchpadInfoTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
	}
}

func (tpt *TokenLaunchpadInfoTool) Get(tokenAddress string) (*GetTokenLaunchpadInfoResp, error) {
	url := "api/v1/token_launchpad_info/sol/" + tokenAddress + "?" + tpt.baseParam
	data, err := HttpGet(tpt.baseUrl + url)
	if err != nil {
		return nil, err
	}

	ret := &GetTokenLaunchpadInfoResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}
