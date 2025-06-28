package gmgn_trade

import (
	"bytes"
	"io"
	"net/http"

	"github.com/alax-mx/geckosdk/proxy"
)

func HttpGet(url string, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Host", "gmgn.ai")
	req.Header.Add("accept", "'application/json, text/plain, */*")
	req.Header.Add("accept-language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("dnt", "1")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://gmgn.ai/?chain=sol")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:82.0) Gecko/20100101 Firefox/82.0")

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

func HttpPostRouter(url string, data []byte, proxyInfo *proxy.STProxyInfo) ([]byte, error) {
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
	res, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
