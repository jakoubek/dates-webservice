package main

import (
	"expvar"
	"net/http"
	"strconv"
	"time"

	"github.com/jakoubek/dates-webservice/internal/dates"
)

func (app *application) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := envelope{
			"index": 1,
		}

		err := app.writeJSON(w, http.StatusOK, data, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}
}

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	requests, _ := strconv.Atoi(expvar.Get("total_requests_received").String())
	data := envelope{
		"version":        version,
		"build_time":     buildTime,
		"server_started": app.startupTime.UTC().String(),
		"requests":       requests,
		"timestamp":      time.Now().Unix(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "OK",
		"info":   "API fully operational",
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) todayHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) tomorrowHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithDayAdd(1),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) yesterdayHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithDayAdd(-1),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) thisYearHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithYearFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) lastYearHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithYearAdd(-1),
		dates.WithYearFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) nextYearHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithYearAdd(1),
		dates.WithYearFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) thisMonthHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithMonthFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) nextMonthHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithMonthAdd(1),
		dates.WithMonthFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) lastMonthHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithMonthAdd(-1),
		dates.WithMonthFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) lastOfMonthHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithLastOfMonth(),
		dates.WithDayFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) weeknumberHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithWeeknumber(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) timestampHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithTimestamp(false),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) timestampMsHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithTimestamp(true),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) timeHandler(w http.ResponseWriter, r *http.Request) {

	dc := dates.NewDateCore(
		dates.WithDatetimeFormat(),
		dates.WithUserFormat(getFormatFromContext(r.Context())),
		dates.WithLanguage(getLangFromContext(r.Context())),
	)

	data := envelope{
		"result": dc.ResultString(),
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
