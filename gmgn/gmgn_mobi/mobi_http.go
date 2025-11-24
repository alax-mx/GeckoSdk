package gmgn_mobi

import (
	"bytes"
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/alax-mx/geckosdk/proxy"
)

var g_configList []*tls.Config

func GetRanomClpherSuites() *tls.Config {
	if len(g_configList) <= 0 {
		g_configList = append(g_configList, &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "gmgn.gracematrix.net",
			SessionTicketsDisabled: true,
		})
		g_configList = append(g_configList, &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "gmgn.gracematrix.net",
			SessionTicketsDisabled: true,
		})
		g_configList = append(g_configList, &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "gmgn.gracematrix.net",
			SessionTicketsDisabled: true,
		})
		g_configList = append(g_configList, &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "gmgn.gracematrix.net",
			SessionTicketsDisabled: true,
		})
		g_configList = append(g_configList, &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12,
			//可以添加其他配置，如支持的Extensions等
			ServerName:             "gmgn.gracematrix.net",
			SessionTicketsDisabled: true,
		})
	}

	tmpRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return g_configList[tmpRand.Intn(5)]
}

func HttpGet(url string, authStr string, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("user-agent", "okhttp/4.9.2")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if len(authStr) > 0 {
		req.Header.Add("authorization", "Bearer "+authStr)
	}
	// client := http.DefaultClient
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: GetRanomClpherSuites(),
		},
	}
	if proxyInfo != nil {
		dialer, err := proxyInfo.GetSocks5()
		if err == nil {
			client = &http.Client{
				Transport: &http.Transport{
					Dial: dialer.Dial,
				},
			}
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}

func HttpPost(url string, param []byte, authStr string, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: GetRanomClpherSuites(),
		},
	}
	if proxyInfo != nil {
		dialer, err := proxyInfo.GetSocks5()
		if err != nil {
			client = &http.Client{
				Transport: &http.Transport{
					Dial: dialer.Dial,
				},
			}
		}
	}
	req, _ := http.NewRequest("POST", url, bytes.NewReader(param))
	req.Header.Add("user-agent", "okhttp/4.9.2")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if len(authStr) > 0 {
		req.Header.Add("authorization", "Bearer "+authStr)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return nil, err2
	}

	return body, nil
}
