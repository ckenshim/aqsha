
package main

import (
	"testing"
)


func TestKeyMetricsUrl(t *testing.T){
	url := keyMetricsUrl("TICKER", "KEY123", 10)

	if url != "https://fmpcloud.io/api/v3/key-metrics/TICKER?limit=10&apikey=KEY123" {
		t.Errorf("Incorrect URL returned %v", url)
	}
}

func TestGetKeyMetrics(t *testing.T){
	body := getKeyMetrics("AAPL", "77dc5e892fbed4514a3ed8e6f5370438")

	t.Errorf(body)
}
