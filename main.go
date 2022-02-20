package golog

import (
	"fmt"
	"log"

	"github.com/juju/errors"
)

// LogMessage() - Takes a log level, which is then compared to the configured log level. If the level provided
// is lower than or equal to the configured log level, the message is logged.
func (g GoLogger) LogMessage(messageLogLevel int, message ...interface{}) {
	if messageLogLevel <= g.LogLevel {
		fmt.Println(message...)
	}
}

// CreateError() - Returns an error with a stack trace
func (e ErrorLogger) CreateError(message ...interface{}) error {
	return errors.Mask(errors.New(fmt.Sprint(message...)))
}

// LogError() - Log error, regardless of logging level. Include stack trace. Provide the error type - ERROR or FATAL.
// ERROR will allow app to continue, FATAL will cause app to shut down.
func (e ErrorLogger) LogError(errorType string, err error) {
	if errorType == ERROR {
		log.Printf("\n%+v\n\n", err)
	} else if errorType == FATAL {
		log.Fatalf("\n%+v\n\n", err)
	} else {
		log.Printf("%s%s\n", "Invalid errorType:", errorType)
	}
}
