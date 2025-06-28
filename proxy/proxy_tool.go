package proxy

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/proxy"
)

type STProxyInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (spi *STProxyInfo) Get() (*url.URL, error) {
	strUrl := "https://" + spi.Username + ":" + spi.Password
	strUrl += "@" + spi.Host + ":" + strconv.Itoa(spi.Port)
	fmt.Println("STProxyInfo strUrl = ", strUrl)
	ret, err := url.Parse(strUrl)
	if err != nil {
		fmt.Println(err)
	}
	return ret, err
}

func (spi *STProxyInfo) GetSocks5() (proxy.Dialer, error) {
	auth := &proxy.Auth{
		User:     spi.Username,
		Password: spi.Password,
	}
	// 创建拨号器
	proxyAddr := fmt.Sprintf("%s:%d", spi.Host, spi.Port)
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, auth, proxy.Direct)
	return dialer, err
}

type GetProxyListResp struct {
	Status string         `json:"status"`
	Count  string         `json:"count"`
	List   []*STProxyInfo `json:"list"`
}

type ProxyTool struct {
	proxyUrl      string
	proxyInfoList []*STProxyInfo
}

func NewProxyTool(proxyUrl string) *ProxyTool {
	return &ProxyTool{
		proxyUrl:      proxyUrl,
		proxyInfoList: make([]*STProxyInfo, 0),
	}
}

func (pt *ProxyTool) Update() bool {
	data, err := HttpGetFullURL(pt.proxyUrl)
	if err != nil {
		fmt.Println(err)
		return false
	}

	ret := &GetProxyListResp{}
	err = json.Unmarshal(data, ret)
	if err != nil {
		fmt.Println(err)
		return false
	}
	pt.proxyInfoList = ret.List
	return true
}

func (pt *ProxyTool) GetRandomProxyInfo() *STProxyInfo {
	count := len(pt.proxyInfoList)
	if count <= 0 {
		return nil
	}
	tmpRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return pt.proxyInfoList[tmpRand.Intn(count)]
}
