package gmgn_mobi

import (
	"bytes"
	"io"
	"net/http"

	"github.com/alax-mx/geckosdk/proxy"
)

func HttpGet(url string, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("user-agent", "okhttp/4.9.2")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	client := http.DefaultClient
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

func HttpPost(url string, param []byte, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
	client := http.DefaultClient
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
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 9.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Mobile Safari/537.36")
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
