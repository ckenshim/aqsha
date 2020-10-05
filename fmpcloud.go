package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getKeyMetrics(ticker string, apikey string) string {
	url := keyMetricsUrl(ticker, apikey, 10)
	resp, err := http.Get(url)
	if err != nil {
		// handle errors
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func keyMetricsUrl(ticker string, apiKey string, limit int) string{
	return fmt.Sprintf("https://fmpcloud.io/api/v3/key-metrics/%s?limit=%d&apikey=%s", ticker, limit, apiKey)
}
