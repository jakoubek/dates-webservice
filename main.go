package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jakoubek/dates-webservice/dates"
	_ "github.com/jakoubek/dates-webservice/l10n"
	"github.com/jakoubek/dates-webservice/requestlogger"
)

var requests int64
var requestsOld int64

var version string

func main() {

	s := NewServer("DatesAPI 1.0", getCounterfile())

	s.logger.Printf("Server Version %s is starting on %s...", version, getServerPort())
	s.logger.Printf("Counter file: %s...", getCounterfile())

	s.setupRoutes()

	http.ListenAndServe(getServerPort(), s.router)

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

func logRequestToConsole(req string) {
	log.Println("Logging request: ", req)
}

func logRequest() {
	requests++
}

type logRequestBody struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	Domain        string `json:"domain"`
	RemoteAddress string `json:"-"`
	path          string `json:"-"`
}

func NewLogRequestBody(path string, address string) *logRequestBody {
	return &logRequestBody{
		Name:          "pageview",
		Url:           fmt.Sprintf("https://api.datesapi.net%s", path),
		Domain:        "api.datesapi.net",
		RemoteAddress: address,
		path:          path,
	}
}

func logRequestToPlausible(lrb *logRequestBody) {

	//log.Println("Logging request:", lrb.path)

	postBody, err := json.Marshal(lrb)
	if err != nil {
		log.Println(err.Error())
	}

	responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request

	request, err := http.NewRequest("POST", "https://plausible.io/api/event", responseBody)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("User-Agent", "API")
	request.Header.Set("X-Forwarded-For", lrb.RemoteAddress)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()

}

func rootInfo(w http.ResponseWriter, r *http.Request) {

	type result struct {
		Result string   `json:"result"`
		Info   string   `json:"info"`
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
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.LastOfMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processToday(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.Today(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processTomorrow(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.Tomorrow(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processYesterday(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.Yesterday(),
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
		Result string `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result: dc.ThisYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processLastYear(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result: dc.LastYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processNextYear(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore()

	result := answer{
		Result: dc.NextYear(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processLastMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.LastMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processThisMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.ThisMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processNextMonth(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result string `json:"result"`
	}

	dc := dates.NewDateCore(
		dates.WithLanguage(r.URL.Query().Get("lang")),
		dates.WithFormat(r.URL.Query().Get("format")),
	)

	result := answer{
		Result: dc.NextMonth(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func processTimestamp(w http.ResponseWriter, r *http.Request) {

	logRequest()

	type answer struct {
		Result int64 `json:"result"`
	}

	result := answer{
		Result: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (s *server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(s.logInfo)
	}
}

func (s *server) handleHealthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Result string `json:"result"`
			Info   string `json:"info"`
		}{
			Result: "OK",
			Info:   "API fully operational",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
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
