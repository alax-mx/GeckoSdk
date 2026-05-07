package main

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/alax-mx/geckosdk/baseutils"
	"github.com/alax-mx/geckosdk/gmgn/gmgn_mobi"
)

func main() {
	TestFreshToken()
}

func TestFreshToken() {
	// deviceInfo := loaddevice("device.json")
	// if deviceInfo == nil {
	// 	return
	// }
	// gmgnTool := gmgn.NewGmgnTool("", "", deviceInfo, nil, "")
	// _, err := gmgnTool.GetMobiTool().GetTokenNewPairTool().Get("sol", gmgn_mobi.NEW_PAIR_PERIOD_1M, 5, gmgn_mobi.NEW_PAIR_ORDER_BY_CREATE_TIMESTAMP)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	TestWSS()
}

func loaddevice(path string) *gmgn_mobi.DeviceInfo {
	data, err := baseutils.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cfg := &gmgn_mobi.DeviceInfo{}
	err2 := json.Unmarshal(data, cfg)
	if err2 != nil {
		fmt.Println(err)
		return nil
	}
	return cfg
}

// Go 语言最常见的生成方式（你现在用的很可能就是这个）
func generateID() string {
	b := make([]byte, 16)
	rand.Read(b) // crypto/rand
	return hex.EncodeToString(b)
	// 输出示例：ffe7e5e5214b8a02
}

func TestWSS() {
	// 1. 创建跳过证书验证的 Dialer（只用于测试/自签证书！！！）
	dialer := &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "ws.gmgn.mobi",
			SessionTicketsDisabled: true,
		},
	}

	// 2. 生产环境用系统根证书池（推荐）
	// dialer := websocket.DefaultDialer // 默认就是这个，已经自动加载系统根证书

	// 3. 连接 wss
	url := "wss://ws.gmgn.mobi/bg_ws?uuid=63ee67fad1f2cdf0&device_id=1d51590f146adc56&client_id=gmgn_android_20405001&from_app=gmgn&app_ver=20405001&pkg=com.gmgn.app&app_lang=zh-CN&sys_lang=zh-CN&brand=HUAWEI&model=LMR-AL10&os=android&os_api=9&tz_name=Asia%2FShanghai&tz_offset=-480&gpv=10000001" // 必须是 wss://
	// 如果你需要带 token
	header := http.Header{}
	header.Add("User-Agent", "ReactNative")

	conn, resp, err := dialer.Dial(url, header)
	if err != nil {
		log.Fatal("wss 连接失败:", err)
		if resp != nil {
			log.Println("HTTP 状态码:", resp.StatusCode)
		}
		return
	}
	defer conn.Close()

	log.Println("wss 连接成功！")

	data := []byte("{\"id\":\"00e1fad8d50ba404\",\"action\":\"subscribe\",\"channel\":\"new_pool_info\",\"data\":[{\"chain\":\"sol\"}]}")
	conn.WriteMessage(websocket.TextMessage, data)

	for {
		var msg map[string]any
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		msgType, _ := msg["type"].(string)
		switch msgType {
		case "welcome":
			log.Println("服务器欢迎:", msg["msg"])
		case "pong":
			// log.Println("收到服务器心跳回应")
		case "chat":
			log.Printf("新消息 from %s: %s", msg["user"], msg["content"])

		case "notification":
			log.Println("通知:", msg["title"], msg["body"])

		default:
			// 打印未知消息，方便调试
			data, _ := json.MarshalIndent(msg, "", "  ")
			log.Println("收到消息:", string(data))
		}
	}
}
