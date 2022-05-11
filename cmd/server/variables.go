package main

import "github.com/jakoubek/dates-webservice/internal/vcs"

var (
	version     = vcs.Version()
	buildTime   string
	isDebugMode string = "false"
)
