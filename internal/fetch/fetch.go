package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func FetchJson(url string, target interface{}) error {
	data, err := Fetch(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}
