package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao acessar API")
	}

	return ioutil.ReadAll(resp.Body)
}