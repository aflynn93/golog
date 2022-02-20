package main

const (
	ERROR = "ERROR"
	FATAL = "FATAL"
)

// Initialize a GoLogger with a log level to be able to log messages depending on the configured log level
type GoLogger struct {
	LogLevel int
}

// Initialize a ErrorLogger to be able to log errors, regardless of log level
type ErrorLogger struct {
}
