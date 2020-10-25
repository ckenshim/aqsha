package main

import (
	"io/ioutil"
	"testing"
)


func TestKeyMetricsParser(t *testing.T) {
	dat, err := ioutil.ReadFile("sample_test_files/keymetrics.json")

	if err != nil {
		t.Errorf("Error while reading the file %v", err)
	}

	metrics, err := parseKeyMetrics(dat)

	if err != nil {
		t.Errorf("Error while decoding. %v", err)
	}

	if len(metrics) != 7 {
		t.Errorf("Expecting key metrics for the last 7 years")
	}

	if metrics[0].Symbol != "AHH" {
		t.Errorf("Ticker symbol is not correct. Expecting AHH")
	}
}

