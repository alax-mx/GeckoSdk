package gmgn_http

import (
	"io"
	"net/http"
)

func HttpGetFullURL(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
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
