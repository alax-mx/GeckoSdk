package main

import (
	"log"
	"math/big"
	"time"

	"github.com/alax-mx/geckosdk/gmgn/evm_trade"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// data, _ := baseutils.ReadFile("device.json")
	// deviceInfo := &gmgn_mobi.DeviceInfo{}
	// json.Unmarshal(data, deviceInfo)
	// gmgnTool := gmgn.NewGmgnTool("",
	// 	"",
	// 	deviceInfo, gmgn_define.CHAIN_TYPE_ETH)

	// solDecimals := solana.DecimalsInBigInt(uint32(eth_trade.ETH_DECIMALS))
	// tmpMount := int64(0.005 * float64(solDecimals.Int64()))
	// routerResp, err := gmgnTool.GetEthTradeTool().GetAvailableRouter("eth", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	// 	"0xcf91b70017eabde82c9671e30e5502d312ea6eb2", int(tmpMount))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// baseutils.ShowObjectValue(routerResp)

	// routerData, err := json.Marshal(routerResp.Data.Routes[0])
	// fmt.Println(string(routerData))
	// gmgnTool.GetEthTradeTool().Swap(&routerResp.Data.Routes[0], 10)
	// resp, err := gmgnTool.GetMobiTool().GetTokenPoolTool().GetPoolInfoEvm(gmgn_define.CHAIN_TYPE_ETH, "0xcf91b70017eabde82c9671e30e5502d312ea6eb2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	time.Sleep(50 * time.Second)
	// 	return
	// }

	// rpcURL := "https://mainnet.infura.io/v3/c696611f821e4cd8a8eb3911c2e5d5a7"
	// pubKey := "0x87efc84DF132e006012387B41288FB28A336e1bb"
	// priKey := "0x5af616d5aad4b117da457cf8719ee4b6a113dbbe987f431e72508c0549834aae"
	// ethTool := gmgn_trade.NewETHTadeTool2(rpcURL, pubKey, priKey, nil)
	// solDecimals := solana.DecimalsInBigInt(uint32(evm_trade.ETH_DECIMALS))
	// ethAmount := big.NewInt(int64(0.005 * float64(solDecimals.Int64())))
	// ethTool.Approve("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", ethAmount) // gmgn_trade.NewETHTadeTool2(nil, nil).Swap("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xcf91b70017eabde82c9671e30e5502d312ea6eb2", 0.005)

	// amount, err := ethTool.Allowance("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	// if err != nil {
	// 	println("Allowance err:", err.Error())
	// 	return
	// }
	// println("Allowance:", amount.Int64())
	// baseutils.ShowObjectValue(resp)
	// count := 0
	// for {
	// 	count++
	// 	_, err := gmgnTool.GetMobiTool().GetWalletStatTool().Get("4S9U8HckRngscHWrW418cG6Suw62dhEZzmyrT2hxSye5", "7d")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// baseutils.ShowObjectValue(resp)
	// 	fmt.Println("count = ", count)
	// }
	TestEth()
}

func TestEth() {
	// 创建价格监控器
	routerAddress := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	monitor, err := evm_trade.NewPriceMonitor("https://mainnet.infura.io/v3/c696611f821e4cd8a8eb3911c2e5d5a7", routerAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 定义交易路径: ETH -> USDT
	path := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // WETH
		common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), // USDT
	}

	// 监控价格变化
	go monitor.MonitorPriceChanges(path, 10*time.Second, func(price *big.Float) {
		log.Printf("当前价格: 1 ETH = %s USDT", price.String())
	})

	select {}
}
