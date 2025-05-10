package tw

import (
	"io"
	"net/http"
)

var G_BASE_URL = "https://api.twitterapi.io/twitter/"

func HttpGet(url string, apiKey string) ([]byte, error) {
	req, _ := http.NewRequest("GET", G_BASE_URL+url, nil)
	req.Header.Add("X-API-Key", apiKey)
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
