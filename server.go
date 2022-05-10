package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// server is a struct for the server instance
type server struct {
	serverName  string
	router      *mux.Router
	logger      *log.Logger
	logInfo     *logStruct
	counterFile string
}

// logStruct is a struct for the log info structure
type logStruct struct {
	ServerStartedAt time.Time   `json:"server_started"`
	Requests        int         `json:"requests"`
	LastRequestAt   time.Time   `json:"lastrequest"`
	Routes          []*logRoute `json:"routes"`
	isDirty         bool
}

// logRoute is a struct for logging the requests
// to one route
type logRoute struct {
	RouteName     string    `json:"route"`
	Requests      int       `json:"requests"`
	LastRequestAt time.Time `json:"lastrequest"`
}

// NewServer is the factory function for returning a
// server instance.
func NewServer(serverName, counterFile string) *server {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	s := &server{
		serverName:  serverName,
		router:      mux.NewRouter(),
		logger:      logger,
		logInfo:     nil,
		counterFile: counterFile,
	}
	s.InitLogStruct()
	s.initLogWriter()
	return s
}

// handleNotFound is the handler function to respond
// on requests for not defined routes.
func (s *server) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//s.logRouteRequest("/notfound")
		response := struct {
			Code    string `json:code`
			Message string `json:message`
		}{
			Code:    "404",
			Message: fmt.Sprintf("Route %s not found", r.URL.Path),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
	}
}

// InitLogStruct initializes the log info structure
// with default values.
func (s *server) InitLogStruct() {
	s.logInfo = &logStruct{
		Requests: 0,
		Routes:   nil,
	}
	s.readCounterFile()
	s.logInfo.ServerStartedAt = time.Now().UTC()
}

// initLogWriter is a separate goroutine that writes
// the server's log info structure to the JSON file
// - only if there was a request since the last write.
func (s *server) initLogWriter() {
	go func() {
		for true {
			if s.logInfo.isDirty {
				s.saveCounterFile()
			}
			time.Sleep(5 * time.Minute)
		}
	}()
}

// readCounterFile reads the from the counterFile JSON file (if existing)
// and populates the server's log info structure.
// Gets executed at server start one (and only - obviously - if the file exists).
func (s *server) readCounterFile() {
	info, err := os.Stat(s.counterFile)
	if os.IsNotExist(err) {
		s.logger.Printf("No counter file found - initialized new log structure")
		return
	}
	if !info.IsDir() {
		file, _ := ioutil.ReadFile(s.counterFile)
		err = json.Unmarshal([]byte(file), &s.logInfo)
		if err != nil {
			return
		}
		s.logger.Printf("Initialized log structure from existing counter file %s", s.counterFile)
	}
	return
}

// saveCounterFile saves the server's current log info structure to a JSON file.
// The JSON filename are taken from the server instance (counterFile).
func (s *server) saveCounterFile() {
	s.logger.Printf("writing stats to counter file %s", s.counterFile)
	file, _ := json.Marshal(s.logInfo)
	_ = ioutil.WriteFile(s.counterFile, file, 0644)
	s.logInfo.isDirty = false
}
