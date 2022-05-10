package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jakoubek/dates-webservice/dates"
)

func (s *server) sendResponse(w http.ResponseWriter, r *http.Request, dc dates.DateCore) {

	response := struct {
		Result string `json:"result"`
	}{
		Result: dc.ResultString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (s *server) handleDemo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dcValue := r.Context().Value("datecore")
		dc, ok := dcValue.(dates.DateCore)
		if ok {
			s.sendResponse(w, r, dc)
		} else {
			s.sendResponse(w, r, dc)
		}

	}
}

func SetupRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "datecore", dates.NewDateCore())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Host, "localhost") {
			log.Println("Log request for route:", r.RequestURI)
		} else {
			go logRequestToPlausible(NewLogRequestBody(r.RequestURI, r.Header.Get("X-Forwarded-For")))
		}
		next.ServeHTTP(w, r)
	})
}
