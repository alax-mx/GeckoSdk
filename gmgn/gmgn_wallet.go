package gmgn

import "fmt"

type GmgnWalletTool struct {
}

func NewGmgnWalletTool() *GmgnWalletTool {
	return &GmgnWalletTool{}
}

func (gwt *GmgnWalletTool) GetWalletTokenDistribution(walletAddress string) {
	// url := "/v1/rank/sol/wallets/" + walletAddress + "/unique_token_7d?interval=7d"
	url := "/v1/smartmoney/sol/walletNew/" + walletAddress + "?period=7d"
	data, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
