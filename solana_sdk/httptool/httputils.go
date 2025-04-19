package httptool

import (
	"fmt"
	"io"
	"net/http"
)

func HttpGet(url string, apiKey string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("token", apiKey)
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

func HttpGetRug(token string) ([]byte, error) {
	url := fmt.Sprintf("https://api.rugcheck.xyz/v1/tokens/%s/report", token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
