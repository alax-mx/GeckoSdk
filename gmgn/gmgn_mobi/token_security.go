package gmgn_mobi

import (
	"encoding/json"

	"github.com/alax-mx/geckosdk/proxy"
)

type STLockSummary struct {
	IsLocked        bool   `json:"is_locked"`
	LockDetail      any    `json:"lock_detail"`
	LockTags        any    `json:"lock_tags"`
	LockPercent     string `json:"lock_percent"`
	LeftLockPercent string `json:"left_lock_percent"`
}
type STSecurity struct {
	Address                string        `json:"address"`
	IsShowAlert            bool          `json:"is_show_alert"`
	Top10HolderRate        string        `json:"top_10_holder_rate"`
	RenouncedMint          bool          `json:"renounced_mint"`
	RenouncedFreezeAccount bool          `json:"renounced_freeze_account"`
	BurnRatio              string        `json:"burn_ratio"`
	BurnStatus             string        `json:"burn_status"`
	DevTokenBurnAmount     string        `json:"dev_token_burn_amount"`
	DevTokenBurnRatio      string        `json:"dev_token_burn_ratio"`
	IsOpenSource           any           `json:"is_open_source"`
	OpenSource             int           `json:"open_source"`
	IsBlacklist            any           `json:"is_blacklist"`
	Blacklist              int           `json:"blacklist"`
	IsHoneypot             any           `json:"is_honeypot"`
	Honeypot               int           `json:"honeypot"`
	IsRenounced            any           `json:"is_renounced"`
	Renounced              any           `json:"renounced"`
	CanSell                int           `json:"can_sell"`
	CanNotSell             int           `json:"can_not_sell"`
	BuyTax                 string        `json:"buy_tax"`
	SellTax                string        `json:"sell_tax"`
	AverageTax             string        `json:"average_tax"`
	HighTax                string        `json:"high_tax"`
	Flags                  []any         `json:"flags"`
	LockInfo               any           `json:"lockInfo"`
	LockSummary            STLockSummary `json:"lock_summary"`
	HideRisk               bool          `json:"hide_risk"`
}
type STTokenSecurity struct {
	Address   string     `json:"address"`
	Security  STSecurity `json:"security"`
	Launchpad any        `json:"launchpad"`
}

type GetMutilWindowTokenSecurityResp struct {
	Code    int             `json:"code"`
	Reason  string          `json:"reason"`
	Message string          `json:"message"`
	Data    STTokenSecurity `json:"data"`
}

type TokenSecurityTool struct {
	baseUrl   string
	baseParam string
	proxyInfo *proxy.STProxyInfo
}

func NewTokenSecurityTool(baseUrl string, baseParam string) *TokenSecurityTool {
	return &TokenSecurityTool{
		baseUrl:   baseUrl,
		baseParam: baseParam,
		proxyInfo: nil,
	}
}

func (gst *TokenSecurityTool) SetProxy(proxyInfo *proxy.STProxyInfo) {
	gst.proxyInfo = proxyInfo
}

func (gst *TokenSecurityTool) Get(chainType string, tokenAddress string) (*GetMutilWindowTokenSecurityResp, error) {
	url := "api/v1/mutil_window_token_security_launchpad/" + chainType + "/" + tokenAddress + "?" + gst.baseParam
	data, err := HttpGet(gst.baseUrl+url, gst.proxyInfo)
	if err != nil {
		return nil, err
	}
	ret := &GetMutilWindowTokenSecurityResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
