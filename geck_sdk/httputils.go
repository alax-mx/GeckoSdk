package geck_sdk

import (
	"io"
	"net/http"
)

var G_base_url = "https://api.geckoterminal.com/api/v2"

func HttpGet(url string) ([]byte, error) {

	req, _ := http.NewRequest("GET", G_base_url+url, nil)
	req.Header.Add("Accept", "application/json;version=20230302")
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
