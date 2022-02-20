package golog

const (
	ERROR   = "ERROR"
	FATAL   = "FATAL"
	INFO    = 1
	DEBUG   = 2
	VERBOSE = 3
)

// Initialize a GoLogger with a log level to be able to log messages depending on the configured log level
type GoLogger struct {
	LogLevel int
}

// Initialize a ErrorLogger to be able to log errors, regardless of log level
type ErrorLogger struct {
}
