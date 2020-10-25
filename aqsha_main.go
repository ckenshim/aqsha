
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PhilPage struct {
	Ticker string
	RequestStatus string
	KeyMetricsArr []KeyMetrics
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/phil/", philTownHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "pages/main.html")
}

func philTownHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/phil/" {
		http.NotFound(w, r)
		return
	}
	ticker := r.FormValue("ticker")
	fmpkey := r.FormValue("fmpkey")
	log.Printf("Request for %s with key %s", ticker, fmpkey)

	doCalc := false
	if ticker != "" && fmpkey != "" {
		doCalc = true
	}

	var page PhilPage
	if doCalc {
		page = calculateBigFiveNumbers(ticker, fmpkey)
	}

	t, _ := template.ParseFiles("pages/philtown.html")

	_ = t.Execute(w, page)
}

func calculateBigFiveNumbers(ticker string, fmpkey string) PhilPage {

	metricsJson := getKeyMetrics(ticker, fmpkey)
	keyMetrics, err := parseKeyMetrics(metricsJson)

	if err != nil {
		log.Print(err)
	}
	page := PhilPage{Ticker: ticker}

	if len(keyMetrics) == 0 {
		page.RequestStatus = fmt.Sprintf("Could not acquire results for %v", ticker)
	} else {
		page.RequestStatus = ""
		page.KeyMetricsArr = keyMetrics
	}

	return page
}