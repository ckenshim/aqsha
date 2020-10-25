package main

import "encoding/json"

type KeyMetrics struct {
	Symbol string
	Date string
	NetIncomePerShare float64
	BookValuePerShare float64
	Roic float64
}

func parseKeyMetrics(jsonData []byte) ([]KeyMetrics, error) {
	var keyMetricsArr []KeyMetrics
	err := json.Unmarshal(jsonData, &keyMetricsArr)

	return keyMetricsArr, err
}