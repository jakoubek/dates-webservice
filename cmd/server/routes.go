package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	//router.Use(middleware.Logger)

	//router.Use(app.initContext)
	router.Use(app.metrics)
	router.Use(app.recoverPanic)
	router.Use(app.rateLimit)
	router.Use(app.enableCORS)

	router.NotFound(http.HandlerFunc(app.notFoundResponse))
	router.MethodNotAllowed(http.HandlerFunc(app.methodNotAllowedResponse))

	router.Get("/", app.indexHandler)
	router.Get("/status", app.statusHandler)
	router.Get("/healthz", app.healthcheckHandler)

	router.Group(func(router chi.Router) {
		router.Use(app.checkNoLogging)
		router.Use(app.logRequests)
		router.Use(app.logRequestsToDatabase)
		router.Use(app.readHeaders)
		router.Use(app.readQueryParams)

		router.Get("/today", app.todayHandler)
		router.Get("/tomorrow", app.tomorrowHandler)
		router.Get("/yesterday", app.yesterdayHandler)
		router.Get("/this-year", app.thisYearHandler)
		router.Get("/next-year", app.nextYearHandler)
		router.Get("/last-year", app.lastYearHandler)
		router.Get("/this-month", app.thisMonthHandler)
		router.Get("/next-month", app.nextMonthHandler)
		router.Get("/last-month", app.lastMonthHandler)
		router.Get("/last-of-month", app.lastOfMonthHandler)
		router.Get("/weeknumber", app.weeknumberHandler)
		router.Get("/timestamp", app.timestampHandler)
		router.Get("/timestampms", app.timestampMsHandler)
		router.Get("/time", app.timeHandler)
	})

	return router
}
