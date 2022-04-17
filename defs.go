package golog

const (
	ERROR   = "ERROR"
	FATAL   = "FATAL"
	INFO    = 1
	DEBUG   = 2
	VERBOSE = 3
)

var (
	logLevel int
)
