package main

import (
	"fmt"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn"
)

func main() {
	tokenSecurityTool := gmgn.NewGmgnTool("", "", nil).GetMobiTool().GetTokenPriceTool()
	resp, err := tokenSecurityTool.GetTokenPriceInfo("8YLdiDNrQgnJ6ki9TSeEmdwxhEtesu2hMLx7vLb43o54")
	if err != nil {
		fmt.Println(err)
		return
	}
	baseutils.ShowObjectValue(resp)
}
