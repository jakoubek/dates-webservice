package main

func (s *server) setupRoutes() {

	demoRoutes := s.router.Methods("GET").Subrouter()
	demoRoutes.Use(SetupRequestMiddleware)
	demoRoutes.HandleFunc("/demo", s.handleDemo()).Methods("GET")

	svcRoutes := s.router.Methods("GET").Subrouter()
	svcRoutes.HandleFunc("/today", processToday).Methods("GET")
	svcRoutes.HandleFunc("/tomorrow", processTomorrow).Methods("GET")
	svcRoutes.HandleFunc("/yesterday", processYesterday).Methods("GET")
	svcRoutes.HandleFunc("/this-year", processThisYear).Methods("GET")
	svcRoutes.HandleFunc("/last-year", processLastYear).Methods("GET")
	svcRoutes.HandleFunc("/next-year", processNextYear).Methods("GET")
	svcRoutes.HandleFunc("/last-month", processLastMonth).Methods("GET")
	svcRoutes.HandleFunc("/this-month", processThisMonth).Methods("GET")
	svcRoutes.HandleFunc("/next-month", processNextMonth).Methods("GET")
	svcRoutes.HandleFunc("/last-of-month", processLastOfMonth)
	svcRoutes.HandleFunc("/weeknumber", processWeeknumber).Methods("GET")
	svcRoutes.HandleFunc("/timestamp", processTimestamp).Methods("GET")
	svcRoutes.HandleFunc("/time", s.handleTime()).Methods("GET")
	svcRoutes.Use(SetupRequestMiddleware)
	svcRoutes.Use(LogRequestMiddleware)

	baseRoutes := s.router.Methods("GET").Subrouter()
	baseRoutes.HandleFunc("/", rootInfo).Methods("GET")
	baseRoutes.HandleFunc("/status", s.handleStatus()).Methods("GET")
	baseRoutes.HandleFunc("/healthz", s.handleHealthz()).Methods("GET")
	baseRoutes.NotFoundHandler = s.handleNotFound()
	baseRoutes.Use(SetupRequestMiddleware)

}
