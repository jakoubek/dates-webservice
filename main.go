package main

import (
	"encoding/json"
	"fmt"
	"github.com/jakoubek/dates-webservice/dates"
	_ "github.com/jakoubek/dates-webservice/l10n"
	"github.com/jakoubek/dates-webservice/requestlogger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var starttime time.Time
var requests int64
var requestsOld int64

func main() {
	starttime = time.Now()
	loadRequestsFromFile()
	initLogWriter()

	r := mux.NewRouter()
	r.HandleFunc("/", rootInfo).Methods("GET")
	r.HandleFunc("/today", processToday).Methods("GET")
	r.HandleFunc("/tomorrow", processTomorrow).Methods("GET")
	r.HandleFunc("/yesterday", processYesterday).Methods("GET")
	r.HandleFunc("/this-year", processThisYear).Methods("GET")
	r.HandleFunc("/last-year", processLastYear).Methods("GET")
	r.HandleFunc("/next-year", processNextYear).Methods("GET")
	r.HandleFunc("/last-month", processLastMonth).Methods("GET")
	r.HandleFunc("/this-month", processThisMonth).Methods("GET")
	r.HandleFunc("/next-month", processNextMonth).Methods("GET")
	r.HandleFunc("/last-of-month", processLastOfMonth)
	r.HandleFunc("/weeknumber", processWeeknumber).Methods("GET")
	r.HandleFunc("/timestamp", processTimestamp).Methods("GET")
	r.HandleFunc("/status", processStatus).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	log.Print("Starting server on " + getServerPort())
	http.ListenAndServe(getServerPort(), r)
}

func NotFound(w http.ResponseWriter, r *http.Request) {

	type result struct {
		Result string `json:"result"`
		Info   string `json:"info"`
	}

	response := &result{
		Result: "404 Not found",
		Info:   fmt.Sprintf("The requested resource %s was not found. Try starting with /", r.RequestURI),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)

}

func initLogWriter() {
	go func() {
		for true {
			if requests > requestsOld {
				requestlogger.SaveCounterfile(getCounterfile(), requests)
				requestsOld = requests
			}
			time.Sleep(5 * time.Minute)
		}
	}()
}

func loadRequestsFromFile() {
	requests = requestlogger.ReadCounterfile(getCounterfile())
	requestsOld = requests
}

func logRequest() {
	requests++
}

func rootInfo(w http.ResponseWriter, r *http.Request) {

	type result struct {
		Result string `json:"result"`
		Info   string `json:"info"`
		Routes []string `json:"routes"`
	}

	routes := []string{
		"https://api.datesapi.net/status",
		"https://api.datesapi.net/this-month",
		"https://api.datesapi.net/next-month",
		"https://api.datesapi.net/last-month",
		"https://api.datesapi.net/this-year",
		"https://api.datesapi.net/next-year",
		"https://api.datesapi.net/last-year",
		"https://api.datesapi.net/today",
		"https://api.datesapi.net/tomorrow",
		"https://api.datesapi.net/yesterday",
		"https://api.datesapi.net/last-of-month",
		"https://api.datesapi.net/weeknumber",
		"https://api.datesapi.net/timestamp",
	}

	response := result{
		Result: "OK",
		Info:   "Go to https://www.datesapi.net for information on how to access the API. See /status for API health.",
		Routes: routes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func processLastOfMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result:        dc.LastOfMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processToday(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.Today(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processTomorrow(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.Tomorrow(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processYesterday(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.Yesterday(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processWeeknumber(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result: dc.Weeknumber(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processThisYear(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result:        dc.ThisYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processLastYear(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result:        dc.LastYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processNextYear(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result:        dc.NextYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processLastMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.LastMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processThisMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.ThisMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processNextMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        string    `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
		)

	result := answer{
		Result:        dc.NextMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processTimestamp(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result        int64    `json:"result"`
	}

	result := answer{
		Result:        time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processStatus(w http.ResponseWriter, r *http.Request) {

	type answer struct {
		Result        string    `json:"result"`
		Info          string    `json:"info"`
		ServerStarted time.Time `json:"server_started"`
		Timestamp     int64     `json:"timestamp"`
		Requests      int64     `json:"requests"`
	}

	result := answer{
		Result:        "OK",
		Info:          "API fully operational",
		ServerStarted: starttime,
		Timestamp:     time.Now().Unix(),
		Requests:      requests,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func getCounterfile() string {
	if filename, ok := os.LookupEnv("COUNTERFILE"); ok {
		return filename
	}
	return "counter.json"
}

func getServerPort() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return ":3000"
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if you keep being sent to an "+
		"invalid page.</p>")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}