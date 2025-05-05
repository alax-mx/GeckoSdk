package gecko_sdk

import (
	"io"
	"net/http"
)

var G_base_url = "https://pro-api.coingecko.com/api/v3/onchain"

func HttpGet(apiKey, url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", G_base_url+url, nil)
	req.Header.Add("Accept", "application/json;version=20230302")
	req.Header.Add("x-cg-pro-api-key", apiKey)
	res, err := http.DefaultClient.Do(req)
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
